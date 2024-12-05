package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	pb "latencytester/gen/latencytester"

	"google.golang.org/grpc"
)

// Servidor con referencia al canal de logs para enviar mensajes
type server struct {
	pb.UnimplementedLatencyServiceServer
	logChan chan string
}

func (s *server) MeasureLatency(ctx context.Context, req *pb.LatencyRequest) (*pb.LatencyResponse, error) {
	// Enviar el mensaje al canal de logs de forma asíncrona
	msg := fmt.Sprintf("Nuevo cliente procesando mensaje: %s", req.Message)
	select {
	case s.logChan <- msg:
		// Log enviado al canal, no bloqueamos
	default:
		// Si el canal está lleno, podemos descartar el log o implementar una lógica
		// alternativa (ej: incrementar un contador de logs perdidos).
	}

	// Responder sin retener la ejecución
	return &pb.LatencyResponse{Message: "pong"}, nil
}

func main() {
	port := flag.String("port", "50051", "El puerto en el que el servidor escuchará")
	flag.Parse()

	address := fmt.Sprintf(":%s", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Sprintf("Error al iniciar el servidor: %v", err))
	}

	// Crear el canal de logs con buffer para evitar bloqueos
	logChan := make(chan string, 1000)

	// Lanzar una goroutine para procesar logs asíncronamente
	go func() {
		for msg := range logChan {
			// Aquí puedes elegir cómo loguear: a stdout, a un archivo, etc.
			// Esto se ejecuta en un hilo separado, sin afectar la latencia del handler
			fmt.Println(time.Now().Format(time.RFC3339), msg)
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterLatencyServiceServer(grpcServer, &server{logChan: logChan})

	fmt.Printf("Servidor gRPC escuchando en el puerto %s...\n", *port)
	if err := grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("Error al ejecutar el servidor: %v", err))
	}
}
