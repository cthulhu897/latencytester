package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"
	"time"

	pb "latencytester/gen/latencytester"

	"google.golang.org/grpc"
)

// analyzeLatencies:
// 1. Crea una copia del slice.
// 2. Ordena las latencias de menor a mayor.
// 3. Elimina las 5 latencias más bajas y las 5 más altas.
// 4. Retorna el mínimo, máximo y la media del subset restante.
func analyzeLatencies(latencies []float32) (float32, float32, float32) {
	if len(latencies) <= 10 {
		return 0, 0, 0
	}

	arr := append([]float32(nil), latencies...)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	arr = arr[3 : len(arr)-3]
	min, max := arr[0], arr[len(arr)-1]

	var sum float32
	for _, v := range arr {
		sum += v
	}
	avg := sum / float32(len(arr))

	return min, max, avg
}

func main() {
	serverAddress := flag.String("server", "localhost:50051", "Dirección del servidor (host:puerto)")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewLatencyServiceClient(conn)
	req := &pb.LatencyRequest{Message: "Prueba"}

	fmt.Println("Presiona Enter para detener la prueba.")
	stop := make(chan struct{})

	// Goroutine para detener la ejecución con Enter
	go func() {
		fmt.Scanln()
		close(stop)
	}()

	var latencies []float32

	for {
		select {
		case <-stop:
			fmt.Println("Finalizando pruebas...")
			// Analizar las latencias
			min, max, avg := analyzeLatencies(latencies)
			if len(latencies) > 10 {
				fmt.Println("Análisis de latencias (descartando bajas/altas):")
				fmt.Printf("Mínimo: %.3f ms\n", min)
				fmt.Printf("Máximo: %.3f ms\n", max)
				fmt.Printf("Media: %.3f ms\n", avg)
			} else {
				fmt.Println("No se recopilaron suficientes datos para el análisis filtrado.")
			}
			return
		default:
			// Medir latencia
			start := time.Now()
			_, err := client.MeasureLatency(context.Background(), req)
			if err != nil {
				// Si hay error, solo lo registramos y continuamos
				log.Printf("Error en la solicitud: %v", err)
				continue
			}
			duration := float32(time.Since(start).Seconds() * 1000)
			fmt.Printf("Latencia: %.3f ms\n", duration)

			// Guardar la latencia
			latencies = append(latencies, duration)
		}
	}
}
