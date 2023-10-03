package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/text"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/spf13/cobra"
)

var createTextArgs text.CreateTextArgs
var updateTextArgs text.UpdateTextArgs
var deleteTextArgs text.DeleteTextArgs
var listTextArgs text.ListTextArgs

func (d *Delivery) createText() *cobra.Command {

	return &cobra.Command{
		Use:     "createText",
		Aliases: []string{"createTex"},
		Short:   "Create new text data in the service.",
		Long:    `This command create a new text data: Keeper client createText --title=<title> --data=<text text> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			createTextArgs.UserID = *d.userID
			dText, err := text.ToEncCreateText(d.enc, &createTextArgs)
			if err != nil {
				fmt.Printf("error when encrypting the text data: %s\n", err.Error())
				return
			}

			tex, err := d.ucTextStorage.Create(ctx, dText)
			if err != nil {
				fmt.Printf("Error create text data: %s\n", err.Error())
				return
			}

			res, err := d.ucTextClient.Create(ctx, tex)
			if err != nil {
				fmt.Printf("Error creating an text data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The text data with ID created: %s\n",
				res.ID().String(),
			)

		},
	}
}

func (d *Delivery) updateText() *cobra.Command {

	return &cobra.Command{
		Use:     "updateText",
		Aliases: []string{"updateTex"},
		Short:   "Update a text data in the service.",
		Long:    `This command update a account: Keeper client updateText --id=<uuid> --title=<title> --data=<text text> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			updateTextArgs.UserID = *d.userID

			dText, err := text.ToEncUpdateText(d.enc, &updateTextArgs)
			if err != nil {
				fmt.Printf("error when encrypting the text data: %s\n", err.Error())
				return
			}

			tex, err := d.ucTextStorage.Update(ctx, dText)
			if err != nil {
				fmt.Printf("Error update text data: %s\n", err.Error())
				return
			}

			res, err := d.ucTextClient.Update(ctx, tex)
			if err != nil {
				fmt.Printf("Error update an text data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The text data with ID updated: %s\n",
				res.ID().String(),
			)

		},
	}
}

func (d *Delivery) deleteText() *cobra.Command {

	return &cobra.Command{
		Use:     "deleteText",
		Aliases: []string{"deleteTex"},
		Short:   "Delete text data in the service.",
		Long:    `This command delete a text data: Keeper client deleteText --id=<id>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			id := converter.StringToUUID(deleteTextArgs.ID)

			err := d.ucTextStorage.Delete(ctx, id)
			if err != nil {
				fmt.Printf("Error delete text data: %s\n", err.Error())
				return
			}

			err = d.ucTextClient.Delete(ctx, id)
			if err != nil {

				fmt.Printf("Error delete an text data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The text data with ID delete: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) listText() *cobra.Command {

	return &cobra.Command{
		Use:     "listText",
		Aliases: []string{"listTex"},
		Short:   "List text data in the service.",
		Long:    `This command list a text data: Keeper client listText --limit=<10> --offset=<0>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			var parameter queryParameter.QueryParameter
			parameter.Pagination.Limit = listAccountArgs.Limit
			parameter.Pagination.Offset = listAccountArgs.Offset

			list, err := d.ucTextClient.List(ctx, parameter)

			if err != nil {
				fmt.Printf("Error list text data: %s\n", err.Error())
				return
			}

			for _, value := range list.Data {
				val, err := text.ToDecText(d.enc, value)
				if err != nil {
					fmt.Printf("Error list text data: %s\n", err.Error())
					return
				}

				msg := fmt.Sprintf("ID: %s\nTitle: %s\nData: %s\nComment: %s\nCreatedAt: %s\nUpdatedAt: %s\nIsDeleted: %v\n",
					val.ID().String(),
					val.Title().String(),
					string(val.Data()),
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

func initCreateTextArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&createTextArgs.Title, "title", "", "", "text title value.")
	cmd.Flags().StringVarP(&createTextArgs.Data, "data", "", "", "text data value.")
	cmd.Flags().StringVarP(&createTextArgs.Comment, "comment", "", "", "text comment value.")

	err := cmd.MarkFlagRequired("data")
	if err != nil {
		return
	}

}

func initUpdateTextArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&updateTextArgs.ID, "id", "", "", "text id value.")
	cmd.Flags().StringVarP(&updateTextArgs.Title, "title", "", "", "text title value.")
	cmd.Flags().StringVarP(&updateTextArgs.Data, "data", "", "", "text data value.")
	cmd.Flags().StringVarP(&updateTextArgs.Comment, "comment", "", "", "text comment value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("data")
	if err != nil {
		return
	}

}

func initDeleteTextArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&deleteTextArgs.ID, "id", "", "", "text id value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}
}

func initListTextArgs(cmd *cobra.Command) {
	cmd.Flags().Uint64VarP(&listTextArgs.Limit, "limit", "", 10, "text list limit.")
	cmd.Flags().Uint64VarP(&listTextArgs.Offset, "offset", "", 0, "offset of the text list.")
}
