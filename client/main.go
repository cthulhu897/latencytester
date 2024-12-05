package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "latencytester/gen/latencytester"

	"google.golang.org/grpc"
)

func main() {
	// Parámetros de línea de comandos
	serverAddress := flag.String("server", "localhost:50051", "Dirección del servidor (host:puerto)")
	numTests := flag.Int("tests", 15, "Número de pruebas a realizar")
	flag.Parse()

	// Conexión al servidor
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewLatencyServiceClient(conn)

	fmt.Printf("Iniciando pruebas de latencia hacia %s (%d pruebas)\n", *serverAddress, *numTests)

	// Latencia inicial (0 para la primera solicitud)
	var lastLatency float32 = 0

	for i := 0; i < *numTests; i++ {
		// Construir la solicitud con la última latencia medida
		req := &pb.LatencyRequest{
			Message:   "ping",
			LatencyMs: lastLatency, // 0 en la primera solicitud
		}

		start := time.Now()
		resp, err := client.MeasureLatency(context.Background(), req)
		durationMs := float32(time.Since(start).Seconds() * 1000) // Latencia actual en ms

		if err != nil {
			log.Printf("Error en la solicitud: %v\n", err)
			continue
		}

		// Imprimir la latencia medida y el historial recibido del servidor
		fmt.Printf("Prueba %d: Latencia medida = %.3f ms\n", i+1, durationMs)
		fmt.Printf("Latencias recientes recibidas del servidor: %v\n", resp.RecentLatencies)

		// Actualizar la última latencia medida
		lastLatency = durationMs

		// Esperar 500 ms antes de la siguiente prueba
		time.Sleep(500 * time.Millisecond)
	}
}
