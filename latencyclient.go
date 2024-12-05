package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "latencytester/gen/latencypb"
)

func main() {
	conn, err := grpc.Dial("34.51.5.52:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewLatencyServiceClient(conn)

	fmt.Println("Presiona Enter para detener la prueba.")
	stop := make(chan struct{})

	// Goroutine para detectar entrada de usuario
	go func() {
		fmt.Scanln()
		close(stop)
	}()

	// Ciclo para medir latencia
	for {
		select {
		case <-stop:
			fmt.Println("Finalizando pruebas...")
			return
		default:
			start := time.Now()
			_, err := client.MeasureLatency(context.Background(), &pb.LatencyRequest{Message: "Prueba"})
			if err != nil {
				log.Printf("Error en la solicitud: %v", err)
				continue
			}
			duration := time.Since(start)
			fmt.Printf("Latencia: %.3f ms\n", duration.Seconds()*1000)
		}
	}
}
