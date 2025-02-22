package client

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofer",
	Short: "Gofer",
	Long:  "Gofer - a simple task management tool",
}

func Init() *cobra.Command {
	rootCmd.AddCommand(initTaskCmd())

	return rootCmd
}
