package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sort"
)

func makePlainText(method, endpoint string, params map[string]interface{}) (plainText string, err error) {

	plainText += method
	plainText += endpoint
	plainText += "?"

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var plainParms string
	for i := range keys {
		k := keys[i]
		plainParms += "&" + fmt.Sprintf("%v", k) + "=" + fmt.Sprintf("%v", params[k])
	}
	plainText += plainParms[1:]

	return plainText, nil
}

// Sign implement the qcloud sign method, refer https://www.qcloud.com/doc/api/229/1227
func Sign(method, endpoint string, params map[string]interface{}, secretKey string) (sign string, err error) {

	var source string

	source, err = makePlainText(method, endpoint, params)
	if err != nil {
		return sign, err
	}

	obj := hmac.New(sha1.New, []byte(secretKey))
	obj.Write([]byte(source))

	sign = base64.StdEncoding.EncodeToString(obj.Sum(nil))

	return sign, nil
}
