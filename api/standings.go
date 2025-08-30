package api

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Response structure from <baseURL>/api/v1/standings
type standingsResponse struct {
	Records []struct {
		Division struct {
			ID uint8 `json:"id"`
		} `json:"division"`
		TeamRecords []struct {
			Team struct {
				ID   uint8  `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			LeagueRecord struct {
				Wins       uint8  `json:"wins"`
				Losses     uint8  `json:"losses"`
				Percentage string `json:"pct"`
			} `json:"leagueRecord"`
			DivisionRank string `json:"divisionRank"`
			// GameBack is either a float as a string, or a dash for the 1st place team
			GamesBack string `json:"gamesBack"`
		} `json:"teamRecords"`
	} `json:"records"`
}

func (c *client) Standings(leagueID uint8) (*standingsResponse, error) {

	params := fmt.Sprintf("leagueId=%d&season=%d&standingsTypes=regularSeason", leagueID, time.Now().Year())

	url, err := c.baseURL.Parse("/api/v1/standings?" + params)
	if err != nil {
		return nil, err
	}

	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var standings standingsResponse
	err = json.Unmarshal(body, &standings)
	if err != nil {
		return nil, err
	}

	return &standings, nil
}
