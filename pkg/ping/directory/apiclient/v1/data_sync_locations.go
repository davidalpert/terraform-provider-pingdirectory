package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type LocationListResponse struct {
	Schemas      []string                   `json:"schemas"`
	TotalResults int64                      `json:"totalResults"`
	Resources    []LocationResourceResponse `json:"Resources"`
}

type LocationResourceResponse struct {
	Schemas                   []string `json:"schemas"`
	ID                        string   `json:"id"`
	Description               string   `json:"description"`
	PreferredFailoverLocation []string `json:"preferredFailoverLocation"`
}

type Location struct {
	Name                      string   `json:"locationName"`
	Description               string   `json:"description"`
	PreferredFailoverLocation []string `json:"preferredFailoverLocation"`
}

func (c *DataSyncService) LocationsGetAll(ctx context.Context) ([]Location, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/config/locations", c.BaseUrl), nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("expected 200 OK; received %s", resp.Status)
	}

	body := LocationListResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, nil, err
	}

	result := make([]Location, len(body.Resources))
	for i, r := range body.Resources {
		result[i] = Location{
			Name:                      r.ID,
			Description:               r.Description,
			PreferredFailoverLocation: r.PreferredFailoverLocation,
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, resp, nil
}

func (c *DataSyncService) LocationsGet(ctx context.Context, name string) (*Location, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/config/locations/%s", c.BaseUrl, name), nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("expected 200 OK; received %s", resp.Status)
	}

	body := LocationResourceResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, nil, err
	}

	result := &Location{
		Name:                      body.ID,
		Description:               body.Description,
		PreferredFailoverLocation: body.PreferredFailoverLocation,
	}

	return result, resp, nil
}

func (c *DataSyncService) LocationsCreate(ctx context.Context, l Location) (*http.Response, error) {
	payload, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/config/locations", c.BaseUrl), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("expected 200 OK; received %s", resp.Status)
	}

	return resp, nil
}

func (c *DataSyncService) LocationUpdate(ctx context.Context, l Location) (*http.Response, error) {
	payload, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/config/locations/%s", c.BaseUrl, l.Name), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected 200 OK; received %s", resp.Status)
	}

	return resp, nil
}

func (c *DataSyncService) LocationDeleteByName(ctx context.Context, name string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/config/locations/%s", c.BaseUrl, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("expected 204 No Content; received %s", resp.Status)
	}

	return resp, nil
}
