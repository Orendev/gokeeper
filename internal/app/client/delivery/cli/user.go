package cli

import (
	"context"
	"fmt"
	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/user"
	domainUser "github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
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
			)

			us, err := d.ucUserStorage.Add(ctx, *dUser)
			if err != nil {
				logger.Log.Info("create user", zap.Error(err))
				return
			}

			userExternal, err := d.ucUserClient.Register(ctx, *us)
			if err != nil {
				logger.Log.Info("register user", zap.Error(err))
				return
			}

			dUser.SetToken(userExternal.Token())

			_, err = d.ucUserStorage.Update(ctx, *dUser)
			if err != nil {
				logger.Log.Info("update token user", zap.Error(err))
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

			var dUser *domainUser.User

			userInternal, err := d.ucUserStorage.Get(ctx)
			if err != nil {
				fmt.Printf("User locally: %s\n", err.Error())

				dUser, err = domainUser.NewWithID(
					userExternal.ID(),
					*passwordUser,
					*emailUser,
					userExternal.Role(),
					userExternal.Name(),
					userExternal.Token(),
					userExternal.CreatedAt(),
					time.Now().UTC(),
				)
				if err != nil {
					fmt.Printf("Error create User locally: %s\n", err.Error())
					return
				}

				_, err := d.ucUserStorage.Add(ctx, *dUser)
				if err != nil {
					fmt.Printf("Error create User locally: %s\n", err.Error())
					return
				}

			} else {
				dUser, err = domainUser.NewWithID(
					userInternal.ID(),
					*passwordUser,
					*emailUser,
					userInternal.Role(),
					userInternal.Name(),
					userExternal.Token(),
					userInternal.CreatedAt(),
					time.Now().UTC(),
				)

				_, err = d.ucUserStorage.Update(ctx, *dUser)
				if err != nil {
					fmt.Printf("Error update token user: %s\n", err.Error())
					return
				}
			}

		},
	}
}

func (d *Delivery) getUser() *cobra.Command {

	return &cobra.Command{
		Use:     "getUser",
		Aliases: []string{"get"},
		Short:   "Get user in the service.",
		Long:    `This command find user: Keeper client getUser.`,

		Run: func(cmd *cobra.Command, args []string) {

			userFind, err := d.ucUserStorage.Get(context.Background())
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			msg := fmt.Sprintf("ID: %s\nName: %s\nEmail: %s\nRole: %s\nToken: %s\nPassword: %s\nCreatedAt: %s\nUpdatedAt: %s\n",
				userFind.ID().String(),
				userFind.Name().String(),
				userFind.Email().String(),
				userFind.Role().String(),
				userFind.Token().String(),
				userFind.Password().String(),
				userFind.CreatedAt().String(),
				userFind.UpdatedAt().String(),
			)

			fmt.Println(msg)
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
