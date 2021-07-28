package WeChatCustomerServiceSDK

import (
	"reflect"
	"strings"
)

// Error 错误
type Error string

const (
	// SDKInitFailed 错误码：50001
	SDKInitFailed Error			= "SDK初始化失败"
	// SDKCacheUnavailable 错误码：50002
	SDKCacheUnavailable	Error	= "缓存无效"
	// SDKUnknownError 错误码：50003
	SDKUnknownError Error		= "未知错误"
	// SDKInvalidCredential 错误码：40001
	SDKInvalidCredential Error	= "不合法的secret参数"
	// SDKAccessTokenInvalid 错误码：40014
	SDKAccessTokenInvalid Error = "AccessToken 无效"
	// SDKAccessTokenExpired 错误码：42001
	SDKAccessTokenExpired Error = "AccessToken 已过期"
	// SDKWeWorkAlready 错误码：95011
	SDKWeWorkAlready Error 		= "已在企业微信使用微信客服"
)

//输出错误信息
func (r Error) Error() string {
	return reflect.ValueOf(r).String()
}

// NewSDKErr 初始化SDK实例错误信息
func NewSDKErr(code int, msgList ...string) Error {
	switch code {
	case 50001:
		return SDKInitFailed
	case 50002:
		return SDKCacheUnavailable
	case 40001:
		return SDKInvalidCredential
	case 42001:
		return SDKAccessTokenExpired
	case 40014:
		return SDKAccessTokenInvalid
	case 95011:
		return SDKWeWorkAlready
	default:
		//返回未知的自定义错误
		if len(msgList) > 0 {
			return Error(strings.Join(msgList, ","))
		}
		return SDKUnknownError
	}
}