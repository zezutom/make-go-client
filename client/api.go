package client

import (
	"go-make/model"
)

type Api struct {
	Version model.Version
	Client  Client
}

type ApiV2 Api

func (c *Client) NewApiV2() *ApiV2 {
	return &ApiV2{
		Version: model.V2,
		Client:  *c,
	}
}
