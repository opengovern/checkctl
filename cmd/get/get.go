/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: checkctl get controls|benchmarks --page-size")
	},
}

func init() {

	GetCmd.AddCommand(controlsCmd)
	GetCmd.AddCommand(benchmarksCmd)

	GetCmd.PersistentFlags().Int("page-size", 25, "Defines page size of response")
	GetCmd.PersistentFlags().Int("page-number", 1, "Defines page number of response")

}
