package api

import (
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
}

// This takes a full URL complete with endpoint and params.
func (c *client) get(url *url.URL) (*http.Response, error) {

	slog.Debug("Client making a GET request.", "url", url.String())

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)

	return c.httpClient.Do(req)
}

/* New returns a new MLB API client.
 */
func New() *client {

	baseURL, err := url.Parse("https://statsapi.mlb.com")
	if err != nil {
		return nil
	}

	return &client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   baseURL,
		userAgent: "big-league-stats",
	}
}
