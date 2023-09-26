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
)

var userArgs user.RegisterUserArgs
var loginUserArgs user.LoginUserArgs

func (d *Delivery) createUser() *cobra.Command {

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

			userRegister, err := d.ucUserClient.Register(ctx, *us)
			if err != nil {
				logger.Log.Info("register user", zap.Error(err))
				return
			}

			dUser.SetToken(userRegister.Token())

			_, err = d.ucUserStorage.UpdateToken(ctx, *dUser)
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

			//_, err = d.ucUserStorage.Login(context.Background(), *emailUser, *passwordUser)
			_, err = d.ucUserClient.Login(context.Background(), *emailUser, *passwordUser)
			if err != nil {
				logger.Log.Info("create user", zap.Error(err))
				return
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

func initCreateUserArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&userArgs.Name, "name", "n", "", "user name value.")
	cmd.Flags().StringVarP(&userArgs.Email, "email", "e", "", "user email value.")
	cmd.Flags().StringVarP(&userArgs.Password, "password", "p", "", "user password value.")
}

func initLoginUserArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&userArgs.Email, "email", "e", "", "user email value.")
	cmd.Flags().StringVarP(&userArgs.Password, "password", "p", "", "user password value.")
	err := cmd.MarkFlagRequired("email")
	if err != nil {
		return
	}
	err = cmd.MarkFlagRequired("password")
	if err != nil {
		return
	}
}
