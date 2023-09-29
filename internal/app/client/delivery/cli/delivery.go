package cli

import (
	"github.com/Orendev/gokeeper/internal/app/client/useCase/client"
	"github.com/Orendev/gokeeper/internal/app/client/useCase/storage"
	"github.com/spf13/cobra"
)

type Delivery struct {
	ucUserStorage storage.User
	ucUserClient  client.User

	ucAccountStorage storage.Account
	ucAccountClient  client.Account

	rootCmd *cobra.Command
}

var version = "0.0.1"

func New(
	ucUserStorage storage.User,
	ucUserClient client.User,
	ucAccountStorage storage.Account,
	ucAccountClient client.Account,
) *Delivery {

	rootCmd := &cobra.Command{
		Use:     "keeper",
		Version: version,
		Short:   "Keeper - A simple CLI is a service for storing and protecting your important data",
		Long:    `It can be used to store and protect your important data`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var d = &Delivery{
		ucUserStorage:    ucUserStorage,
		ucUserClient:     ucUserClient,
		ucAccountStorage: ucAccountStorage,
		ucAccountClient:  ucAccountClient,
		rootCmd:          rootCmd,
	}

	registerUser := d.registerUser()
	loginUser := d.loginUser()
	getUser := d.getUser()

	rootCmd.AddCommand(registerUser)
	initRegisterUserArgs(registerUser)

	rootCmd.AddCommand(loginUser)
	initLoginUserArgs(loginUser)

	rootCmd.AddCommand(getUser)

	createAccount := d.createAccount()
	rootCmd.AddCommand(createAccount)
	initCreateAccountArgs(createAccount)

	return d
}

func (d *Delivery) Run() error {
	err := d.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
