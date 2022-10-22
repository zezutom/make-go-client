//go:build integration
// +build integration

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-make/client"
	"go-make/model"
	"os"
	"testing"
)

func TestListConnections(t *testing.T) {
	c := client.NewClient(os.Getenv("MAKE_API_TOKEN"), model.EU)
	apiV2 := c.NewApiV2()
	res, err := apiV2.ListConnections(context.Background(), &client.ListConnectionsRequest{
		TeamId: os.Getenv("MAKE_TEAM_ID"),
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	assert.GreaterOrEqual(t, len(res.Connections), 1, "expecting at least 1 connection")
}
