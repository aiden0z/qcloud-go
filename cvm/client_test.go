package cvm

import "testing"

const (
	Region    = "sh"
	SecretId  = "your secret id"
	SecretKey = "your secret key"
)

func TestClient(t *testing.T) {
	args := map[string]interface{}{"offset": 0, "limit": 1}
	client := NewClient(Region, SecretId, SecretKey)
	js, err := client.Do("DescribeInstances", args)
	if err != nil {
		t.Error(err)
		prettyData, _ := js.EncodePretty()
		t.Log(string(prettyData))
	}
}
