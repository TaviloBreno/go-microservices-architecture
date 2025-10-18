package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("🧩 User Service rodando na porta 50051...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
