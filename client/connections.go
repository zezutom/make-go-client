package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ListConnectionsReq struct {
	TeamId string   `json:"teamId"`
	Type   []string `json:"type"`
	Cols   []string `json:"cols"`
}
type ListConnectionsRes struct {
	Connections []Connection `json:"connections"`
}
type CreateConnectionReq struct {
	TeamId string `json:"teamId"`
	Body   CreateConnectionReqBody
}
type CreateConnectionReqBody struct {
	AccountName  string   `json:"accountName"`
	AccountType  string   `json:"accountType"`
	Scopes       []string `json:"scopes"`
	ClientId     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
}
type CreateConnectionRes struct {
	Connection Connection `json:"connection"`
}
type Connection struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	AccountName  string      `json:"accountName"`
	AccountLabel interface{} `json:"accountLabel"`
	PackageName  string      `json:"packageName"`
	Expire       interface{} `json:"expire"`
	Metadata     interface{} `json:"metadata"`
	TeamID       int         `json:"teamId"`
	Theme        string      `json:"theme"`
	Upgradeable  bool        `json:"upgradeable"`
	Scopes       int         `json:"scopesCnt"`
	Scoped       bool        `json:"scoped"`
	AccountType  string      `json:"accountType"`
	Editable     bool        `json:"editable"`
	UID          string      `json:"uid"`
}

func (apiV2 *ApiV2) ListConnections(ctx context.Context, request *ListConnectionsReq) (*ListConnectionsRes, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s/connections?teamId=%s",
			apiV2.Client.BaseURL,
			apiV2.Version,
			request.TeamId),
		nil)
	if err != nil {
		return nil, err
	}
	res := ListConnectionsRes{}
	if err := apiV2.Client.SendRequest(ctx, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (apiV2 *ApiV2) CreateConnection(ctx context.Context, request *CreateConnectionReq) (*CreateConnectionRes, error) {
	body, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s/connections?teamId=%s",
			apiV2.Client.BaseURL,
			apiV2.Version,
			request.TeamId),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	res := CreateConnectionRes{}
	if err := apiV2.Client.SendRequest(ctx, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func toQueryParam(name string, values []string) string {
	if len(values) == 0 {
		return ""
	} else {
		builder := strings.Builder{}
		for i := 0; i < len(values); i++ {
			builder.WriteString(fmt.Sprintf("&%s=%s", name, values[i]))
		}
		return builder.String()
	}
}
