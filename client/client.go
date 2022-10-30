package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-make/model"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL    string
	AuthToken  string
	region     model.Region
	HTTPClient HTTPClient
}

func NewClient(authToken string, region model.Region, httpClient HTTPClient) *Client {
	// https://www.make.com/en/api-documentation

	return &Client{
		BaseURL:   fmt.Sprintf("https://%s.make.com/api", toDomain(region)),
		AuthToken: authToken,
		region:    region,
		HTTPClient: httpClient,
	}
}

func (c *Client) SendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.AuthToken))

	fmt.Println(req)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes model.ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return fmt.Errorf("uknown error, status code: %d", res.StatusCode)
		}
		return errors.New(errRes.Message)
	}

	successRes := model.SuccessResponse{
		Data: v,
	}
	if err := json.NewDecoder(res.Body).Decode(&successRes.Data); err != nil {
		return err
	}
	return nil
}

func toDomain(region model.Region) string {
	var domain string
	if region == model.EU {
		domain = "eu1"
	} else if region == model.US {
		domain = "us1"
	} else {
		panic(fmt.Sprintf("Unsupported region: %s", region))
	}
	return domain
}
