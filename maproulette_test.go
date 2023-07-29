package maproulette

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testBaseURL = "https://staging.maproulette.org/api/v2"
const testApiKey = "10406|0d174387-370f-4d91-901c-8bb1c77edbf0"

func TestGetChallenge(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() != "/api/v2/challenge/1" {
			t.Errorf("want: '/api/v2/challenge/1', got: '%s'", req.URL.String())
		}
		// Send response to be tested
		rw.Write([]byte(`{"id": 1, "name": "challenge1"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Server's URL as base URL
	mr := NewMapRouletteClient(testApiKey)
	mr.BaseURL = testBaseURL
	challenge, err := mr.GetChallenge(1)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if challenge.ID != 1 {
		t.Errorf("Expected challenge ID 1, got %d", challenge.ID)
	}
	if challenge.Name != "Old Aerodromes USA" {
		t.Errorf("Expected challenge name 'Default Dummy Survey', got '%s'", challenge.Name)
	}
}

// a test function that posts a new challenge to the API
func TestPostChallenge(t *testing.T) {
	mr := NewMapRouletteClient(testApiKey)
	mr.BaseURL = testBaseURL

	// first, get the current highest challenge ID
	challenges, err := mr.GetChallenges(1)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	highestID := challenges[0].ID
	challenge := Challenge{
		Name: "Test Challenge",
		General: General{
			Owner:       1,
			Parent:      Project{ID: 1},
			Instruction: "Test instructions",
			Difficulty:  1,
			Blurb:       "Test blurb",
			Enabled:     true,
			Featured:    false,
		},
		Creation: Creation{
			OverpassQL: "Test Overpass QL",
		},
		Extra: Extra{},
	}
	newChallenge, err := mr.PostChallenge(challenge)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if newChallenge.ID != highestID+1 {
		t.Errorf("Expected challenge ID %d, got %d", highestID+1, newChallenge.ID)
	}
	if newChallenge.Name != "Test Challenge" {
		t.Errorf("Expected challenge name 'Test Challenge', got '%s'", newChallenge.Name)
	}
}

func TestGetChallenges(t *testing.T) {
	// Use Server's URL as base URL
	mr := NewMapRouletteClient(testApiKey)
	challenges, err := mr.GetChallenges(5)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if len(challenges) != 5 {
		t.Errorf("Expected 5 challenges, got %d", len(challenges))
	}
	for i, challenge := range challenges {
		if challenge.ID != i+1 {
			t.Errorf("Expected challenge ID %d, got %d", i+1, challenge.ID)
		}
		if challenge.Name != "challenge"+fmt.Sprint(i+1) {
			t.Errorf("Expected challenge name 'challenge%d', got '%s'", i+1, challenge.Name)
		}
	}
}
