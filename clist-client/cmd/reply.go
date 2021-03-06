/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"clist-client/client"

	"github.com/spf13/cobra"
)

// replyCmd represents the reply command
var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Adds a reply to a ticket",
	Long: `
Usage:

clist reply ticketId - Opens a multi-line prompt for you to input your reply. When finished, type ~ and hit enter to save your reply, or ctrl+c to cancel.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client.Reply(args[0])
	},
}

func init() {
	rootCmd.AddCommand(replyCmd)
}
