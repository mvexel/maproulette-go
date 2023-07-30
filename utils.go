package maproulette

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// doRequest performs the specified request, adding the API key and referer
// headers.
// doRequest performs a request to the specified URL, using the specified
// API key for authentication, and unmarshals the response into the target
// interface. The HTTP method and payload (for POST and PUT requests) are specified as parameters.
func (mr *MapRoulette) doRequest(ctx context.Context, method, url string, payload, target interface{}) error {
	var req *http.Request
	var err error

	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	case http.MethodPost, http.MethodPut:
		// Convert payload to JSON
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
	default:
		return fmt.Errorf("Invalid method: %s", method)
	}

	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("API-Key", mr.APIKey)
	req.Header.Set("Referer", "https://github.com/mvexel/maproulette-go v"+Version+" ")

	resp, err := mr.Client.Do(req)
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
