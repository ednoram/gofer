package client

import (
	"encoding/json"
	"gofer/schemas"
	"io"
	"log"
	"net/http"

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
			params["completed"] = "0"
		}

		resp, err := sendApiRequest("GET", "/tasks", nil, params)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Unsuccessful response: %v", string(bodyBytes))
		}

		var tasks []schemas.TaskResponse
		err = json.Unmarshal(bodyBytes, &tasks)
		if err != nil {
			log.Fatalf("Error parsing response: %v", err)
		}

		printTasks(tasks)
	},
}

func Init() *cobra.Command {
	// Commands
	rootCmd.AddCommand(listTasksCommand)
	// Flags
	listTasksCommand.Flags().Bool("all", false, "--all")

	return rootCmd
}
