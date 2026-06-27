package cmd

import (
	"github.com/spf13/cobra"
)

// mlbCmd represents the mlb command
var mlbCmd = &cobra.Command{
	Use: "mlb",
}

func init() {
	rootCmd.AddCommand(mlbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mlbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mlbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
