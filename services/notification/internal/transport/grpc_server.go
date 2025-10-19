package transport

import (
	"context"

	"notification-service/internal/service"
	pb "notification-service/proto"
)

// NotificationGRPCServer implementa o servidor gRPC do notification service
type NotificationGRPCServer struct {
	pb.UnimplementedNotificationServiceServer
	notificationService service.NotificationService
}

// NewNotificationGRPCServer cria uma nova instância do servidor gRPC
func NewNotificationGRPCServer(notificationService service.NotificationService) *NotificationGRPCServer {
	return &NotificationGRPCServer{
		notificationService: notificationService,
	}
}

// ListNotifications retorna todas as notificações
func (s *NotificationGRPCServer) ListNotifications(ctx context.Context, req *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	notifications, err := s.notificationService.GetAllNotifications()
	if err != nil {
		return nil, err
	}

	var pbNotifications []*pb.NotificationResponse
	for _, notification := range notifications {
		pbNotification := &pb.NotificationResponse{
			Id:        uint32(notification.ID),
			PaymentId: uint32(notification.PaymentID),
			OrderId:   uint32(notification.OrderID),
			Message:   notification.Message,
			Status:    notification.Status,
			CreatedAt: notification.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		pbNotifications = append(pbNotifications, pbNotification)
	}

	return &pb.ListNotificationsResponse{
		Notifications: pbNotifications,
	}, nil
}
