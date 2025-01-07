package main

import (
	"gofer/client"
	"log"
)

func main() {
	rootCmd := client.Init()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
