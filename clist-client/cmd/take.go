package cmd

import (
	"clist-client/client"

	"github.com/spf13/cobra"
)

// takeCmd represents the take command.
// the take command is similar to assign, but automatically assigns the ticket to yourself, rather than another user.
var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client.Take(args[0])
	},
}

func init() {
	rootCmd.AddCommand(takeCmd)
}
