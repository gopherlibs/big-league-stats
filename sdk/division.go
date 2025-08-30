package sdk

import (
	"log"
	"sort"
)

type Division struct {
	// This is an ID as defined by the MLB Stats API.
	ID    uint8
	Name  string
	Slug  string
	teams []*team
}

func (d *Division) Rank() {

	if d.teams == nil {
		log.Fatal("Can't rank because teams is nil")
	}

	sort.Slice(d.teams, func(i, j int) bool {
		return d.teams[i].divisionRank < d.teams[j].divisionRank
	})
}

func (d *Division) Teams() []*team {
	return d.teams
}
