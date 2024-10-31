package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://pterodactyl.file.properties/api/client/servers"

// APIClient represents the API client
type APIClient struct {
	APIKey   string
	ServerID string
}

// NewAPIClient creates a new API client
func NewAPIClient(apiKey, serverID string) *APIClient {
	return &APIClient{
		APIKey:   apiKey,
		ServerID: serverID,
	}
}

// sendRequest sends a request to the Pterodactyl API
func (c *APIClient) sendRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", baseURL, c.ServerID, endpoint)
	var jsonBody []byte
	if body != nil {
		var err error
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// StartServer starts the server
func (c *APIClient) StartServer() (*http.Response, error) {
	return c.sendRequest("POST", "power", map[string]string{"action": "start"})
}

// StopServer stops the server
func (c *APIClient) StopServer() (*http.Response, error) {
	return c.sendRequest("POST", "power", map[string]string{"action": "stop"})
}

// RestartServer restarts the server
func (c *APIClient) RestartServer() (*http.Response, error) {
	return c.sendRequest("POST", "power", map[string]string{"action": "restart"})
}
