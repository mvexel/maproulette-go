// Package maproulette provides a Go client for the MapRoulette API.
//
// The starting point for most applications will be the NewClient function,
// which creates a new API client instance. The client provides methods
// to interact with MapRoulette API resources, such as challenges.
package maproulette

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Version string

const baseURL = "https://maproulette.org/api/v2"

// MapRoulette represents a client for the MapRoulette API.
//
// To create a new client, use the NewMapRouletteClient function.
type MapRoulette struct {
	APIKey string
}

// TODO - use real GeoJSON types from twpayne/go-geojson?
type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"` // []float64 for Point, [][]float64 for Polygon
}

// Challenge represents a challenge in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// challenge data.
type Challenge struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Created              string `json:"created"`
	Modified             string `json:"modified"`
	Description          string `json:"description"`
	Deleted              bool   `json:"deleted"`
	InfoLink             string `json:"infoLink"`
	General              General
	Creation             Creation
	Priority             Priority
	Extra                Extra
	Status               int     `json:"status"`
	StatusMessage        string  `json:"statusMessage"`
	LastTaskRefresh      string  `json:"lastTaskRefresh"`
	DataOriginDate       string  `json:"dataOriginDate"`
	Location             GeoJSON `json:"location"`
	Bounding             GeoJSON `json:"bounding"`
	CompletionPercentage int     `json:"completionPercentage"`
	TasksRemaining       int     `json:"tasksRemaining"`
}

// General represents the general section of a challenge with basic information
// about its owner, parent project, instructions, etc.
type General struct {
	Owner           int     `json:"owner"`
	Parent          Project `json:"parent"`
	Instruction     string  `json:"instruction"`
	Difficulty      int     `json:"difficulty"`
	Blurb           string  `json:"blurb"`
	Enabled         bool    `json:"enabled"`
	Featured        bool    `json:"featured"`
	CooperativeType int     `json:"cooperativeType"`
	Popularity      int     `json:"popularity"`
	CheckinComment  string  `json:"checkinComment"`
	CheckinSource   string  `json:"checkinSource"`
	VirtualParents  []int   `json:"virtualParents"`
	RequiresLocal   bool    `json:"requiresLocal"`
}

// Creation represents the creation section of a challenge with information
// about the overpass query or remote GeoJSON link.
type Creation struct {
	OverpassQL         string `json:"overpassQL"`
	RemoteGeoJson      string `json:"remoteGeoJson"`
	OverpassTargetType string `json:"overpassTargetType"`
}

// Priority represents the priority section of a challenge with information
// about the default priority and priority rules.
type Priority struct {
	DefaultPriority    int    `json:"defaultPriority"`
	HighPriorityRule   string `json:"highPriorityRule"`
	MediumPriorityRule string `json:"mediumPriorityRule"`
	LowPriorityRule    string `json:"lowPriorityRule"`
}

// Extra represents the extra section of a challenge with information
// which are optional to set when creating a challenge.
type Extra struct {
	DefaultZoom          int      `json:"defaultZoom"`
	MinZoom              int      `json:"minZoom"`
	MaxZoom              int      `json:"maxZoom"`
	DefaultBasemap       int      `json:"defaultBasemap"`
	DefaultBasemapId     string   `json:"defaultBasemapId"`
	CustomBasemap        string   `json:"customBasemap"`
	UpdateTasks          bool     `json:"updateTasks"`
	ExportableProperties string   `json:"exportableProperties"`
	OsmIdProperty        string   `json:"osmIdProperty"`
	PreferredTags        string   `json:"preferredTags"`
	PreferredReviewTags  string   `json:"preferredReviewTags"`
	LimitTags            bool     `json:"limitTags"`
	LimitReviewTags      bool     `json:"limitReviewTags"`
	TaskStyles           string   `json:"taskStyles"`
	TaskBundleIdProperty string   `json:"taskBundleIdProperty"`
	IsArchived           bool     `json:"isArchived"`
	ReviewSetting        int      `json:"reviewSetting"`
	SystemArchivedAt     int      `json:"systemArchivedAt"`
	Presets              []string `json:"presets"`
}

// Project represents a project in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// project data.
type Project struct {
	ID          int     `json:"id"`
	Owner       int     `json:"owner"`
	Name        string  `json:"name"`
	Created     string  `json:"created"`
	Modified    string  `json:"modified"`
	Description string  `json:"description"`
	Grants      []Grant `json:"grants"`
	Enabled     bool    `json:"enabled"`
	DisplayName string  `json:"displayName"`
	Deleted     bool    `json:"deleted"`
	Featured    bool    `json:"featured"`
}

// Grant represents a grant in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// grant data.
type Grant struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Grantee Grantee `json:"grantee"`
	Role    int     `json:"role"`
	Target  Target  `json:"target"`
}

// Grantee represents a grantee in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// grantee data.
type Grantee struct {
	GranteeType GranteeType `json:"granteeType"`
	GranteeID   int         `json:"granteeId"`
}

// GranteeType represents a grantee type in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// grantee type data.
type GranteeType struct {
	ID int `json:"id"`
}

// Target represents a target in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// target data.
type Target struct {
	ObjectType ObjectType `json:"objectType"`
	ObjectID   int        `json:"objectId"`
}

// ObjectType represents an object type in MapRoulette.
//
// This struct is used by various methods in the Client to receive and send
// object type data.
type ObjectType struct {
	ID int `json:"id"`
}

// getJSON performs a GET request to the specified URL, using the specified
// API key for authentication, and unmarshals the response into the target
// interface.
func getJSON(url string, apiKey string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("API-Key", apiKey)
	req.Header.Set("Referer", "https://github.com/mvexel/maproulette-go v"+Version+" ")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

// NewMapRouletteClient creates a new MapRoulette API client.
//
// The client communicates with the MapRoulette API at the specified baseURL,
// and uses the specified API key for authentication.
func NewMapRouletteClient(apiKey string) *MapRoulette {
	return &MapRoulette{
		APIKey: apiKey,
	}
}

// GetChallenges returns a list of challenges from the MapRoulette API.
//
// The limit parameter specifies the maximum number of challenges to return.

func (mr *MapRoulette) GetChallenges(limit int) ([]Challenge, error) {
	var challenges []Challenge
	url := fmt.Sprintf("%s/challenges?limit=%d", baseURL, limit)
	err := getJSON(url, mr.APIKey, &challenges)
	return challenges, err
}

// GetChallenge returns a challenge from the MapRoulette API.
//
// The id parameter specifies the ID of the challenge to return.
func (mr *MapRoulette) GetChallenge(id int) (Challenge, error) {
	var challenge Challenge
	url := fmt.Sprintf("%s/challenge/%d", baseURL, id)
	err := getJSON(url, mr.APIKey, &challenge)
	return challenge, err
}
