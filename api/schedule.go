package api

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Response structure from <baseURL>/api/v1/schedule
type scheduleResponse struct {
	TotalGames int `json:"totalGames"`
	Dates      []struct {
		Date       string `json:"date"`
		TotalGames int    `json:"totalGames"`
		Games      []struct {
			GamePK   int    `json:"gamePk"`
			GameType string `json:"gameType"`
			Season   string `json:"season"`
			GameDate string `json:"gameDate"`
			Teams    struct {
				Away struct {
					Team struct {
						ID uint8 `json:"id"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID uint8 `json:"id"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
		} `json:"games"`
	} `json:"dates"`
}

/*
 * Schedule returns a struct with the output from the schedule API endpoint.
 */
func (c *client) Schedule(teamID, daysBefore, daysAfter uint8) (*scheduleResponse, error) {

	format := "2006-01-02"
	startDate := time.Now().AddDate(0, 0, int(daysBefore)*-1)
	endDate := time.Now().AddDate(0, 0, int(daysAfter))

	params := fmt.Sprintf("sportId=1&teamId=%d&startDate=%s&endDate=%s", teamID, startDate.Format(format), endDate.Format(format))

	url, err := c.baseURL.Parse("/api/v1/schedule?" + params)
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

	var schedule scheduleResponse
	err = json.Unmarshal(body, &schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}
