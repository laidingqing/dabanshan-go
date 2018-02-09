package service

import (
	"context"

	"github.com/laidingqing/dabanshan/notifications/model"
	"github.com/laidingqing/dabanshan/pb"
)

// RPCNotificationServer is used to implement notification_service.UserServiceServer.
type RPCNotificationServer struct{}

//Send send a notifaction message.
func (s *RPCNotificationServer) Send(ctx context.Context, in *pb.NotificationRequest) (*pb.NotificationReply, error) {
	notification := model.PushNotification{
		Channel:          int(in.Channel),
		Tokens:           in.Tokens,
		Message:          in.Message,
		Title:            in.Title,
		Sound:            in.Sound,
		ContentAvailable: in.ContentAvailable,
	}

	if in.Alert != nil {
		notification.Alert = model.Alert{
			Title:        in.Alert.Title,
			Body:         in.Alert.Body,
			Subtitle:     in.Alert.Subtitle,
			Action:       in.Alert.Action,
			ActionLocKey: in.Alert.Action,
			LaunchImage:  in.Alert.LaunchImage,
			LocArgs:      in.Alert.LocArgs,
			LocKey:       in.Alert.LocKey,
			TitleLocArgs: in.Alert.TitleLocArgs,
			TitleLocKey:  in.Alert.TitleLocKey,
		}
	}

	go SendNotification(notification)

	return &pb.NotificationReply{
		Success: false,
		Counts:  int32(len(notification.Tokens)),
	}, nil
}
