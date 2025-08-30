package sdk

import "fmt"

type team struct {
	// This is an ID as defined by the MLB Stats API.
	ID       uint8  `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	LeagueID uint8  `json:"league"`

	// Stats
	Wins       uint8
	Losses     uint8
	Percentage float64

	//=== Division related fields
	DivisionID uint8 `json:"division"`
	// a value of zero means undefined
	divisionRank uint8
	gamesBack    float64
}

func (t *team) GamesBack() string {

	if t.gamesBack == 0 {
		return fmt.Sprintf("%4s", "-")
	}

	return fmt.Sprintf("%4.1f", t.gamesBack)
}

func (t *team) WinPercentage() string {

	// Remove a decimal place if winning % is 100.
	if t.Percentage == 1 {
		return "1.00"
	}

	return fmt.Sprintf("%3.3f", t.Percentage)[1:]
}

var teams []*team

func Teams() []*team {
	return teams
}

func TeamByID(id uint8) *team {

	for _, t := range teams {
		if t.ID == id {
			return t
		}
	}

	return nil
}
