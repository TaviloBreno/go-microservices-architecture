package clients

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UserRequest representa uma requisição de criação de usuário
type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse representa um usuário
type UserResponse struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// ListUsersRequest representa uma requisição para listar usuários
type ListUsersRequest struct{}

// ListUsersResponse representa a resposta com lista de usuários
type ListUsersResponse struct {
	Users []*UserResponse `json:"users"`
}

// UserClient interface para comunicação com user-service
type UserClient interface {
	CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error)
	ListUsers(ctx context.Context) ([]*UserResponse, error)
	Close() error
}

// UserServiceClient interface gRPC (simplificada)
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

// userServiceClient implementa UserServiceClient
type userServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewUserServiceClient cria um novo cliente gRPC para UserService
func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// grpcUserClient implementa UserClient usando gRPC
type grpcUserClient struct {
	conn   *grpc.ClientConn
	client UserServiceClient
}

// NewUserClient cria um novo cliente para user-service
func NewUserClient(grpcAddr string, timeout time.Duration) (UserClient, error) {
	log.Printf("🔌 Conectando ao User Service em %s...", grpcAddr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Printf("❌ Erro ao conectar com User Service: %v", err)
		return nil, err
	}

	client := NewUserServiceClient(conn)
	log.Printf("✅ Conectado ao User Service com sucesso!")

	return &grpcUserClient{
		conn:   conn,
		client: client,
	}, nil
}

// CreateUser cria um novo usuário
func (c *grpcUserClient) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	log.Printf("👤 Criando usuário: %+v", req)

	response, err := c.client.CreateUser(ctx, req)
	if err != nil {
		log.Printf("❌ Erro ao criar usuário: %v", err)
		return nil, err
	}

	log.Printf("✅ Usuário criado com sucesso: ID %d", response.ID)
	return response, nil
}

// ListUsers lista todos os usuários
func (c *grpcUserClient) ListUsers(ctx context.Context) ([]*UserResponse, error) {
	log.Printf("👥 Listando usuários...")

	response, err := c.client.ListUsers(ctx, &ListUsersRequest{})
	if err != nil {
		log.Printf("❌ Erro ao listar usuários: %v", err)
		return nil, err
	}

	log.Printf("✅ Encontrados %d usuários", len(response.Users))
	return response.Users, nil
}

// Close fecha a conexão
func (c *grpcUserClient) Close() error {
	log.Printf("🔌 Fechando conexão com User Service...")
	return c.conn.Close()
}
