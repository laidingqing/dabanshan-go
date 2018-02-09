package service

import (
	"log"

	"github.com/laidingqing/dabanshan/notifications/model"
)

var (
	// QueueNotification is chan type
	QueueNotification chan model.PushNotification
)

// InitWorkers for initialize all workers.
func InitWorkers(workerNum int64, queueNum int64) {
	log.Printf("worker number is %v , queue number is %v", workerNum, queueNum)
	QueueNotification = make(chan model.PushNotification, queueNum)
	for i := int64(0); i < workerNum; i++ {
		go startWorker()
	}
}

//SendNotification send a notification
func SendNotification(msg model.PushNotification) {
	switch msg.Channel {
	case model.ChannelEmail:
		// PUSH to WEIXIN
	case model.ChannelMessage:
		// PUSH to EMAIL
	case model.ChannelWeChat:
		// PUSH to Wechat
	}
}

func startWorker() {
	for {
		notification := <-QueueNotification
		SendNotification(notification)
	}
}
