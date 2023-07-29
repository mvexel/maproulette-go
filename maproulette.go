// Package maproulette provides a Go client for the MapRoulette API.
//
// The starting point for most applications will be the NewClient function,
// which creates a new API client instance. The client provides methods
// to interact with MapRoulette API resources, such as challenges.
package maproulette

import (
	"fmt"
	"net/http"
)

// Version is the version of this library.
var Version string

// baseURL is the base URL for the MapRoulette API
const prodBaseURL = "https://maproulette.org/api/v2"

// NewMapRouletteClient creates a new MapRoulette API client.
//
// The client communicates with the MapRoulette API at the specified baseURL,
// and uses the specified API key for authentication.
func NewMapRouletteClient(options *MapRouletteClientOptions) *MapRoulette {
	// Set default values if not provided.
	if options.BaseURL == "" {
		options.BaseURL = prodBaseURL
	}
	if options.Client == nil {
		options.Client = http.DefaultClient
	}

	return &MapRoulette{
		APIKey:  options.APIKey,
		BaseURL: options.BaseURL,
		Client:  options.Client,
	}
}

// GetChallenges returns a list of challenges from the MapRoulette API.
//
// The limit parameter specifies the maximum number of challenges to return.

func (mr *MapRoulette) GetChallenges(limit int) ([]Challenge, error) {
	var challenges []Challenge
	url := fmt.Sprintf("%s/challenges?limit=%d", mr.BaseURL, limit)
	err := mr.getJSON(url, &challenges)
	return challenges, err
}

// GetChallenge returns a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
func (mr *MapRoulette) GetChallenge(id int) (Challenge, error) {
	var challenge Challenge
	url := fmt.Sprintf("%s/challenge/%d", mr.BaseURL, id)
	err := mr.getJSON(url, &challenge)
	return challenge, err
}

// PostChallenge creates a new challenge on the MapRoulette API.
//
// The challenge parameter specifies the challenge to create.
func (mr *MapRoulette) PostChallenge(challenge Challenge) (Challenge, error) {
	var newChallenge Challenge
	url := fmt.Sprintf("%s/challenge", mr.BaseURL)
	err := mr.postJSON(url, challenge, &newChallenge)
	return newChallenge, err
}

// GetChallengeTasks returns a list of tasks for a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
// the optional limit parameter specifies the maximum number of tasks to return.
// The default limit is 10.
func (mr *MapRoulette) GetChallengeTasks(id int, limit ...int) ([]Task, error) {
	var tasks []Task
	url := fmt.Sprintf("%s/challenge/%d/tasks", mr.BaseURL, id)
	if len(limit) > 0 {
		url = fmt.Sprintf("%s?limit=%d", url, limit[0])
	}
	err := mr.getJSON(url, &tasks)
	return tasks, err
}
