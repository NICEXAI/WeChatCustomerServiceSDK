package WeChatCustomerServiceSDK

import (
	"github.com/NICEXAI/WeChatCustomerServiceSDK/cache"
	"sync"
	"time"
)

// BaseModel 基础数据
type BaseModel struct {
	ErrCode int    `json:"errcode"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg"`  // 返回码提示语
}

// Options 微信客服初始化参数
type Options struct {
	CorpID         string        // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	Secret         string        // Secret是微信客服用于校验开发者身份的访问密钥，企业成功注册微信客服后，可在「微信客服管理后台-开发配置」处获取
	Token          string        // 用于生成签名校验回调请求的合法性
	EncodingAESKey string        // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	Cache          cache.Cache   // 数据缓存
	ExpireTime     time.Duration // 令牌过期时间
	IsCloseCache   bool          // 是否关闭自动缓存AccessToken, 默认缓存
}

// Client 微信客服实例
type Client struct {
	corpID         string        // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	secret         string        // Secret是微信客服用于校验开发者身份的访问密钥，企业成功注册微信客服后，可在「微信客服管理后台-开发配置」处获取
	token          string        // 用于生成签名校验回调请求的合法性
	encodingAESKey string        // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	expireTime     time.Duration // 令牌过期时间
	cache          cache.Cache
	eventQueue     sync.Map //事件队列
	mutex          sync.Mutex
	accessToken    string // 用户访问凭证
	isCloseCache   bool   // 是否自动缓存AccessToken, 默认缓存
}

// New 初始化微信客服实例
func New(options Options) (client *Client, err error) {
	if options.Cache == nil {
		return nil, NewSDKErr(50001)
	}

	if options.ExpireTime == 0 {
		options.ExpireTime = 6000
	}

	client = &Client{
		corpID:         options.CorpID,
		secret:         options.Secret,
		token:          options.Token,
		encodingAESKey: options.EncodingAESKey,
		expireTime:     options.ExpireTime,
		cache:          options.Cache,
		eventQueue:     sync.Map{},
		mutex:          sync.Mutex{},
		isCloseCache:   options.IsCloseCache,
	}

	if options.Secret != "" {
		if err = client.initAccessToken(); err != nil {
			return nil, err
		}
	}

	return client, nil
}
