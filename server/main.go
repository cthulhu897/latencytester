package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "latencytester/gen/latencytester"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLatencyServiceServer
}

func (s *server) MeasureLatency(ctx context.Context, req *pb.LatencyRequest) (*pb.LatencyResponse, error) {
	// Respuesta estática, sin logs ni concatenación
	return &pb.LatencyResponse{Message: "pong"}, nil
}

func main() {
	port := flag.String("port", "50051", "El puerto en el que el servidor escuchará")
	flag.Parse()

	address := fmt.Sprintf(":%s", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	// Ajustar buffers opcionalmente (si deseas experimentar):
	grpcServer := grpc.NewServer(
		grpc.WriteBufferSize(64*1024),
		grpc.ReadBufferSize(64*1024),
	)

	pb.RegisterLatencyServiceServer(grpcServer, &server{})

	// Servidor sin impresiones ni logs
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
