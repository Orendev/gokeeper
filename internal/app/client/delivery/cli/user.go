package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/user"
	domainUser "github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user/name"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var userArgs user.RegisterUserArgs

func (d *Delivery) createUser() *cobra.Command {

	return &cobra.Command{
		Use:     "registerUser",
		Aliases: []string{"regUs"},
		Short:   "Register new user in the service.",
		Long:    `This command register a new user: Keeper client registerUser --name=<name> --email=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {

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

			_, err = d.ucUser.Create(context.Background(), *dUser)
			if err != nil {
				logger.Log.Info("create user", zap.Error(err))
			}
		},
	}
}

func (d *Delivery) findUser() *cobra.Command {

	return &cobra.Command{
		Use:     "findUser",
		Aliases: []string{"findUs"},
		Short:   "Find user in the service.",
		Long:    `This command find user: Keeper client findUser --email=<email>.`,

		Run: func(cmd *cobra.Command, args []string) {

			emailUser, err := email.New(userArgs.Email)
			if err != nil {
				logger.Log.Info("email input fields", zap.Error(err))
				return
			}

			userFind, err := d.ucUser.Find(context.Background(), emailUser.Email())
			if err != nil {
				logger.Log.Info("user search by email", zap.Error(err))
				return
			}

			msg := fmt.Sprintf("ID: %s\nName: %s\nEmail: %s\nRole: %s\nPassword: %s\nCreatedAt: %s\nUpdatedAt: %s\n",
				userFind.ID().String(),
				userFind.Name().String(),
				userFind.Email().String(),
				userFind.Role().String(),
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

func initFindUserArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&userArgs.Email, "email", "e", "", "user email value.")
	err := cmd.MarkFlagRequired("email")
	if err != nil {
		return
	}
}
