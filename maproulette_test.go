package maproulette

import (
	"testing"
	// maproulette "github.com/mvexel/maproulette-go"
)

const testBaseURL = "https://staging.maproulette.org/api/v2"
const testApiKey = "10406|0d174387-370f-4d91-901c-8bb1c77edbf0"

// use staging server to test getting challenge number 1
// TODO: mock the API so we won't need to use the staging server
func TestGetChallenge(t *testing.T) {
	mr := NewMapRouletteClient(&MapRouletteClientOptions{
		APIKey:  testApiKey,
		BaseURL: testBaseURL,
	})
	challenge, err := mr.GetChallenge(1)
	if err != nil {
		t.Errorf("Error getting challenge: %v", err)
	}
	if challenge.ID != 1 {
		t.Errorf("Expected challenge ID 1, got %d", challenge.ID)
	}
	if challenge.Name != "string" {
		t.Errorf("Expected challenge name \"string\", got %s", challenge.Name)
	}
	if challenge.CompletionPercentage != 40 {
		t.Errorf("Expected completion percentage 40, got %d", challenge.CompletionPercentage)
	}
	if challenge.TasksRemaining != 7067 {
		t.Errorf("Expected tasks remaining 7067, got %d", challenge.TasksRemaining)
	}
}

func TestGetTasks(t *testing.T) {
	mr := NewMapRouletteClient(&MapRouletteClientOptions{
		APIKey:  testApiKey,
		BaseURL: testBaseURL,
	})
	tasks, err := mr.GetChallengeTasks(1)
	if err != nil {
		t.Errorf("Error getting task: %v", err)
	}
	// if tasks count is not 10, we have a problem
	if len(tasks) != 10 {
		t.Errorf("Expected 10 tasks, got %d", len(tasks))
	}
}
