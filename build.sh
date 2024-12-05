#!/bin/bash

# Crear directorio para archivos generados
mkdir -p gen/latencypb

# Generar archivos Protobuf en la ubicación correcta
protoc --go_out=. --go-grpc_out=. proto/latency.proto

# Mover archivos generados al directorio adecuado
mv latencytester/gen/latencypb/* gen/latencypb/
rm -rf latencytester/gen

# Limpiar dependencias y compilar
go mod tidy
go build latencyclient.go
go build latencyserver.go

