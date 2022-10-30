package mock

import (
	"bytes"
	"go-make/client"
	"go-make/model"
	"io/ioutil"
	"net/http"
)

type MockClient struct {
	GetDoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error)  {
	return m.GetDoFunc(req)
}

func Success(json string) *client.ApiV2 {
	return client.NewClient("TEST_TOKEN", model.EU, &MockClient{
		GetDoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}).NewApiV2()
}