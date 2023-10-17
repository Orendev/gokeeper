package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/configs"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/useCase"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	memory "github.com/Orendev/gokeeper/pkg/tools/fileStorage"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/spf13/cobra"
)

type Delivery struct {
	ucUserStorage useCase.User
	ucUserClient  useCase.User

	ucAccountStorage useCase.Account
	ucAccountClient  useCase.Account

	ucTextStorage useCase.Text
	ucTextClient  useCase.Text

	ucBinaryStorage useCase.Binary
	ucBinaryClient  useCase.Binary

	ucCardStorage useCase.Card
	ucCardClient  useCase.Card

	userID      string
	enc         *encryption.Enc
	fileStorage *memory.FileStorage
	key         string
	rootCmd     *cobra.Command
}

func New(
	ucUserStorage useCase.User,
	ucUserClient useCase.User,
	ucAccountStorage useCase.Account,
	ucAccountClient useCase.Account,
	ucTextStorage useCase.Text,
	ucTextClient useCase.Text,
	ucBinaryStorage useCase.Binary,
	ucBinaryClient useCase.Binary,

	ucCardStorage useCase.Card,
	ucCardClient useCase.Card,

	fileStorage *memory.FileStorage,
	key string,
) *Delivery {

	rootCmd := &cobra.Command{
		Use:     "keeper",
		Version: fmt.Sprintf("%s\nbuild date %s", configs.BuildVersion, configs.BuildDate),
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
		ucBinaryStorage:  ucBinaryStorage,
		ucBinaryClient:   ucBinaryClient,
		ucCardStorage:    ucCardStorage,
		ucCardClient:     ucCardClient,
		fileStorage:      fileStorage,
		rootCmd:          rootCmd,
	}

	registerUser := d.registerUser()
	loginUser := d.loginUser()

	rootCmd.AddCommand(registerUser)
	rootCmd.AddCommand(loginUser)
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

	createBinary := d.createBinary()
	updateBinary := d.updateBinary()
	deleteBinary := d.deleteBinary()
	listBinary := d.listBinary()
	rootCmd.AddCommand(createBinary)
	rootCmd.AddCommand(updateBinary)
	rootCmd.AddCommand(deleteBinary)
	rootCmd.AddCommand(listBinary)
	initCreateBinaryArgs(createBinary)
	initUpdateBinaryArgs(updateBinary)
	initDeleteBinaryArgs(deleteBinary)
	initListBinaryArgs(listBinary)

	createCard := d.createCard()
	updateCard := d.updateCard()
	deleteCard := d.deleteCard()
	listCard := d.listCard()
	rootCmd.AddCommand(createCard)
	rootCmd.AddCommand(updateCard)
	rootCmd.AddCommand(deleteCard)
	rootCmd.AddCommand(listCard)
	initCreateCardArgs(createCard)
	initUpdateCardArgs(updateCard)
	initDeleteCardArgs(deleteCard)
	initListCardArgs(listCard)

	d.key = key

	return d
}

func (d *Delivery) Run() error {

	err := d.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func (d *Delivery) Init() error {
	ctx := context.Background()

	var parameter queryParameter.QueryParameter
	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	m, err := d.fileStorage.Memory()
	if err != nil {
		return ErrCannotAuthorize
	}

	e, err := email.New(m.Email)
	if err != nil {
		return ErrCannotAuthorize
	}

	p, err := password.New(m.Password)
	if err != nil {
		return ErrCannotAuthorize
	}

	user, err := d.ucUserStorage.Login(ctx, *e, *p)
	if err != nil {
		return repository.ErrNoAuth
	}

	d.ucUserClient.SetToken(ctx, user)
	d.key = user.Password().String()
	d.userID = user.ID().String()

	d.enc = encryption.New(d.key)

	return nil
}
