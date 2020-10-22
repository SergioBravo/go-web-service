package jokes

import (
	"encoding/json"
	"fmt"
	"go-web-service/internal/api"
	"net/http"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

// NewJokeClient creates a new joke client
func NewJokeClient(baseURL string) *JokeClient {
	return &JokeClient{
		url: baseURL,
	}
}

// GetJoke returns joke
func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s",
			http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, nil
	}
	return &data, nil
}
