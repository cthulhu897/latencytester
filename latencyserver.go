package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "latencytester/gen/latencypb"
)

type server struct {
	pb.UnimplementedLatencyServiceServer
}

func (s *server) MeasureLatency(ctx context.Context, req *pb.LatencyRequest) (*pb.LatencyResponse, error) {
	// Responder con un mensaje rápido
	return &pb.LatencyResponse{Message: "Respuesta recibida: " + req.Message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLatencyServiceServer(grpcServer, &server{})

	fmt.Println("Servidor gRPC escuchando en el puerto 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}
