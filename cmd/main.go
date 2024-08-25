package main

import (
	"core/internal/server"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ServerCore",
		Short: "ServerCore CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			// Default action: Start the server
			server.Init()
		},
	}

	// Server command
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Init()
		},
	}

	// Admin command
	var adminCmd = &cobra.Command{
		Use:   "admin",
		Short: "Run the admin interface",
		Run: func(cmd *cobra.Command, args []string) {
			startAdmin()
		},
	}

	// Add other commands here
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(adminCmd)
	// Add more subcommands as needed

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startAdmin() {
	// Admin interface logic
}
