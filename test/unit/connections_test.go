package unit

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"go-make/client"
	"go-make/mock"
	"go-make/model"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var (
	apiV2 = client.NewClient("TEST_TOKEN", model.EU, &mock.MockClient{}).NewApiV2()
)

func Test_ListConnections(t *testing.T) {
	json := `{"connections": [{
				"id": 1,
				"name": "Test Connection",
				"accountName": "google",
				"accountLabel": "Google",
				"metadata": {"type": "email", "value": "j.doe@gmail.com"},
				"teamId": 1,
				"theme": "#fecd5f",
				"upgradeable": false,
				"scopesCnt": 5,
				"scoped": true,
				"accountType": "oauth",
				"editable": false,
				"uid": "testuid"
			}]}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	res, err := apiV2.ListConnections(context.Background(), &client.ListConnectionsRequest{
		TeamId: os.Getenv("MAKE_TEAM_ID"),
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	if (res != nil) {
		assert.GreaterOrEqual(t, len(res.Connections), 1, "expecting at least 1 connection")
	}
}
