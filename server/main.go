package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "latencytester/gen/latencytester"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLatencyServiceServer
}

func (s *server) MeasureLatency(ctx context.Context, req *pb.LatencyRequest) (*pb.LatencyResponse, error) {
	// Imprimir cuando hay un nuevo cliente
	log.Printf("Nuevo cliente procesando mensaje: %s", req.Message)
	// Responder con un mensaje rápido
	return &pb.LatencyResponse{Message: "pong" + req.Message}, nil
}

func main() {
	// Recibir el puerto desde argumentos de línea de comandos
	port := flag.String("port", "50051", "El puerto en el que el servidor escuchará")
	flag.Parse()

	address := fmt.Sprintf(":%s", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	log.Printf("Servidor gRPC escuchando en el puerto %s...", *port)

	grpcServer := grpc.NewServer()
	pb.RegisterLatencyServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}
