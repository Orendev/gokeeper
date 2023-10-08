package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/card"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/spf13/cobra"
)

var createCardArgs card.CreateCardArgs
var updateCardArgs card.UpdateCardArgs
var deleteCardArgs card.DeleteCardArgs
var listCardArgs card.ListCardArgs

func (d *Delivery) createCard() *cobra.Command {

	return &cobra.Command{
		Use:     "createCard",
		Aliases: []string{"createCar"},
		Short:   "Create new card data in the service.",
		Long:    `This command create a new card data: Keeper client createCard --name=<title> --number=<card> --number=<date> --number=<cvc> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}
			createCardArgs.UserID = d.userID
			dCard, err := card.ToEncCreateCard(d.enc, &createCardArgs)
			if err != nil {
				fmt.Printf("error when encrypting the card data: %s\n", err.Error())
				return
			}

			bin, err := d.ucCardStorage.Create(ctx, dCard)
			if err != nil {
				fmt.Printf("Error create card data: %s\n", err.Error())
				return
			}

			res, err := d.ucCardClient.Create(ctx, bin)
			if err != nil {
				fmt.Printf("Error creating an card data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The card data with ID created: %s\n",
				res.ID().String(),
			)

		},
	}
}

func (d *Delivery) updateCard() *cobra.Command {

	return &cobra.Command{
		Use:     "updateCard",
		Aliases: []string{"updateCar"},
		Short:   "Update a card data in the service.",
		Long:    `This command update a account: Keeper client updateCard --id=<uuid> --title=<title> --data=<card> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}
			updateCardArgs.UserID = d.userID

			dcard, err := card.ToEncUpdateCard(d.enc, &updateCardArgs)
			if err != nil {
				fmt.Printf("error when encrypting the card data: %s\n", err.Error())
				return
			}

			bin, err := d.ucCardStorage.Update(ctx, dcard)
			if err != nil {
				fmt.Printf("Error update card data: %s\n", err.Error())
				return
			}

			res, err := d.ucCardClient.Update(ctx, bin)
			if err != nil {
				fmt.Printf("Error update an card data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The card data with ID updated: %s\n",
				res.ID().String(),
			)

		},
	}
}

func (d *Delivery) deleteCard() *cobra.Command {

	return &cobra.Command{
		Use:     "deleteCard",
		Aliases: []string{"deleteCar"},
		Short:   "Delete card data in the service.",
		Long:    `This command delete a card data: Keeper client deleteCard --id=<id>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := d.Init()
			if err != nil {
				fmt.Printf("failed to init client: %s\n", err.Error())
				return
			}
			id := converter.StringToUUID(deleteCardArgs.ID)

			err = d.ucCardStorage.Delete(ctx, id)
			if err != nil {
				fmt.Printf("Error delete card data: %s\n", err.Error())
				return
			}

			err = d.ucCardClient.Delete(ctx, id)
			if err != nil {

				fmt.Printf("Error delete an card data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The card data with ID delete: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) listCard() *cobra.Command {

	return &cobra.Command{
		Use:     "listCard",
		Aliases: []string{"listBin"},
		Short:   "List card data in the service.",
		Long:    `This command list a card data: Keeper client listCard --limit=<10> --offset=<0>.`,
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

			list, err := d.ucCardClient.List(ctx, parameter)

			if err != nil {
				fmt.Printf("Error list card data: %s\n", err.Error())
				return
			}

			for _, value := range list.Data {
				val, err := card.ToDecCard(d.enc, value)
				if err != nil {
					fmt.Printf("Error list card data: %s\n", err.Error())
					return
				}

				msg := fmt.Sprintf("ID: %s\nCardName: %s\nCardNumber: %s\nCardDate: %s\nCVC: %s\nComment: %s\nCreatedAt: %s\nUpdatedAt: %s\nIsDeleted: %v\n",
					val.ID().String(),
					string(val.CardName()),
					string(val.CardNumber()),
					string(val.CardDate()),
					string(val.CVC()),
					string(val.Comment()),
					val.CreatedAt().String(),
					val.UpdatedAt().String(),
					val.IsDeleted(),
				)

				fmt.Println(msg)
			}

			msg := fmt.Sprintf("Limit: %v\nOffset: %v\nTotal: %v\n",
				list.Limit,
				list.Offset,
				list.Total,
			)

			fmt.Println(msg)

		},
	}
}

func initCreateCardArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&createCardArgs.CardName, "name", "", "", "card title value.")
	cmd.Flags().StringVarP(&createCardArgs.CardNumber, "number", "", "", "card name value.")
	cmd.Flags().StringVarP(&createCardArgs.CardDate, "date", "", "", "card date value.")
	cmd.Flags().StringVarP(&createCardArgs.CVC, "cvc", "", "", "card cvc value.")
	cmd.Flags().StringVarP(&createCardArgs.Comment, "comment", "", "", "card comment value.")

	err := cmd.MarkFlagRequired("name")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("number")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("date")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("cvc")
	if err != nil {
		return
	}
}

func initUpdateCardArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&updateCardArgs.ID, "id", "", "", "card id value.")
	cmd.Flags().StringVarP(&updateCardArgs.CardName, "name", "", "", "card name value.")
	cmd.Flags().StringVarP(&updateCardArgs.CardNumber, "number", "", "", "card number value.")
	cmd.Flags().StringVarP(&updateCardArgs.CardDate, "date", "", "", "card date value.")
	cmd.Flags().StringVarP(&updateCardArgs.CVC, "cvc", "", "", "card cvc value.")
	cmd.Flags().StringVarP(&updateCardArgs.Comment, "comment", "", "", "card comment value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("number")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("date")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("cvc")
	if err != nil {
		return
	}

}

func initDeleteCardArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&deleteCardArgs.ID, "id", "", "", "card id value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}
}

func initListCardArgs(cmd *cobra.Command) {
	cmd.Flags().Uint64VarP(&listCardArgs.Limit, "limit", "", 10, "card list limit.")
	cmd.Flags().Uint64VarP(&listCardArgs.Offset, "offset", "", 0, "offset of the card list.")
}
