package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/account"
	domainAccount "github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
	"github.com/spf13/cobra"
)

var createAccountArgs account.CreateAccountArgs

func (d *Delivery) createAccount() *cobra.Command {

	return &cobra.Command{
		Use:     "createAccount",
		Aliases: []string{"createAc"},
		Short:   "Create new account in the service.",
		Long:    `This command create a new account: Keeper client createAccount --title=<title> --login=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			titleObj, err := title.New(createAccountArgs.Title)
			if err != nil {
				fmt.Printf("Error title input fields: %s\n", err.Error())
				return
			}

			loginObj, err := login.New(createAccountArgs.Login)
			if err != nil {
				fmt.Printf("Error login input fields: %s\n", err.Error())
				return
			}

			passwordObj, err := password.New(createAccountArgs.Password)
			if err != nil {
				fmt.Printf("Error password input fields: %s\n", err.Error())
				return
			}

			urlObj, err := url.New(createAccountArgs.URL)
			if err != nil {
				fmt.Printf("Error url input fields: %s\n", err.Error())
				return
			}

			commentObj, err := comment.New(createAccountArgs.Comment)
			if err != nil {
				fmt.Printf("Error comment input fields: %s\n", err.Error())
				return
			}

			dAccount, err := domainAccount.New(
				*titleObj,
				*loginObj,
				*passwordObj,
				*urlObj,
				*commentObj,
			)

			ac, err := d.ucAccountStorage.Create(ctx, *dAccount)
			if err != nil {
				fmt.Printf("Error create account: %s\n", err.Error())
				return
			}

			accountExternal, err := d.ucAccountClient.Create(ctx, *ac)
			if err != nil {
				fmt.Printf("Error creating an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("ID: %s\nTitle: %s\nLogin: %s\nPassword: %s\nURL: %s\nComment: %s\nCreatedAt: %s\nUpdatedAt: %s\n",
				accountExternal.ID().String(),
				accountExternal.Title().String(),
				accountExternal.Login().String(),
				accountExternal.Password().String(),
				accountExternal.URL().String(),
				accountExternal.Comment().String(),
				accountExternal.CreatedAt().String(),
				accountExternal.UpdatedAt().String(),
			)

		},
	}
}

func initCreateAccountArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&createAccountArgs.Title, "title", "t", "", "account title value.")
	cmd.Flags().StringVarP(&createAccountArgs.Login, "login", "l", "", "account login value.")
	cmd.Flags().StringVarP(&createAccountArgs.Password, "password", "p", "", "account password value.")
	cmd.Flags().StringVarP(&createAccountArgs.Comment, "comment", "c", "", "account comment value.")
	cmd.Flags().StringVarP(&createAccountArgs.URL, "url", "u", "", "account url value.")

	err := cmd.MarkFlagRequired("login")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("password")
	if err != nil {
		return
	}
}
