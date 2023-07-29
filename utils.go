package maproulette

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// doRequest performs the specified request, adding the API key and referer
// headers.
func (mr *MapRoulette) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("API-Key", mr.APIKey)
	req.Header.Set("Referer", "https://github.com/mvexel/maproulette-go v"+Version+" ")

	return mr.Client.Do(req)
}

// getJSON performs a GET request to the specified URL, using the specified
// API key for authentication, and unmarshals the response into the target
// interface.
func (mr *MapRoulette) getJSON(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := mr.doRequest(req)
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

// postJSON performs a POST request to the specified URL, using the specified
// API key for authentication, and unmarshals the response into the target
// interface.
func (mr *MapRoulette) postJSON(url string, payload interface{}, target interface{}) error {
	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("API-Key", mr.APIKey)
	req.Header.Set("Content-Type", "application/json")
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
