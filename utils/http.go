package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func HttpRequest(method, url string, headers, queryParam map[string]string, body interface{}, timeout int) (resp *resty.Response, err error) {

	client := resty.New().SetTimeout(time.Duration(timeout) * time.Millisecond)

	var r = client.R().SetHeaders(headers).SetQueryParams(queryParam)
	switch method {
	case http.MethodGet:
		resp, err = r.Get(url)
	case http.MethodPost:
		resp, err = r.SetBody(body).Post(url)
	case http.MethodPatch:
		resp, err = r.SetBody(body).Patch(url)
	case http.MethodDelete:
		resp, err = r.SetBody(body).Delete(url)
	case http.MethodPut:
		resp, err = r.SetBody(body).Put(url)
	}

	// // marshal body for logging
	// bodyByte, _ := json.Marshal(body)
	// headerByte, _ := json.Marshal(headers)
	// queryByte, _ := json.Marshal(queryParam)

	if resp == nil || err != nil {
		fmt.Println(err)
		// Logger.LogHttpRequest(logId, method, url, fmt.Sprintf("%s", r.Header), string(bodyByte), string(queryByte))
		// Logger.LogError(logId, err.Error())
		return
	}

	bodyResBuff := &bytes.Buffer{}
	if err := json.Compact(bodyResBuff, resp.Body()); err != nil {
		fmt.Println(err)
	}

	// Logger.LogHttpRequest(logId, method, url, string(headerByte), string(bodyByte), string(queryByte))
	// Logger.LogHttpResponse(logId, method, url, resp.StatusCode(), bodyResBuff.String(), resp.Time().String())

	return
}
