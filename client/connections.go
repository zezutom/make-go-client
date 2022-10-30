package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type ListConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}
type ListConnectionsRes struct {
	Connections []Connection `json:"connections"`
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

type ListConnectionsRequest struct {
	TeamId string   `json:"teamId"`
	Type   []string `json:"type"`
	Cols   []string `json:"cols"`
}

type ConnectionList struct {
	Connections []Connection
}

func (apiV2 *ApiV2) ListConnections(ctx context.Context, request *ListConnectionsRequest) (*ListConnectionsRes, error) {
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
