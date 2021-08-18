package cmd

import (
	"clist-client/client"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `
	
	Usage:

	clist search -m matchString - Searches for a tickets with a subject similar to or matching matchString.

	clist search -t tag - Returns a list of open tickets with the matching tag.

	clist search -u username - Returns a list of open tickets assigned to the username specified.


	`,
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetString("tag")
		query, _ := cmd.Flags().GetBool("query")
		user, _ := cmd.Flags().GetString("username")

		// execute a function based on the present tag
		if tag != "" {
			client.SearchTag(tag)
			return
		}

		if query {
			client.SearchQuery()
			return
		}

		if user != "" {
			client.SearchUser(user)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Defining flags for the search command
	searchCmd.Flags().BoolP("query", "q", false, "Search for a subject that contains the match value")
	searchCmd.Flags().StringP("tag", "t", "", "Search for tickets with the tag value")
	searchCmd.Flags().StringP("username", "u", "", "Search for tickts assigned to a user")
}
