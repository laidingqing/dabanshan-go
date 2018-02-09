package model

import (
	"sync"

	"gopkg.in/mgo.v2/bson"
)

// D provide string array
type D map[string]interface{}

//NotifactionChannel 通知通道
type NotifactionChannel int

const (
	//ChannelEmail 邮件
	ChannelEmail = iota + 1
	//ChannelMessage 短信
	ChannelMessage
	//ChannelWeChat 微信
	ChannelWeChat
)

// LogPushEntry is push response log
type LogPushEntry struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Type    string        `bson:"type" json:"type"`
	Channel string        `bson:"channel" json:"channel"`
	Token   string        `bson:"token" json:"token"`
	Message string        `bson:"message" json:"message"`
	Error   string        `bson:"error" json:"error"`
}

// Alert is APNs payload
type Alert struct {
	Action       string   `json:"action,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	Body         string   `json:"body,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	Title        string   `json:"title,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
}

// RequestPush support multiple notification request.
type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
}

// PushNotification is single notification request
type PushNotification struct {
	Tokens           []string `json:"tokens" binding:"required"`
	TemplateID       string   `json:"template_id" binding:"required"`
	Channel          int      `json:"channel" binding:"required"`
	Message          string   `json:"message,omitempty"`
	Title            string   `json:"title,omitempty"`
	Priority         string   `json:"priority,omitempty"`
	ContentAvailable bool     `json:"content_available,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	Data             D        `json:"data,omitempty"`
	Retry            int      `json:"retry,omitempty"`
	Alert            Alert    `json:"alert,omitempty"`
	wg               *sync.WaitGroup
	log              *[]LogPushEntry
}
