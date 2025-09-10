package cmd

import (
	"fmt"

	"github.com/gopherlibs/big-league-stats/sdk"
	"github.com/spf13/cobra"
)

var (
	team   string
	teamID uint8
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Display the schedule for the provided team(s)",
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if team == "" && teamID == 0 {
			return fmt.Errorf("--team or --teamID must be set")
		}

		// team and teamID are mutually exclusive flags
		if team != "" && teamID != 0 {
			return fmt.Errorf("--team and --teamID are mutually exclusive")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		//slog.SetLogLoggerLevel(slog.LevelDebug)

		sched := sdk.GetSchedule(teamID, 1, 4)

		fmt.Printf("Recent & Upcoming Games\n")
		for _, g := range sched.Games {
			fmt.Printf("==============\n")
			fmt.Printf("%s\n", g.Date.Format("01/02"))
			fmt.Printf("%s\n", g.Away.Name)
			fmt.Printf("%s\n", g.Home.Name)
		}

		return nil
	},
}

func init() {

	scheduleCmd.Flags().Uint8Var(&teamID, "teamID", 0, "Team ID")
	scheduleCmd.Flags().StringVar(&team, "team", "", "Team name")

	mlbCmd.AddCommand(scheduleCmd)
}
