package client

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var ListTasksCommand = &cobra.Command{
	Use:   "task list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := SendApiRequest("GET", "/tasks", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(body)
	},
}
