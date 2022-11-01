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
		Metadata: map[string]interface{}{
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

func Test_CreateConnection(t *testing.T) {
	json := `{
			   "connection":{
				  "id":764265,
				  "name":"Slack Test",
				  "accountName":"slack",
				  "accountLabel":"Slack",
				  "packageName":null,
				  "expire":null,
				  "metadata":null,
				  "teamId":1,
				  "theme":"#4a154b",
				  "upgradeable":false,
				  "scopes":0,
				  "scoped":true,
				  "accountType":"oauth",
				  "editable":false,
				  "uid":null
			   }
			}`
	res, err := mock.Success(json).CreateConnection(context.Background(), &client.CreateConnectionReq{
		TeamId: "1",
		Body: 	client.CreateConnectionReqBody{
			AccountName:  "Slack Test",
			AccountType:  "slack",
			Scopes:       []string{"chat:write"},
			ClientId:     "123456",
			ClientSecret: "clientSecret",
		},
	})
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(t, client.Connection{
		ID:           764265,
		Name:         "Slack Test",
		AccountName:  "slack",
		AccountLabel: "Slack",
		PackageName:  "",
		Expire:       nil,
		Metadata: nil,
		TeamID:      1,
		Theme:       "#4a154b",
		Upgradeable: false,
		Scopes:      0,
		Scoped:      true,
		AccountType: "oauth",
		Editable:    false,
		UID:         "",
	}, res.Connection, "expecting the connection be equal to the mock connection")
}
