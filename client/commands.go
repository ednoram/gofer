package client

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofer",
	Short: "Gofer",
	Long:  "Gofer - a simple task management tool",
}

var listTasksCommand = &cobra.Command{
	Use:   "task list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		allFlag, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}
		params := make(map[string]string)
		if !allFlag {
			params["completed"] = "1"
		}
		body, err := SendApiRequest("GET", "/tasks", nil, params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(body)
	},
}

func Init() *cobra.Command {
	// Commands
	rootCmd.AddCommand(listTasksCommand)
	// Flags
	listTasksCommand.Flags().Bool("all", false, "--all")

	return rootCmd
}
