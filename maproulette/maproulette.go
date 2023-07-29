package maproulette

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://maproulette.org/api/v2"

type MapRoulette struct {
	APIKey string
}

type Challenge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Instruction string `json:"instruction"`
	Active      bool   `json:"active"`
}

func getJSON(url string, apiKey string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("API-Key", apiKey)

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

func NewMapRouletteClient(apiKey string) *MapRoulette {
	return &MapRoulette{
		APIKey: apiKey,
	}
}

func (mr *MapRoulette) GetChallenges() ([]Challenge, error) {
	var challenges []Challenge
	err := getJSON(baseURL+"/challenge?limit=100", mr.APIKey, &challenges)
	return challenges, err
}
