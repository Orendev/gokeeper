package cli

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/client/useCase/client"
	"github.com/Orendev/gokeeper/internal/app/client/useCase/storage"
	"github.com/Orendev/gokeeper/internal/pkg/useCase"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/spf13/cobra"
)

type Delivery struct {
	ucUserStorage storage.User
	ucUserClient  client.User

	ucAccountStorage storage.Account
	ucAccountClient  client.Account

	ucTextStorage useCase.Text
	ucTextClient  useCase.Text

	userID *string
	enc    *encryption.Enc

	rootCmd *cobra.Command
}

var version = "0.0.1"

func New(
	ucUserStorage storage.User,
	ucUserClient client.User,
	ucAccountStorage storage.Account,
	ucAccountClient client.Account,
	ucTextStorage useCase.Text,
	ucTextClient useCase.Text,
	key string,
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
		ucTextStorage:    ucTextStorage,
		ucTextClient:     ucTextClient,
		rootCmd:          rootCmd,
	}

	registerUser := d.registerUser()
	loginUser := d.loginUser()
	getUser := d.getUser()

	rootCmd.AddCommand(registerUser)
	rootCmd.AddCommand(loginUser)
	rootCmd.AddCommand(getUser)
	initRegisterUserArgs(registerUser)
	initLoginUserArgs(loginUser)

	createAccount := d.createAccount()
	updateAccount := d.updateAccount()
	deleteAccount := d.deleteAccount()
	listAccount := d.listAccount()

	rootCmd.AddCommand(createAccount)
	rootCmd.AddCommand(updateAccount)
	rootCmd.AddCommand(deleteAccount)
	rootCmd.AddCommand(listAccount)
	initCreateAccountArgs(createAccount)
	initUpdateAccountArgs(updateAccount)
	initDeleteAccountArgs(deleteAccount)
	initListAccountArgs(listAccount)

	createText := d.createText()
	updateText := d.updateText()
	deleteText := d.deleteText()
	listText := d.listText()
	rootCmd.AddCommand(createText)
	rootCmd.AddCommand(updateText)
	rootCmd.AddCommand(deleteText)
	rootCmd.AddCommand(listText)
	initCreateTextArgs(createText)
	initUpdateTextArgs(updateText)
	initDeleteTextArgs(deleteText)
	initListTextArgs(listText)

	user, err := d.ucUserStorage.Get(context.Background())
	if err == nil {
		d.ucUserClient.SetToken(*user)
		key = user.ID().String()
		userID := user.ID().String()

		d.userID = &userID
	}

	d.enc = encryption.New(key)

	return d
}

func (d *Delivery) Run() error {

	err := d.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
