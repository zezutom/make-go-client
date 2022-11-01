package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-make/client"
	"go-make/model"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	apiV2 = client.NewClient(os.Getenv("MAKE_API_TOKEN"), model.EU, &http.Client{
		Timeout: time.Minute,
	}).NewApiV2()
)

func Test_ListConnections(t *testing.T) {
	res, err := apiV2.ListConnections(context.Background(), &client.ListConnectionsReq{
		TeamId: os.Getenv("MAKE_TEAM_ID"),
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	if res != nil {
		assert.GreaterOrEqual(t, len(res.Connections), 1, "expecting at least 1 connection")
	}
}

func Test_CreateConnection(t *testing.T) {
	res, err := apiV2.CreateConnection(context.Background(), &client.CreateConnectionReq{
		TeamId: os.Getenv("MAKE_TEAM_ID"),
		Body: client.CreateConnectionReqBody{
			AccountName:  "Slack Test",
			AccountType:  "slack",
			Scopes:       []string{"chat:write"},
			ClientId:     "123456",
			ClientSecret: "secret",
		},
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	if res != nil {
		fmt.Println(res.Connection)
		assert.NotNil(t, res.Connection, "expecting a fully initialized connection")
	}
}
