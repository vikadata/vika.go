// Package http provides basic http request/response method and definition
package http

import (
	"encoding/json"
	"fmt"
	vkerror "github.com/vikadata/vika.go/lib/common/error"
	"io/ioutil"
	//"log"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string
}

type ErrorResponse struct {
	BaseResponse
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return vkerror.NewVikaSDKError(500, msg, "ClientError.ParseJsonError")
	}
	if resp.Code != 200 {
		return vkerror.NewVikaSDKError(resp.Code, resp.Message, resp.Message)
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return vkerror.NewVikaSDKError(500, msg, "ClientError.IOError")
	}
	if !(hr.StatusCode == 200 || hr.StatusCode == 201) {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return vkerror.NewVikaSDKError(hr.StatusCode, msg, "ClientError.HttpStatusCodeError")
	}
	//log.Printf("[DEBUG] Response Body=%s", body)
	err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return vkerror.NewVikaSDKError(500, msg, "ClientError.ParseJsonError")
	}
	return
}
