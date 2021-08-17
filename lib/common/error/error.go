// Package error provides custom sdk error
package error

import "fmt"

type VikaSDKError struct {
	Code      int
	Message   string
	RequestId string
}

func (e *VikaSDKError) Error() string {
	return fmt.Sprintf("[SDKError] Code=%d, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewVikaSDKError(code int, message, requestId string) error {
	return &VikaSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *VikaSDKError) GetCode() int {
	return e.Code
}

func (e *VikaSDKError) GetMessage() string {
	return e.Message
}

func (e *VikaSDKError) GetRequestId() string {
	return e.RequestId
}
