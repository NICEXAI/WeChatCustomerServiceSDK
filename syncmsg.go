package WeChatCustomerServiceSDK

import (
	"WeChatCustomerServiceSDK/util"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	//获取消息
	syncMsgAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg?access_token=%s"
)

// SyncMsgOptions 获取消息查询参数
type SyncMsgOptions struct {
	Cursor string `json:"cursor"`							// 上一次调用时返回的next_cursor，第一次拉取可以不填, 不多于64字节
	Token string `json:"token"`								// 回调事件返回的token字段，10分钟内有效；可不填，如果不填接口有严格的频率限制, 不多于128字节
	Limit uint `json:"limit"`								// 期望请求的数据量，默认值和最大值都为1000, 注意：可能会出现返回条数少于limit的情况，需结合返回的has_more字段判断是否继续请求。
}

// SyncMsgSchema 获取消息查询响应内容
type SyncMsgSchema struct {
	ErrCode int32 `json:"errcode"`							// 返回码
	ErrMsg string `json:"errmsg"`							// 错误码描述
	NextCursor string `json:"next_cursor"`					// 下次调用带上该值则从该key值往后拉，用于增量拉取
	HasMore uint32 `json:"has_more"`						// 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList [][]byte `json:"msg_list"`						// 消息列表
}

// SyncMsg 获取消息
func (r *Client) SyncMsg(options SyncMsgOptions) (info SyncMsgSchema, err error) {
	data, err := util.HttpPost(fmt.Sprintf(syncMsgAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, errors.New(info.ErrMsg)
	}
	return info, nil
}