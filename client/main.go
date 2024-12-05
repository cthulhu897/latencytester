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

func main() {
	// Parámetros de línea de comandos
	serverAddress := flag.String("server", "localhost:50051", "Dirección del servidor (host:puerto)")
	numTests := flag.Int("tests", 21, "Número de pruebas a realizar")
	flag.Parse()

	// Conexión al servidor
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewLatencyServiceClient(conn)

	fmt.Printf("Iniciando pruebas de latencia hacia %s (%d pruebas)\n", *serverAddress, *numTests)

	var lastLatency float32 = 0
	var latencies []float32

	for i := 0; i < *numTests; i++ {
		req := &pb.LatencyRequest{
			Message:   "ping",
			LatencyMs: lastLatency,
		}

		start := time.Now()
		_, err := client.MeasureLatency(context.Background(), req)
		durationMs := float32(time.Since(start).Seconds() * 1000) // Latencia actual en ms

		if err != nil {
			log.Printf("Error en la solicitud: %v\n", err)
			continue
		}

		// Imprimir la latencia medida en esta prueba
		fmt.Printf("Prueba %d: Latencia = %.3f ms\n", i+1, durationMs)

		// Guardar la latencia en la lista para análisis posterior
		latencies = append(latencies, durationMs)

		// Actualizar la última latencia medida
		lastLatency = durationMs

		// Esperar 500 ms antes de la siguiente prueba
		time.Sleep(600 * time.Millisecond)
	}

	// Análisis final de las latencias medidas
	if len(latencies) > 0 {
		min, max, avg := analyzeLatencies(latencies)
		fmt.Println("\n--- Resumen de Latencias ---")
		fmt.Printf("Mejor latencia (mínima): %.3f ms\n", min)
		fmt.Printf("Peor latencia (máxima):  %.3f ms\n", max)
		fmt.Printf("Latencia promedio:       %.3f ms\n", avg)
	} else {
		fmt.Println("No se pudieron medir latencias correctamente.")
	}
}

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

	arr = arr[5 : len(arr)-5]
	min, max := arr[0], arr[len(arr)-1]

	var sum float32
	for _, v := range arr {
		sum += v
	}
	avg := sum / float32(len(arr))

	return min, max, avg
}
