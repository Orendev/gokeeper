package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/account"
	domainAccount "github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
	"github.com/spf13/cobra"
)

var createAccountArgs account.CreateAccountArgs
var updateAccountArgs account.UpdateAccountArgs
var deleteAccountArgs account.DeleteAccountArgs
var listAccountArgs account.ListAccountArgs

func (d *Delivery) createAccount() *cobra.Command {

	return &cobra.Command{
		Use:     "createAccount",
		Aliases: []string{"createAc"},
		Short:   "Create new account in the service.",
		Long:    `This command create a new account: Keeper client createAccount --title=<title> --login=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			err = account.ToEncAccountArgs[account.CreateAccountArgs](d.enc, &createAccountArgs)
			if err != nil {
				fmt.Printf("error when encrypting the account: %s\n", err.Error())
				return
			}

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
			if err != nil {
				fmt.Printf("Error create domainAccount: %s\n", err.Error())
				return
			}

			ac, err := d.ucAccountStorage.Create(ctx, *dAccount)
			if err != nil {
				fmt.Printf("Error create account: %s\n", err.Error())
				return
			}

			id, err := d.ucAccountClient.Create(ctx, *ac)
			if err != nil {
				fmt.Printf("Error creating an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The login/password pair of the account has been created with an ID: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) updateAccount() *cobra.Command {

	return &cobra.Command{
		Use:     "updateAccount",
		Aliases: []string{"updateAc"},
		Short:   "Update a account in the service.",
		Long:    `This command update a account: Keeper client createAccount --id=<uuid> --title=<title> --login=<dev@email.com> --password=<password>.`,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			err = account.ToEncAccountArgs[account.UpdateAccountArgs](d.enc, &updateAccountArgs)
			if err != nil {
				fmt.Printf("error when encrypting the account: %s\n", err.Error())
				return
			}

			ctx := context.Background()
			titleObj, err := title.New(updateAccountArgs.Title)
			if err != nil {
				fmt.Printf("Error title input fields: %s\n", err.Error())
				return
			}

			loginObj, err := login.New(updateAccountArgs.Login)
			if err != nil {
				fmt.Printf("Error login input fields: %s\n", err.Error())
				return
			}

			passwordObj, err := password.New(updateAccountArgs.Password)
			if err != nil {
				fmt.Printf("Error password input fields: %s\n", err.Error())
				return
			}

			urlObj, err := url.New(updateAccountArgs.URL)
			if err != nil {
				fmt.Printf("Error url input fields: %s\n", err.Error())
				return
			}

			commentObj, err := comment.New(updateAccountArgs.Comment)
			if err != nil {
				fmt.Printf("Error comment input fields: %s\n", err.Error())
				return
			}

			dAccount, err := domainAccount.NewWithID(
				converter.StringToUUID(updateAccountArgs.ID),
				*titleObj,
				*loginObj,
				*passwordObj,
				*urlObj,
				*commentObj,
				false,
				time.Now().UTC(),
				time.Now().UTC(),
			)
			if err != nil {
				fmt.Printf("Error update domainAccount: %s\n", err.Error())
				return
			}

			ac, err := d.ucAccountStorage.Update(ctx, *dAccount)
			if err != nil {
				fmt.Printf("Error create account: %s\n", err.Error())
				return
			}

			id, err := d.ucAccountClient.Update(ctx, *ac)
			if err != nil {
				fmt.Printf("Error update an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The login/password pair of the account has been created with an ID: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) deleteAccount() *cobra.Command {

	return &cobra.Command{
		Use:     "deleteAccount",
		Aliases: []string{"deleteAc"},
		Short:   "Delete account in the service.",
		Long:    `This command delete a account: Keeper client deleteAccount --id=<id>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			id := converter.StringToUUID(deleteAccountArgs.ID)

			err := d.ucAccountStorage.Delete(ctx, id)
			if err != nil {
				fmt.Printf("Error delete account: %s\n", err.Error())
				return
			}

			err = d.ucAccountClient.Delete(ctx, id)
			if err != nil {
				fmt.Printf("Error delete an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The login/password pair of the account has been delete with an ID: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) listAccount() *cobra.Command {

	return &cobra.Command{
		Use:     "listAccount",
		Aliases: []string{"listAc"},
		Short:   "List account in the service.",
		Long:    `This command list a account: Keeper client listAccount --limit=<10> --offset=<10>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			var parameter queryParameter.QueryParameter
			parameter.Pagination.Limit = listAccountArgs.Limit
			parameter.Pagination.Offset = listAccountArgs.Offset

			accounts, err := d.ucAccountClient.List(ctx, parameter)
			if err != nil {
				fmt.Printf("Error list account: %s\n", err.Error())
				return
			}

			for _, value := range accounts {
				val, err := account.ToDecAccount(d.enc, value)
				if err != nil {
					fmt.Printf("Error list account: %s\n", err.Error())
					return
				}

				msg := fmt.Sprintf("ID: %s\nTitle: %s\nURL: %s\nComment: %s\nLogin: %s\nPassword: %s\nCreatedAt: %s\nUpdatedAt: %s\nIsDeleted: %v\n",
					val.ID().String(),
					val.Title().String(),
					val.URL().String(),
					val.Comment().String(),
					val.Login().String(),
					val.Password().String(),
					val.CreatedAt().String(),
					val.UpdatedAt().String(),
					val.IsDeleted(),
				)

				fmt.Println(msg)
			}

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

func initUpdateAccountArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&updateAccountArgs.ID, "id", "", "", "account id value.")
	cmd.Flags().StringVarP(&updateAccountArgs.Title, "title", "", "", "account title value.")
	cmd.Flags().StringVarP(&updateAccountArgs.Login, "login", "", "", "account login value.")
	cmd.Flags().StringVarP(&updateAccountArgs.Password, "password", "", "", "account password value.")
	cmd.Flags().StringVarP(&updateAccountArgs.Comment, "comment", "", "", "account comment value.")
	cmd.Flags().StringVarP(&updateAccountArgs.URL, "url", "u", "", "account url value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("login")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("password")
	if err != nil {
		return
	}
}

func initDeleteAccountArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&deleteAccountArgs.ID, "id", "", "", "account id value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}
}

func initListAccountArgs(cmd *cobra.Command) {
	cmd.Flags().Uint64VarP(&listAccountArgs.Limit, "limit", "", 10, "account list limit.")
	cmd.Flags().Uint64VarP(&listAccountArgs.Offset, "offset", "", 0, "offset of the account list.")
}
