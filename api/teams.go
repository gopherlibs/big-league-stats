package api

import (
	"encoding/json"
	"io"
	"strconv"
)

// Response structure from <baseURL>/teams
type teamsResponse struct {
	Teams []struct {
		ID           uint8  `json:"id"`
		Abbreviation string `json:"abbreviation"`
		LocationName string `json:"locationName"`
		TeamName     string `json:"teamName"`
		League       struct {
			ID uint8 `json:"id"`
		} `json:"league"`
		Division struct {
			ID uint8 `json:"id"`
		} `json:"division"`
	} `json:"teams"`
}

func (c *client) Teams(leagueID uint8) (*teamsResponse, error) {

	url, err := c.baseURL.Parse("/api/v1/teams?leagueId=" + strconv.FormatUint(uint64(leagueID), 10))
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

	var teams teamsResponse
	err = json.Unmarshal(body, &teams)
	if err != nil {
		return nil, err
	}

	return &teams, nil
}
