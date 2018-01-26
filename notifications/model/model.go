package model

//NotifactionChannel 通知通道
type NotifactionChannel int

const (
	_ NotifactionChannel = iota
	//EMAIL 邮件
	EMAIL
	//MESSAGE 短信
	MESSAGE
	//WEIXIN 微信
	WEIXIN
)

//NotificationRequest notifaction request.
type NotificationRequest struct {
	To      string             `json:"to"`
	Subject string             `json:"subject"`
	Data    map[string]string  `json:"data"`
	Channel NotifactionChannel `json:"channel"`
}
