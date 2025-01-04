package main

import (
	"log"

	"github.com/spf13/cobra"

	"gofer/client"
)

var rootCmd = &cobra.Command{
	Use:   "gofer",
	Short: "Gofer",
	Long:  "Gofer - a simple task management tool",
}

func main() {
	rootCmd.AddCommand(client.ListTasksCommand)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
