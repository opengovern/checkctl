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
	controlsCmd.PersistentFlags().Int("page-size", 25, "Defines page size of response")
	controlsCmd.PersistentFlags().Int("page-number", 1, "Defines page number of response")

	GetCmd.AddCommand(benchmarksCmd)
	benchmarksCmd.PersistentFlags().Int("page-size", 25, "Defines page size of response")
	benchmarksCmd.PersistentFlags().Int("page-number", 1, "Defines page number of response")

	GetCmd.AddCommand(complianceSummaryForBenchmarkCmd)
	complianceSummaryForBenchmarkCmd.PersistentFlags().StringSlice("benchmark-ids", []string{}, "List of Benchmark IDs to get the summary for (optional)")
	complianceSummaryForBenchmarkCmd.PersistentFlags().Bool("is-root", true, "Whether to return only root benchmarks or not. (matters if benchmark-id list not provided)")

	GetCmd.AddCommand(complianceSummaryForIntegrationCmd)
	complianceSummaryForIntegrationCmd.PersistentFlags().String("integration", "", "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	complianceSummaryForIntegrationCmd.PersistentFlags().String("benchmark-id", "", "Benchmark ID")
}
