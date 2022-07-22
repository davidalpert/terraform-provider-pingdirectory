package apiclient

import (
	"fmt"
	"net/http"
)

type service struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func (s *service) formatPath(path string) string {
	return fmt.Sprintf("%s/%s", s.BaseUrl, path)
}

type Client struct {
	Config *ClientConfig

	common   service
	DataSync *DataSyncService
}

type ClientConfig struct {
	BaseURL  string
	Username string
	password string
}

func NewConfig(baseUrl string) *ClientConfig {
	return &ClientConfig{
		BaseURL:  baseUrl,
		Username: "",
		password: "",
	}
}

func (c *ClientConfig) WithBasicAuth(username, password string) *ClientConfig {
	c.Username = username
	c.password = password
	return c
}

func (c *ClientConfig) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("client must have a non-empty BaseUrl (did you set PING_BASE_URL?)")
	}

	if c.Username == "" {
		return fmt.Errorf("client must have a non-empty BaseUrl (did you set PING_USERNAME?)")
	}

	if c.password == "" {
		return fmt.Errorf("client must have a non-empty BaseUrl (did you set PING_PASSWORD?)")
	}

	return nil
}

func (c *ClientConfig) BuildPingApiClient() (*Client, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	commonService := service{
		BaseUrl: c.BaseURL,
		HTTPClient: &http.Client{
			Transport: BasicAuthTransport{
				Username: c.Username,
				Password: c.password,
			},
			Timeout: 0,
		},
	}
	client := &Client{
		Config:   c,
		common:   commonService,
		DataSync: (*DataSyncService)(&commonService),
	}

	return client, nil
}
