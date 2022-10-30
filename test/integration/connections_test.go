package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-make/client"
	"go-make/model"
	"net/http"
	"os"
	"testing"
	"time"
)

var(
	apiV2 = client.NewClient(os.Getenv("MAKE_API_TOKEN"), model.EU, &http.Client{
		Timeout: time.Minute,
	}).NewApiV2()
)

func Test_IT_ListConnections(t *testing.T) {
	res, err := apiV2.ListConnections(context.Background(), &client.ListConnectionsReq{
		TeamId: os.Getenv("MAKE_TEAM_ID"),
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	if (res != nil) {
		assert.GreaterOrEqual(t, len(res.Connections), 1, "expecting at least 1 connection")
	}
}
