package common

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/aiden0z/qcloud-go/util"
	simplejson "github.com/bitly/go-simplejson"
)

const (
	QcloudHTTPMehtod = "POST"
)

// Client descrips Qcloud service client
type Client struct {
	SecretId   string
	SecretKey  string
	Region     string
	debug      bool
	httpClient *http.Client
	endpoint   string
}

// Init the client
func (c *Client) Init(endpoint, region, secretId, secretKey string) {
	c.endpoint = endpoint
	c.Region = region
	c.SecretId = secretId
	c.SecretKey = secretKey
	c.httpClient = &http.Client{}
}

// Invoke a action
func (c *Client) Invoke(action string, args map[string]interface{}) (js *simplejson.Json, err error) {
	args["SecretId"] = c.SecretId
	args["Timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	args["Nonce"] = fmt.Sprintf("%v", rand.Int())
	args["Region"] = c.Region
	args["Action"] = action

	signature, err := util.Sign(QcloudHTTPMehtod, c.endpoint, args, c.SecretKey)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("Signature", signature)

	for k, v := range args {
		params.Add(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}

	url := "https://" + c.endpoint
	resp, err := c.httpClient.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	js, err = simplejson.NewJson(data)

	if err != nil {
		return nil, err
	}

	return js, nil
}

// Do invoke the action and determine the error code
func (c *Client) Do(action string, args map[string]interface{}) (js *simplejson.Json, err error) {

	js, err = c.Invoke(action, args)
	if err != nil {
		return
	}

	code, err := js.Get("code").Int()
	if err != nil {
		return
	}

	message, err := js.Get("message").String()
	if nil != err {
		return
	}

	if code != 0 {
		err = fmt.Errorf("qlcoud api eror, code: %d, message: %s", code, message)
		return
	}

	return
}
