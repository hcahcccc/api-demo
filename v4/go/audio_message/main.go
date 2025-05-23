package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://api-audio-sh.fengkongcloud.com/audiomessage/v4"
	accessKey := "{ACCESS_KEY}"
	btId := "{BT_ID}"
	filename := "../files/demo.pcm"

	contentBytes, _ := ioutil.ReadFile(filename)
	content := base64.StdEncoding.EncodeToString(contentBytes)
	payload := map[string]interface{}{
		"accessKey":   accessKey,
		"appId":       "default",
		"eventId":     "audio",
		"type":        "POLITY_EROTIC_MOAN_ADVERT",
		"content":     content,
		"contentType": "RAW",
		"btId":        btId,
		"data": map[string]interface{}{
			"format": "pcm",
			"rate":   8000,
			"track":  1,
		},
	}
	b, _ := json.Marshal(payload)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(b))

	var data map[string]interface{}
	if resp != nil {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(respBytes, &data)
		// 通过 code 和 riskLevel 判断返回，具体请参考接口文档
		fmt.Println(data)
	}
}
