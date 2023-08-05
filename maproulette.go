// Package maproulette provides a Go client for the MapRoulette API.
//
// The starting point for most applications will be the NewClient function,
// which creates a new API client instance. The client provides methods
// to interact with MapRoulette API resources, such as challenges.
package maproulette

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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

func (mr *MapRoulette) GetChallenges(ctx context.Context, limit int) ([]Challenge, error) {
	var challenges []Challenge
	url := fmt.Sprintf("%s/challenges?limit=%d", mr.BaseURL, limit)
	err := mr.doRequest(ctx, http.MethodGet, url, nil, &challenges)
	return challenges, err
}

// GetChallenge returns a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
func (mr *MapRoulette) GetChallenge(ctx context.Context, id int) (Challenge, error) {
	var challenge Challenge
	url := fmt.Sprintf("%s/challenge/%d", mr.BaseURL, id)
	err := mr.doRequest(ctx, http.MethodGet, url, nil, &challenge)
	return challenge, err
}

// PostChallenge creates a new challenge on the MapRoulette API.
//
// The challenge parameter specifies the challenge to create.
func (mr *MapRoulette) PostChallenge(ctx context.Context, challenge Challenge) (Challenge, error) {
	var newChallenge Challenge
	url := fmt.Sprintf("%s/challenge", mr.BaseURL)
	err := mr.doRequest(ctx, http.MethodPost, url, challenge, &newChallenge)
	return newChallenge, err
}

// GetChallengeTasks returns a list of tasks for a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
// the optional limit parameter specifies the maximum number of tasks to return.
// The default limit is 10.
func (mr *MapRoulette) GetChallengeTasks(ctx context.Context, id int, limit ...int) ([]Task, error) {
	var tasks []Task
	url := fmt.Sprintf("%s/challenge/%d/tasks", mr.BaseURL, id)
	if len(limit) > 0 {
		url = fmt.Sprintf("%s?limit=%d", url, limit[0])
	}
	err := mr.doRequest(ctx, http.MethodGet, url, nil, &tasks)
	return tasks, err
}

// GetRandomChallengeTasks returns a random task for a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
// The optional search parameter specifies a search string to filter tasks by.
// The optional tags parameter specifies a comma-separated list of tags to filter tasks by.
// The optional limit parameter specifies the maximum number of tasks to return.
// The optional proximity parameter identifies the Task id to use to find a nearby task.
func (mr *MapRoulette) GetRandomChallengeTasks(ctx context.Context, options *GetRandomChallengeTasksOptions) ([]Task, error) {
	var tasks []Task
	url := fmt.Sprintf("%s/challenge/%d/tasks/random", mr.BaseURL, options.ChallengeID)
	if options.Search != "" {
		url = fmt.Sprintf("%s?search=%s", url, options.Search)
	}
	// if length of tags is greater than 0, add comma separated tags to url
	if len(options.Tags) > 0 {
		url = fmt.Sprintf("%s?tags=%s", url, strings.Join(options.Tags, ","))
	}
	if options.Limit > 0 {
		url = fmt.Sprintf("%s?limit=%d", url, options.Limit)
	}
	if options.Proximity > 0 {
		url = fmt.Sprintf("%s?proximity=%d", url, options.Proximity)
	}
	err := mr.doRequest(ctx, http.MethodGet, url, nil, &tasks)
	return tasks, err

}

// AddTasksToChallenge creates new tasks for a challenge on the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to add tasks to.
// The tasks parameter specifies the tasks to add.
// TODO this requires tasks to be wrapped as GeoJSON features
func (mr *MapRoulette) AddTasksToChallenge(ctx context.Context, id int, tasks []Task) ([]Task, error) {
	var newTasks []Task
	url := fmt.Sprintf("%s/challenge/%d/tasks", mr.BaseURL, id)
	err := mr.doRequest(ctx, http.MethodPost, url, tasks, &newTasks)
	if err != nil {
		return nil, err
	}
	return newTasks, nil
}

// AddTasks adds tasks to a Challenge from the GeoJSON payload.
// We make no attempt to validate the GeoJSON payload.
//
// The id parameter specifies the ID of the challenge to add tasks to.
// The payload parameter specifies the GeoJSON payload containing the tasks to add.
// Return value is the HTTP response code.
//
// This function requires an API key with write access to the challenge.
func (mr *MapRoulette) AddTasks(ctx context.Context, id int, payload []byte) error {
	url := fmt.Sprintf("%s/challenge/%d/tasks", mr.BaseURL, id)
	return mr.doRequest(ctx, http.MethodPost, url, payload, nil)
}
