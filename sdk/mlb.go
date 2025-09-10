package sdk

import (
	"log"
	"log/slog"
	"strconv"
	"time"

	"github.com/gopherlibs/big-league-stats/api"
)

type Game struct {
	Date time.Time
	Away *team
	Home *team
}

type Schedule struct {
	Games []Game
}

type MLBLeague struct {
	*basicConference
	West    *Division
	Central *Division
	East    *Division
}

func NewMLBLeague(id uint8, name, nameShort, slug string, west *Division, central *Division, east *Division) *MLBLeague {

	l := &MLBLeague{
		&basicConference{
			id:        id,
			Name:      name,
			NameShort: nameShort,
			Slug:      slug,
		},
		west,
		central,
		east,
	}
	l.add(west)
	l.add(central)
	l.add(east)

	return l
}

func GetSchedule(teamID, daysBefore, daysAfter uint8) Schedule {

	c := api.New()

	rawSched, err := c.Schedule(teamID, daysBefore, daysAfter)
	if err != nil {
		slog.Error("GetSchedule failed.", "err", err)
		log.Fatal("Request failed")
	}

	var sched Schedule

	// for each day
	for _, d := range rawSched.Dates {

		// for each game on the day
		for _, g := range d.Games {

			gameDate, err := time.Parse("2006-01-02", d.Date)
			if err != nil {
				slog.Error("Failed to parse date in schedule.", "date", d.Date)
			}

			sched.Games = append(sched.Games, Game{
				Date: gameDate,
				Away: TeamByID(g.Teams.Away.Team.ID),
				Home: TeamByID(g.Teams.Home.Team.ID),
			})
		}
	}

	return sched
}

func (ml *MLBLeague) Standings() {

	c := api.New()

	standings, err := c.Standings(ml.ID())
	if err != nil {
		log.Fatal("Request failed")
	}

	// for each divison
	for _, r := range standings.Records {

		d := ml.DivisionByID(r.Division.ID)

		// for each team
		for _, tr := range r.TeamRecords {

			t := TeamByID(tr.Team.ID)
			if t == nil {
				log.Fatal("Team is nil")
			}

			// Process divisionRank
			dr, err := strconv.ParseUint(tr.DivisionRank, 10, 8)
			if err != nil {
				log.Fatal("Division rank is not a number")
			}
			t.divisionRank = uint8(dr)

			// Process gamesBack
			if tr.GamesBack == "-" {
				t.gamesBack = 0
			} else {
				t.gamesBack, err = strconv.ParseFloat(tr.GamesBack, 64)
				if err != nil {
					log.Fatal("Division rank is not a number")
				}
			}

			// Process stats
			t.Wins = tr.LeagueRecord.Wins
			t.Losses = tr.LeagueRecord.Losses
			t.Percentage, err = strconv.ParseFloat(tr.LeagueRecord.Percentage, 64)
			if err != nil {
				log.Fatal("Win percentage is not a valid float.")
			}
		}

		d.Rank()
	}
}

type mlbLeagues struct {
	*sport
	AmericanLeague *MLBLeague
	NationalLeague *MLBLeague
}

func NewMLBLeagues(al *MLBLeague, nl *MLBLeague) *mlbLeagues {

	s := &mlbLeagues{&sport{}, al, nl}
	s.add(al)
	s.add(nl)

	return s
}

var MLB *mlbLeagues

func GetTeams(id uint8) {

	c := api.New()

	rawTeams, err := c.Teams(id)
	if err != nil {
		log.Fatal("Request failed")
	}

	// for each team
	for _, t := range rawTeams.Teams {

		teams = append(teams, &team{
			ID:         t.ID,
			Location:   t.LocationName,
			Name:       t.TeamName,
			Code:       t.Abbreviation,
			LeagueID:   t.League.ID,
			DivisionID: t.Division.ID,
		})
	}
}
