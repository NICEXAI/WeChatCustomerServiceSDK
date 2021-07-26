package WeChatCustomerServiceSDK

import (
	"WeChatCustomerServiceSDK/util"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	//获取调用凭证AccessToken
	getTokenAddr = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

// AccessTokenSchema 获取调用凭证响应数据
type AccessTokenSchema struct {
	BaseModel
	AccessToken string `json:"access_token"`				// 获取到的凭证，最长为512字节
	ExpiresIn int `json:"expires_in"`						// 凭证的有效时间（秒）
}

// GetAccessToken 获取调用凭证access_token
func (r *Client) GetAccessToken() (info AccessTokenSchema, err error) {
	data, err := util.HttpGet(fmt.Sprintf(getTokenAddr, r.corpID, r.secret))
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	fmt.Println(string(data))
	if info.ErrCode != 0 {
		return info, errors.New(info.ErrMsg)
	}
	return info, nil
}

func (r *Client) getAccessToken() (string, error) {
	return r.cache.Get("wechat:kf:" + r.corpID)
}

func (r *Client) setAccessToken(token string) error {
	return r.cache.Set("wechat:kf:" + r.corpID, token, r.expireTime)
}
