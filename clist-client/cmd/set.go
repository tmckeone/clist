package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set variables for clist",
	Long: `The clist-client set command allows you to change variables for usage of the client.

-a address - This defines the ip address / hostname of the clist server your client will connect to.
-p port - This defines the port of the clist server your client will connect to.

-u username - This defines the username clist will use to sign into the clist server.
-pw password - This defines the password clist will use to sign into the clist server.

These variables can also be assigned in the .clist-client configuration file. Use clist --config to locate the configuration file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
