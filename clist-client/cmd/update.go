package cmd

import (
	"clist-client/client"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a ticket",
	Long: `
	
	Usage:

	clist update ticketId -t tag - Allows you to change the tag of the ticket id specified.

	clist update ticketId -c - Closes the specified ticket

	clist update ticketId -o - Opens the specified ticket
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetString("tag")
		close, _ := cmd.Flags().GetBool("close")
		open, _ := cmd.Flags().GetBool("open")

		//look for flags and run the corresponding function.
		if tag != "" {
			client.ChangeTag(args[0], tag)
			return
		}

		if close {
			client.UpdateStatus(args[0], false)
			return
		}

		if open {
			client.UpdateStatus(args[0], true)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	//flag definition for the update command
	updateCmd.Flags().StringP("tag", "t", "", "Change the tag of a ticket")
	updateCmd.Flags().BoolP("close", "c", false, "Close a ticket")
	updateCmd.Flags().BoolP("open", "o", false, "Open a ticket")
}
