package syncmsg

import "encoding/json"

// Message 同步的消息内容
type Message struct {
	MsgID string `json:"msgid"`							// 消息ID
	OpenKFID string `json:"open_kfid"`					// 客服帐号ID
	ExternalUserID string `json:"external_userid"`		// 客户UserID
	SendTime uint64 `json:"send_time"`					// 消息发送时间
	Origin uint32 `json:"origin"`						// 消息来源。3-客户回复的消息 4-系统推送的消 息
	MsgType string `json:"msgtype"`						// 消息类型
	EventType string `json:"event_type"`				// 事件类型
	originData []byte 									// 原始数据内容
}

// GetOriginMessage 获取原始消息
func (r Message) GetOriginMessage() (info []byte) {
	return r.originData
}

// GetTextMessage 获取文本消息
func (r Message) GetTextMessage() (info Text, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetImageMessage 获取图片消息
func (r Message) GetImageMessage() (info Image, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetVoiceMessage 获取语音消息
func (r Message) GetVoiceMessage() (info Voice, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetVideoMessage 获取视频消息
func (r Message) GetVideoMessage() (info Video, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetFileMessage 获取文件消息
func (r Message) GetFileMessage() (info File, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetLocationMessage 获取文件消息
func (r Message) GetLocationMessage() (info Location, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetEnterSessionEvent 用户进入会话事件
func (r Message) GetEnterSessionEvent() (info EnterSessionEvent, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}

// GetMsgSendFailEvent 消息发送失败事件
func (r Message) GetMsgSendFailEvent() (info MsgSendFailEvent, err error) {
	err = json.Unmarshal(r.originData, &info)
	return info, err
}