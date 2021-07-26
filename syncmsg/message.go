package syncmsg

// BaseMessage 接收消息
type BaseMessage struct {
	MsgID string `json:"msgid"`							// 消息ID
	OpenKFID string `json:"open_kfid"`					// 客服帐号ID
	ExternalUserID string `json:"external_userid"`		// 客户UserID
	SendTime uint64 `json:"send_time"`					// 消息发送时间
	Origin uint32 `json:"origin"`						// 消息来源。3-客户回复的消息 4-系统推送的消息
}

// Text 文本消息
type Text struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：text
	Text struct{
		Content string `json:"content"`					// 文本内容
		MenuID string `json:"menu_id"`					// 客户点击菜单消息，触发的回复消息中附带的菜单ID
	} `json:"text"`										// 文本消息
}

// Image 图片消息
type Image struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：image
	Image struct{
		MediaID string `json:"media_id"`				// 图片文件ID
	} `json:"image"`									// 图片消息
}

// Voice 语音消息
type Voice struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：voice
	Voice struct{
		MediaID string `json:"media_id"`				// 语音文件ID
	} `json:"voice"`									// 语音消息
}

// Video 视频消息
type Video struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：video
	Video struct{
		MediaID string `json:"media_id"`				// 文件ID
	} `json:"video"`									// 视频消息
}

// File 文件消息
type File struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：file
	File struct{
		MediaID string `json:"media_id"`				// 文件ID
	} `json:"file"`										// 文件消息
}

// Location 地理位置消息
type Location struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：location
	Location struct{
		Latitude float32 `json:"latitude"`				// 纬度
		Longitude float32 `json:"longitude"`			// 经度
		Name string `json:"name"`						// 位置名
		Address string `json:"address"`					// 地址详情说明
	} `json:"location"`									// 地理位置消息
}

// EventMessage 事件消息
type EventMessage struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：event
	Event struct{
		EventType string `json:"event_type"`			// 事件类型
	} `json:"event"`									// 事件消息
}

// EnterSessionEvent 用户进入会话事件
type EnterSessionEvent struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：event
	Event struct{
		EventType string `json:"event_type"`			// 事件类型。此处固定为：enter_session
		OpenKFID string `json:"open_kfid"`				// 客服账号ID
		ExternalUserID string `json:"external_userid"`	// 客户UserID
		Scene string `json:"scene"`						// 进入会话的场景值，获取客服帐号链接开发者自定义的场景值
	} `json:"event"`									// 事件消息
}

// MsgSendFailEvent 消息发送失败事件
type MsgSendFailEvent struct {
	BaseMessage
	MsgType string `json:"msgtype"`						// 消息类型，此时固定为：event
	Event struct{
		EventType string `json:"event_type"`			// 事件类型。此处固定为：msg_send_fail
		OpenKFID string `json:"open_kfid"`				// 客服账号ID
		ExternalUserID string `json:"external_userid"`	// 客户UserID
		FailMsgID string `json:"fail_msgid"`			// 发送失败的消息msgid
		FailType uint32 `json:"fail_type"`				// 失败类型。0-未知原因 10-用户拒收
	} `json:"event"`									// 事件消息
}