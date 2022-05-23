package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IPtracker CLI app 🫠",
		Long:  `IPTracker CLI app 🫠`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
