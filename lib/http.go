package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient is a wrapper for making HTTP requests to the fastotp API.
type APIClient struct {
	BaseURL string
	APIKey  string
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient(baseURL, apiKey string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

// Post sends a POST request to the specified endpoint with the given payload.
func (c *APIClient) Post(endpoint string, payload interface{}) (*http.Response, error) {
	url := c.BaseURL + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)

	client := http.DefaultClient
	return client.Do(req)
}

// Get sends a GET request to the specified endpoint, appending id as a path parameter
func (c *APIClient) Get(id string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, id)
	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)

	client := http.DefaultClient
	return client.Do(req)
}
