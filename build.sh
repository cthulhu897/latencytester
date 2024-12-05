#!/bin/bash

# Crear directorio para archivos generados
mkdir -p gen/latencypb

# Generar archivos Protobuf en la ubicación correcta
protoc --go_out=gen --go-grpc_out=gen proto/latency.proto

# Limpiar dependencias y compilar
go mod tidy
go build -o ./client/ ./client 
go build -o ./server/ ./server 

