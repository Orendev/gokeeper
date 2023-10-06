package cli

import (
	"context"
	"fmt"

	"github.com/Orendev/gokeeper/internal/app/client/delivery/cli/binary"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/spf13/cobra"
)

var createBinaryArgs binary.CreateBinaryArgs
var updateBinaryArgs binary.UpdateBinaryArgs
var deleteBinaryArgs binary.DeleteBinaryArgs
var listBinaryArgs binary.ListBinaryArgs

func (d *Delivery) createBinary() *cobra.Command {

	return &cobra.Command{
		Use:     "createBinary",
		Aliases: []string{"createBin"},
		Short:   "Create new text data in the service.",
		Long:    `This command create a new text data: Keeper client createBinary --title=<title> --data=<binary> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			createBinaryArgs.UserID = *d.userID
			dBinary, err := binary.ToEncCreateBinary(d.enc, &createBinaryArgs)
			if err != nil {
				fmt.Printf("error when encrypting the text data: %s\n", err.Error())
				return
			}

			bin, err := d.ucBinaryStorage.Create(ctx, dBinary)
			if err != nil {
				fmt.Printf("Error create text data: %s\n", err.Error())
				return
			}

			res, err := d.ucBinaryClient.Create(ctx, bin)
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

func (d *Delivery) updateBinary() *cobra.Command {

	return &cobra.Command{
		Use:     "updateBinary",
		Aliases: []string{"updateBin"},
		Short:   "Update a binary data in the service.",
		Long:    `This command update a account: Keeper client updateBinary --id=<uuid> --title=<title> --data=<binary> --comment=<comment>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			updateBinaryArgs.UserID = *d.userID

			dbinary, err := binary.ToEncUpdateBinary(d.enc, &updateBinaryArgs)
			if err != nil {
				fmt.Printf("error when encrypting the binary data: %s\n", err.Error())
				return
			}

			bin, err := d.ucBinaryStorage.Update(ctx, dbinary)
			if err != nil {
				fmt.Printf("Error update binary data: %s\n", err.Error())
				return
			}

			res, err := d.ucBinaryClient.Update(ctx, bin)
			if err != nil {
				fmt.Printf("Error update an binary data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The binary data with ID updated: %s\n",
				res.ID().String(),
			)

		},
	}
}

func (d *Delivery) deleteBinary() *cobra.Command {

	return &cobra.Command{
		Use:     "deleteBinary",
		Aliases: []string{"deleteBin"},
		Short:   "Delete binary data in the service.",
		Long:    `This command delete a binary data: Keeper client deleteBinary --id=<id>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			id := converter.StringToUUID(deleteBinaryArgs.ID)

			err := d.ucBinaryStorage.Delete(ctx, id)
			if err != nil {
				fmt.Printf("Error delete binary data: %s\n", err.Error())
				return
			}

			err = d.ucBinaryClient.Delete(ctx, id)
			if err != nil {

				fmt.Printf("Error delete an binary data on a remote server: %s\n", err.Error())
				return
			}

			fmt.Printf("The binary data with ID delete: %s\n",
				id.String(),
			)

		},
	}
}

func (d *Delivery) listBinary() *cobra.Command {

	return &cobra.Command{
		Use:     "listBinary",
		Aliases: []string{"listBin"},
		Short:   "List binary data in the service.",
		Long:    `This command list a binary data: Keeper client listBinary --limit=<10> --offset=<0>.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			var parameter queryParameter.QueryParameter
			parameter.Pagination.Limit = listAccountArgs.Limit
			parameter.Pagination.Offset = listAccountArgs.Offset

			list, err := d.ucBinaryClient.List(ctx, parameter)

			if err != nil {
				fmt.Printf("Error list binary data: %s\n", err.Error())
				return
			}

			for _, value := range list.Data {
				val, err := binary.ToDecBinary(d.enc, value)
				if err != nil {
					fmt.Printf("Error list binary data: %s\n", err.Error())
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

func initCreateBinaryArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&createBinaryArgs.Title, "title", "", "", "text title value.")
	cmd.Flags().StringVarP(&createBinaryArgs.Data, "data", "", "", "text data value.")
	cmd.Flags().StringVarP(&createBinaryArgs.Comment, "comment", "", "", "text comment value.")

	err := cmd.MarkFlagRequired("data")
	if err != nil {
		return
	}

}

func initUpdateBinaryArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&updateBinaryArgs.ID, "id", "", "", "text id value.")
	cmd.Flags().StringVarP(&updateBinaryArgs.Title, "title", "", "", "text title value.")
	cmd.Flags().StringVarP(&updateBinaryArgs.Data, "data", "", "", "text data value.")
	cmd.Flags().StringVarP(&updateBinaryArgs.Comment, "comment", "", "", "text comment value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}

	err = cmd.MarkFlagRequired("data")
	if err != nil {
		return
	}

}

func initDeleteBinaryArgs(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&deleteBinaryArgs.ID, "id", "", "", "text id value.")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		return
	}
}

func initListBinaryArgs(cmd *cobra.Command) {
	cmd.Flags().Uint64VarP(&listBinaryArgs.Limit, "limit", "", 10, "text list limit.")
	cmd.Flags().Uint64VarP(&listBinaryArgs.Offset, "offset", "", 0, "offset of the text list.")
}
