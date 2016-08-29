package cvm

import (
	"github.com/aiden0z/qcloud-go/common"
)

const (
	// CVMDefaultEndpoint describes the cvm api endpoint
	CVMDefaultEndpoint = "cvm.api.qcloud.com/v2/index.php"
)

// Client describes the cvm api client
type Client struct {
	common.Client
}

// NewClient return a cvm api client
func NewClient(region, secretId, secretKey string) *Client {
	c := &Client{}
	c.Init(CVMDefaultEndpoint, region, secretId, secretKey)
	return c
}
