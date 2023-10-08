package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/account"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
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
			ctx := context.Background()
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}

			createAccountArgs.UserID = d.userID

			dAccount, err := account.ToEncCreateAccount(d.enc, &createAccountArgs)
			if err != nil {
				fmt.Printf("error when encrypting the account: %s\n", err.Error())
				return
			}

			acInternal, err := d.ucAccountStorage.Create(ctx, dAccount)
			if err != nil {
				fmt.Printf("Error create account: %s\n", err.Error())
				return
			}

			acExternal, err := d.ucAccountClient.Create(ctx, acInternal)
			if err != nil {
				fmt.Printf("Error creating an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The login/password pair of the account has been created with an ID: %s\n",
				acExternal.ID().String(),
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
			ctx := context.Background()
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}

			updateAccountArgs.UserID = d.userID
			dAccount, err := account.ToEncUpdateAccount(d.enc, &updateAccountArgs)
			if err != nil {
				fmt.Printf("error when encrypting the account: %s\n", err.Error())
				return
			}

			if err != nil {
				fmt.Printf("Error update domainAccount: %s\n", err.Error())
				return
			}

			acInternal, err := d.ucAccountStorage.Update(ctx, dAccount)
			if err != nil {
				fmt.Printf("Error create account: %s\n", err.Error())
				return
			}

			acExternal, err := d.ucAccountClient.Update(ctx, acInternal)
			if err != nil {
				fmt.Printf("Error update an account on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The login/password pair of the account has been created with an ID: %s\n",
				acExternal.ID().String(),
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
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}
			id := converter.StringToUUID(deleteAccountArgs.ID)

			err = d.ucAccountStorage.Delete(ctx, id)
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
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}

			var parameter queryParameter.QueryParameter
			parameter.Pagination.Limit = listAccountArgs.Limit
			parameter.Pagination.Offset = listAccountArgs.Offset

			list, err := d.ucAccountClient.List(ctx, parameter)
			if err != nil {
				fmt.Printf("Error list account: %s\n", err.Error())
				return
			}

			for _, value := range list.Data {
				val, err := account.ToDecAccount(d.enc, value)
				if err != nil {
					fmt.Printf("Error list account: %s\n", err.Error())
					return
				}

				msg := fmt.Sprintf("ID: %s\nTitle: %s\nURL: %s\nComment: %s\nLogin: %s\nPassword: %s\nCreatedAt: %s\nUpdatedAt: %s\nIsDeleted: %v\n",
					val.ID().String(),
					val.Title().String(),
					string(val.URL()),
					string(val.Comment()),
					string(val.Login()),
					string(val.Password()),
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
