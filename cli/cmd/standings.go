package cmd

import (
	"fmt"

	"github.com/gopherlibs/big-league-stats/sdk"
	"github.com/spf13/cobra"
)

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings <league> <division>",
	Short: "Display MLB standings for a division or wildcard.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		// Parse the league argument
		var c *sdk.MLBLeague
		switch args[0] {
		case "al", "american":
			c = sdk.MLB.AmericanLeague
		case "nl", "national":
			c = sdk.MLB.NationalLeague
		default:
			return fmt.Errorf("Failed to understand MLB league passed.")
		}

		c.Standings()

		// Parse the divison argument
		var d *sdk.Division
		switch args[1] {
		case "west":
			d = c.West
		case "central":
			d = c.Central
		case "east":
			d = c.East
		default:
			return fmt.Errorf("Failed to understand MLB division passed.")
		}

		fmt.Printf(" %s %s      |  W  |  L  |  W%%  |  GB  \n", c.NameShort, d.Name)
		fmt.Printf("==============|=====|=====|======|======\n")
		for _, t := range d.Teams() {
			fmt.Printf(" %-12s | %3d | %3d | %s | %s\n", t.Name, t.Wins, t.Losses, t.WinPercentage(), t.GamesBack())
		}

		return nil
	},
}

func init() {
	mlbCmd.AddCommand(standingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
