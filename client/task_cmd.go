package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gofer/schemas"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
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

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatal(err)
		}
		if title == "" {
			log.Fatal("Title is required")
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}

		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			log.Fatal(err)
		}

		createTaskData := schemas.CreateTask{
			Title:       title,
			Description: description,
			Completed:   completed,
		}
		body, err := json.Marshal(createTaskData)
		if err != nil {
			log.Fatalf("Error marshaling data: %v", err)
		}

		resp, err := sendApiRequest("POST", "/tasks", body, nil)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusCreated {
			log.Fatalf("Unsuccessful response - status_code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}

		// Print result
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
			log.Fatalf("Error marshaling response: %v", err)
		}

		fmt.Println(prettyJSON.String())
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := cmd.Flags().GetInt("taskId")
		if err != nil {
			log.Fatal(err)
		}
		if taskId <= 0 {
			log.Fatal("Task ID must be a positive integer")
		}

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatal(err)
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}

		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			log.Fatal(err)
		}

		// Check explicitly set flags
		updateTaskData := schemas.UpdateTask{}
		if cmd.Flags().Lookup("title").Changed {
			updateTaskData.Title = &title
		}
		if cmd.Flags().Lookup("description").Changed {
			updateTaskData.Description = &description
		}
		if cmd.Flags().Lookup("completed").Changed {
			updateTaskData.Completed = &completed
		}

		body, err := json.Marshal(updateTaskData)
		if err != nil {
			log.Fatalf("Error marshaling data: %v", err)
		}

		endpoint := fmt.Sprintf("/tasks/%d", taskId)
		resp, err := sendApiRequest("PATCH", endpoint, body, nil)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Unsuccessful response - status_code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}

		fmt.Println("Task updated")
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := cmd.Flags().GetInt("taskId")
		if err != nil {
			log.Fatal(err)
		}
		if taskId <= 0 {
			log.Fatal("Task ID must be a positive integer")
		}

		endpoint := fmt.Sprintf("/tasks/%d", taskId)
		resp, err := sendApiRequest("DELETE", endpoint, nil, nil)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Unsuccessful response - status_code: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}

		fmt.Println("Task deleted")
	},
}

func initTaskCmd() *cobra.Command {
	// List tasks command
	taskCmd.AddCommand(listTasksCmd)
	listTasksCmd.Flags().Bool("all", false, "--all")

	// Add task command
	taskCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().String("title", "", "--title")
	addTaskCmd.Flags().String("description", "", "--description")
	addTaskCmd.Flags().Bool("completed", false, "--completed")

	// Update task command
	taskCmd.AddCommand(updateTaskCmd)
	updateTaskCmd.Flags().Int("taskId", 0, "--taskId")
	updateTaskCmd.Flags().String("title", "", "--title")
	updateTaskCmd.Flags().String("description", "", "--description")
	updateTaskCmd.Flags().Bool("completed", false, "--completed")

	// Delete task command
	taskCmd.AddCommand(deleteTaskCmd)
	deleteTaskCmd.Flags().Int("taskId", 0, "--taskId")

	return taskCmd
}
