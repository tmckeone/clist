package cmd

import (
	"clist-client/client"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view [ticketId]",
	Short: "View a ticket",
	Long: `The clist view command allows you to view a ticket's contents.
	Usage:

	clist view [ticketId] - Returns the contents of the specified ticket id.

	If you do not know the ticket id you are looking for, try using the search command.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		mine, _ := cmd.Flags().GetBool("mine")

		//if the -m or --mine flag is present, run the ViewMyTickets function and exit.
		if mine {
			client.ViewMyTickets()
			return
		}

		//if no flags are present, view the ticket number provided.
		client.ViewTicket(args[0])
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	//Define flags for the view command
	viewCmd.Flags().BoolP("mine", "m", false, "Returns a list of all of the tickets assigned to you")
}
