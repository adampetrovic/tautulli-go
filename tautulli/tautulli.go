package tautulli

import (
	"github.com/dghubble/sling"
)

const tautulliApiPath = "/api/v2"

type Client struct {
	sling *sling.Sling

	// Tautulli API Services

}

type DefaultResponse struct {
	Response struct {
		Data    interface{}
		Message string
		Result  string
	}
}

type DefaultParams struct {
	ApiKey ApiKey `json:"api_key"`
}

type ApiKey string

func NewClient(baseUrl string) *Client {
	apiUrl := baseUrl + tautulliApiPath
	return &Client{
		sling: sling.New().Base(apiUrl),
	}
}

func (s *Client) WithCredentials(username string, password string) *Client {
	credentialQuery := struct {
		Cmd      string
		Username string
		Password string
	}{
		Cmd:      "get_apikey",
		Username: username,
		Password: password,
	}

	defaultResponse := new(DefaultResponse)

	s.sling.New().Get("").QueryStruct(credentialQuery).Receive(defaultResponse, nil)
	return s.WithApiKey(defaultResponse.Response.Data.(ApiKey))
}

func (s *Client) WithApiKey(apiKey ApiKey) *Client {
	defaultParams := DefaultParams{
		ApiKey: apiKey,
	}
	return &Client{
		sling: s.sling.New().QueryStruct(defaultParams),
	}
}
