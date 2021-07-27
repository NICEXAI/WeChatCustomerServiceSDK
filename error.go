package WeChatCustomerServiceSDK

import "fmt"

// Error 错误
type Error struct {
	ErrCode int `json:"err_code,omitempty"`
	ErrMsg string `json:"err_msg"`
}

//输出错误信息
func (r Error) Error() string {
	return fmt.Sprintf("%d:%s", r.ErrCode, r.ErrMsg)
}

// NewSDKErr 初始化SDK实例错误信息
func NewSDKErr(code int, msg string) Error {
	return Error{
		ErrCode: code,
		ErrMsg:  msg,
	}
}