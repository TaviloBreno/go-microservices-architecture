package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := ":50053"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

	fmt.Printf("🚀 OrderService rodando em gRPC %s\n", port)

	s := grpc.NewServer()
	// TODO: registrar serviços gRPC aqui

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Erro ao servir: %v", err)
	}
}
