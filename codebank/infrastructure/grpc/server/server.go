package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hrsniper/imersao-fullstack-fullcycle-3/infrastructure/grpc/pb"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/infrastructure/grpc/service"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/usecase"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

// method of GRPCServer struct
func (server GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("could not listen tpc port")
	}

	transactionService := service.NewTransactionService()

	transactionService.ProcessTransactionUseCase = server.ProcessTransactionUseCase

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterPaymentServiceServer(grpcServer, transactionService)

	grpcServer.Serve(lis)
}
