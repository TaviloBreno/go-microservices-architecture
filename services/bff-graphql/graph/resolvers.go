package graph

import (
	"context"
	"log"
	"strconv"

	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/graph/model"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/clients"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/config"
)

// Resolver estrutura principal contendo todos os clientes gRPC
type Resolver struct {
	OrderClient        clients.OrderClient
	UserClient         clients.UserClient
	PaymentClient      clients.PaymentClient
	NotificationClient clients.NotificationClient
	Config             *config.Config
}

// NewResolver cria um novo resolver com todos os clientes inicializados
func NewResolver(cfg *config.Config) (*Resolver, error) {
	log.Println("🚀 Inicializando clientes gRPC...")

	// Inicializar cliente Order Service
	orderClient, err := clients.NewOrderClient(cfg.OrderGRPCAddr, cfg.GRPCTimeout)
	if err != nil {
		log.Printf("❌ Erro ao inicializar Order Client: %v", err)
		return nil, err
	}

	// Inicializar cliente User Service
	userClient, err := clients.NewUserClient(cfg.UserGRPCAddr, cfg.GRPCTimeout)
	if err != nil {
		log.Printf("❌ Erro ao inicializar User Client: %v", err)
		orderClient.Close()
		return nil, err
	}

	// Inicializar cliente Payment Service
	paymentClient, err := clients.NewPaymentClient(cfg.PaymentGRPCAddr, cfg.GRPCTimeout)
	if err != nil {
		log.Printf("❌ Erro ao inicializar Payment Client: %v", err)
		orderClient.Close()
		userClient.Close()
		return nil, err
	}

	// Inicializar cliente Notification Service
	notificationClient, err := clients.NewNotificationClient(cfg.NotificationGRPCAddr, cfg.GRPCTimeout)
	if err != nil {
		log.Printf("❌ Erro ao inicializar Notification Client: %v", err)
		orderClient.Close()
		userClient.Close()
		paymentClient.Close()
		return nil, err
	}

	log.Println("✅ Todos os clientes gRPC inicializados com sucesso!")

	return &Resolver{
		OrderClient:        orderClient,
		UserClient:         userClient,
		PaymentClient:      paymentClient,
		NotificationClient: notificationClient,
		Config:             cfg,
	}, nil
}

// Close fecha todas as conexões gRPC
func (r *Resolver) Close() error {
	log.Println("🔌 Fechando todas as conexões gRPC...")

	if r.OrderClient != nil {
		r.OrderClient.Close()
	}
	if r.UserClient != nil {
		r.UserClient.Close()
	}
	if r.PaymentClient != nil {
		r.PaymentClient.Close()
	}
	if r.NotificationClient != nil {
		r.NotificationClient.Close()
	}

	return nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error) {
	log.Printf("📦 GraphQL: Criando pedido para usuário %d", input.UserID)

	// Converter para o formato do cliente gRPC
	orderReq := &clients.OrderRequest{
		Customer:  strconv.Itoa(input.UserID), // Converter userID para string como customer
		ProductID: 1,                          // Assumindo produto padrão
		Quantity:  int32(input.Quantity),
		Price:     input.Price,
	}

	// Chamar o serviço via gRPC
	orderResp, err := r.OrderClient.CreateOrder(ctx, orderReq)
	if err != nil {
		log.Printf("❌ Erro ao criar pedido via gRPC: %v", err)
		return nil, err
	}

	// Converter resposta para modelo GraphQL
	return &model.Order{
		ID:          strconv.Itoa(int(orderResp.ID)),
		UserID:      input.UserID,
		ProductName: input.ProductName,
		Quantity:    input.Quantity,
		Price:       orderResp.Price,
		Status:      "created",
		CreatedAt:   orderResp.CreatedAt,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	log.Printf("👤 GraphQL: Criando usuário %s", input.Name)

	// Converter para o formato do cliente gRPC
	userReq := &clients.UserRequest{
		Name:  input.Name,
		Email: input.Email,
	}

	// Chamar o serviço via gRPC
	userResp, err := r.UserClient.CreateUser(ctx, userReq)
	if err != nil {
		log.Printf("❌ Erro ao criar usuário via gRPC: %v", err)
		return nil, err
	}

	// Converter resposta para modelo GraphQL
	return &model.User{
		ID:        strconv.Itoa(int(userResp.ID)),
		Name:      userResp.Name,
		Email:     userResp.Email,
		CreatedAt: userResp.CreatedAt,
	}, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	log.Printf("📋 GraphQL: Listando pedidos")

	// Chamar o serviço via gRPC
	orders, err := r.OrderClient.ListOrders(ctx)
	if err != nil {
		log.Printf("❌ Erro ao listar pedidos via gRPC: %v", err)
		return nil, err
	}

	// Converter resposta para modelo GraphQL
	result := make([]*model.Order, len(orders))
	for i, order := range orders {
		userID, _ := strconv.Atoi(order.Customer) // Converter customer de volta para userID
		result[i] = &model.Order{
			ID:          strconv.Itoa(int(order.ID)),
			UserID:      userID,
			ProductName: "Produto " + strconv.Itoa(int(order.ProductID)), // Nome baseado no ID
			Quantity:    int(order.Quantity),
			Price:       order.Price,
			Status:      "completed",
			CreatedAt:   order.CreatedAt,
		}
	}

	return result, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Printf("👥 GraphQL: Listando usuários")

	// Chamar o serviço via gRPC
	users, err := r.UserClient.ListUsers(ctx)
	if err != nil {
		log.Printf("❌ Erro ao listar usuários via gRPC: %v", err)
		return nil, err
	}

	// Converter resposta para modelo GraphQL
	result := make([]*model.User, len(users))
	for i, user := range users {
		result[i] = &model.User{
			ID:        strconv.Itoa(int(user.ID)),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	return result, nil
}

// Payments is the resolver for the payments field.
func (r *queryResolver) Payments(ctx context.Context) ([]*model.Payment, error) {
	log.Printf("💳 GraphQL: Listando pagamentos")

	// Como o PaymentService só tem GetPaymentStatus por orderID,
	// vamos simular uma lista baseada nos pedidos conhecidos
	orders, err := r.OrderClient.ListOrders(ctx)
	if err != nil {
		log.Printf("❌ Erro ao listar pedidos para pagamentos: %v", err)
		return nil, err
	}

	result := make([]*model.Payment, 0)
	for _, order := range orders {
		// Tentar obter status do pagamento para cada pedido
		payment, err := r.PaymentClient.GetPaymentStatus(ctx, order.ID)
		if err != nil {
			log.Printf("⚠️ Erro ao obter pagamento do pedido %d: %v", order.ID, err)
			continue
		}

		userID, _ := strconv.Atoi(order.Customer)
		result = append(result, &model.Payment{
			ID:            strconv.Itoa(int(payment.OrderID)), // Usando orderID como ID do payment
			OrderID:       int(payment.OrderID),
			UserID:        userID,
			Amount:        payment.Amount,
			Status:        payment.Status,
			PaymentMethod: "card", // Valor padrão
			CreatedAt:     order.CreatedAt,
		})
	}

	return result, nil
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context) ([]*model.Notification, error) {
	log.Printf("📢 GraphQL: Listando notificações")

	// Chamar o serviço via gRPC
	notifications, err := r.NotificationClient.ListNotifications(ctx)
	if err != nil {
		log.Printf("❌ Erro ao listar notificações via gRPC: %v", err)
		return nil, err
	}

	// Converter resposta para modelo GraphQL
	result := make([]*model.Notification, len(notifications))
	for i, notification := range notifications {
		result[i] = &model.Notification{
			ID:        strconv.Itoa(int(notification.ID)),
			PaymentID: int(notification.PaymentID),
			OrderID:   int(notification.OrderID),
			Message:   notification.Message,
			Status:    notification.Status,
			CreatedAt: notification.CreatedAt,
		}
	}

	return result, nil
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	log.Printf("📦 GraphQL: Buscando pedido ID %s", id)

	// Para simplicidade, listar todos e filtrar (em produção seria uma busca direta)
	orders, err := r.OrderClient.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	orderIDInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		if order.ID == uint32(orderIDInt) {
			userID, _ := strconv.Atoi(order.Customer)
			return &model.Order{
				ID:          strconv.Itoa(int(order.ID)),
				UserID:      userID,
				ProductName: "Produto " + strconv.Itoa(int(order.ProductID)),
				Quantity:    int(order.Quantity),
				Price:       order.Price,
				Status:      "completed",
				CreatedAt:   order.CreatedAt,
			}, nil
		}
	}

	return nil, nil // Pedido não encontrado
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Printf("👤 GraphQL: Buscando usuário ID %s", id)

	// Para simplicidade, listar todos e filtrar
	users, err := r.UserClient.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	userIDInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == uint32(userIDInt) {
			return &model.User{
				ID:        strconv.Itoa(int(user.ID)),
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
			}, nil
		}
	}

	return nil, nil // Usuário não encontrado
}

// Payment is the resolver for the payment field.
func (r *queryResolver) Payment(ctx context.Context, id string) (*model.Payment, error) {
	log.Printf("💳 GraphQL: Buscando pagamento ID %s", id)

	orderID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	payment, err := r.PaymentClient.GetPaymentStatus(ctx, uint32(orderID))
	if err != nil {
		return nil, err
	}

	return &model.Payment{
		ID:            strconv.Itoa(int(payment.OrderID)),
		OrderID:       int(payment.OrderID),
		UserID:        0, // Seria necessário consultar o pedido para obter o userID
		Amount:        payment.Amount,
		Status:        payment.Status,
		PaymentMethod: "card",
		CreatedAt:     "",
	}, nil
}

// OrderSummary is the resolver for the orderSummary field.
func (r *queryResolver) OrderSummary(ctx context.Context, orderID string) (*model.OrderSummary, error) {
	log.Printf("📊 GraphQL: Buscando resumo do pedido ID %s", orderID)

	// Buscar pedido
	order, err := r.Order(ctx, orderID)
	if err != nil || order == nil {
		return nil, err
	}

	// Buscar usuário
	user, err := r.User(ctx, strconv.Itoa(order.UserID))
	if err != nil {
		user = nil // Continuar mesmo se não encontrar usuário
	}

	// Buscar pagamento
	payment, err := r.Payment(ctx, orderID)
	if err != nil {
		payment = nil // Continuar mesmo se não encontrar pagamento
	}

	// Buscar notificações do pedido
	allNotifications, err := r.Notifications(ctx)
	if err != nil {
		allNotifications = nil
	}

	// Filtrar notificações do pedido específico
	orderNotifications := make([]*model.Notification, 0)
	if allNotifications != nil {
		orderIDInt, _ := strconv.Atoi(orderID)
		for _, notification := range allNotifications {
			if notification.OrderID == orderIDInt {
				orderNotifications = append(orderNotifications, notification)
			}
		}
	}

	return &model.OrderSummary{
		Order:         order,
		User:          user,
		Payment:       payment,
		Notifications: orderNotifications,
	}, nil
}

// Health is the resolver for the health field.
func (r *queryResolver) Health(ctx context.Context) (string, error) {
	log.Printf("🏥 GraphQL: Health check")
	return "BFF GraphQL is healthy! 🚀", nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
