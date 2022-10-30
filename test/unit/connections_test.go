package unit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-make/client"
	"go-make/mock"
	"testing"
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
	res, err := mock.Success(json).ListConnections(context.Background(), &client.ListConnectionsReq{
		TeamId: "1",
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(t, len(res.Connections), 1, "expecting 1 connection")
	assert.Equal(t, client.Connection{
		ID:           1,
		Name:         "Test Connection",
		AccountName:  "google",
		AccountLabel: "Google",
		PackageName:  "",
		Expire:       nil,
		Metadata:     map[string]interface{}{
			"type":  "email",
			"value": "j.doe@gmail.com",
		},
		TeamID:      1,
		Theme:       "#fecd5f",
		Upgradeable: false,
		Scopes:      5,
		Scoped:      true,
		AccountType: "oauth",
		Editable:    false,
		UID:         "testuid",
}, res.Connections[0], "expecting the connection be equal to the mock connection")
}
