package main

//go:generate mkdir -p pb
//go:generate protoc --go-grpc_out=./pb --go-grpc_opt=paths=source_relative esquema.proto
