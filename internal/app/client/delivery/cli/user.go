package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/user"
	domainUser "github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/logger"
	memory "github.com/Orendev/gokeeper/pkg/tools/fileStorage"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var userArgs user.RegisterUserArgs
var loginUserArgs user.LoginUserArgs

func (d *Delivery) registerUser() *cobra.Command {

	return &cobra.Command{
		Use:     "registerUser",
		Aliases: []string{"regUs"},
		Short:   "Register new user in the service.",
		Long:    `This command register a new user: Keeper client registerUser --name=<name> --email=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			nameUser, err := name.New(userArgs.Name)
			if err != nil {
				logger.Log.Error("name input fields", zap.Error(err))
				return
			}

			emailUser, err := email.New(userArgs.Email)
			if err != nil {
				logger.Log.Error("email input fields", zap.Error(err))
				return
			}

			passwordUser, err := password.New(userArgs.Password)
			if err != nil {
				logger.Log.Error("password input fields", zap.Error(err))
				return
			}

			dUser, err := domainUser.New(
				*passwordUser,
				*emailUser,
				*nameUser,
				*token.New(""),
			)

			_, err = d.ucUserStorage.Create(ctx, dUser)
			if err != nil {
				logger.Log.Info("create user", zap.Error(err))
				return
			}

			userExternal, err := d.ucUserClient.Create(ctx, dUser)
			if err != nil {
				logger.Log.Info("register user", zap.Error(err))
				return
			}

			dUser.SetToken(userExternal.Token())

			if !d.ucUserStorage.SetToken(ctx, dUser) {
				fmt.Printf("Error update token user: %s\n", err.Error())
				return
			}

			err = d.fileStorage.Save(memory.Memory{
				ID:       dUser.ID().String(),
				Email:    dUser.Email().String(),
				Password: dUser.Password().String(),
				Token:    dUser.Token().String(),
			})
			if err != nil {
				logger.Log.Info("register user", zap.Error(err))
				return
			}
		},
	}
}

func (d *Delivery) loginUser() *cobra.Command {

	return &cobra.Command{
		Use:     "loginUser",
		Aliases: []string{"logUs"},
		Short:   "Login user in the service.",
		Long:    `This command login a user: Keeper client loginUser --email=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			emailUser, err := email.New(loginUserArgs.Email)
			if err != nil {
				logger.Log.Error("email input fields", zap.Error(err))
				return
			}

			passwordUser, err := password.New(loginUserArgs.Password)
			if err != nil {
				logger.Log.Error("password input fields", zap.Error(err))
				return
			}

			userExternal, err := d.ucUserClient.Login(ctx, *emailUser, *passwordUser)
			if err != nil {
				logger.Log.Info("user authorization", zap.Error(err))
				return
			}

			dUser, err := domainUser.NewWithID(
				userExternal.ID(),
				*passwordUser,
				*emailUser,
				userExternal.Name(),
				userExternal.Role(),
				userExternal.CreatedAt(),
				time.Now().UTC(),
			)

			dUser.SetToken(userExternal.Token())

			err = d.fileStorage.Save(memory.Memory{
				ID:       dUser.ID().String(),
				Email:    dUser.Email().String(),
				Password: dUser.Password().String(),
				Token:    dUser.Token().String(),
			})
			if err != nil {
				logger.Log.Info("register user", zap.Error(err))
				return
			}

			var parameter queryParameter.QueryParameter
			parameter.Pagination.Limit = 1
			parameter.Pagination.Offset = 0

			total, err := d.ucUserStorage.Count(ctx, parameter)
			if err != nil {
				logger.Log.Info("user authorization", zap.Error(err))
				return
			}

			if total == 0 {
				_, err = d.ucUserStorage.Create(ctx, dUser)
				if err != nil {
					logger.Log.Info("create user", zap.Error(err))
					return
				}

				return
			}

			if !d.ucUserStorage.SetToken(ctx, dUser) {
				fmt.Println("Error update token user")
				return
			}

		},
	}
}

func initRegisterUserArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&userArgs.Name, "name", "n", "", "user name value.")
	cmd.Flags().StringVarP(&userArgs.Email, "email", "e", "", "user email value.")
	cmd.Flags().StringVarP(&userArgs.Password, "password", "p", "", "user password value.")
}

func initLoginUserArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&loginUserArgs.Email, "email", "e", "", "user email value.")
	cmd.Flags().StringVarP(&loginUserArgs.Password, "password", "p", "", "user password value.")
	err := cmd.MarkFlagRequired("email")
	if err != nil {
		return
	}
	err = cmd.MarkFlagRequired("password")
	if err != nil {
		return
	}
}
