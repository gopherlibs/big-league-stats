package sdk

// Organize the league structure and preload JSON data.
func init() {

	MLB = NewMLBLeagues(
		// Build MLB American League
		NewMLBLeague(
			103,
			"American League",
			"AL",
			"al",
			&Division{
				ID:   200,
				Name: "West",
			},
			&Division{
				ID:   202,
				Name: "Central",
			},
			&Division{
				ID:   201,
				Name: "East",
			},
		),

		// Build MLB National League
		NewMLBLeague(
			104,
			"National League",
			"NL",
			"nl",
			&Division{
				ID:   203,
				Name: "West",
			},
			&Division{
				ID:   205,
				Name: "Central",
			},
			&Division{
				ID:   204,
				Name: "East",
			},
		),
	)

	// load MLB teams
	GetTeams(103)
	GetTeams(104)

	// populate MLB teams into divisions
	for _, t := range teams {

		l := MLB.LeagueByID(t.LeagueID)
		d := l.DivisionByID(t.DivisionID)

		d.teams = append(d.teams, t)
	}
}
