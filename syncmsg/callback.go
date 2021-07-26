package syncmsg

type Event struct {
	ToUserName string 					// 微信客服组件ID
	CreateTime int						// 消息创建时间，unix时间戳
	MsgType	string						// 消息的类型，此时固定为 event
	Event string						// 事件的类型，此时固定为 kf_msg_or_event
	Token string						// 调用拉取消息接口时，需要传此token，用于校验请求的合法性
}
