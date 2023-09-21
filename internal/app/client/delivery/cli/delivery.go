package cli

import (
	"github.com/Orendev/gokeeper/internal/app/client/useCase"
	"github.com/spf13/cobra"
)

type Delivery struct {
	ucUser  useCase.User
	rootCmd *cobra.Command
}

var version = "0.0.1"

func New(
	ucUser useCase.User,
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
		ucUser:  ucUser,
		rootCmd: rootCmd,
	}

	createUser := d.createUser()
	findUser := d.findUser()
	rootCmd.AddCommand(createUser)
	initCreateUserArgs(createUser)
	rootCmd.AddCommand(findUser)
	initFindUserArgs(findUser)

	return d
}

func (d *Delivery) Run() error {
	err := d.rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
