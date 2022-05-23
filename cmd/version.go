package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of IP-CLI",
	Long:  `All software has versions. This is IP_CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("IP-CLI v1.0.1")
	},
}
