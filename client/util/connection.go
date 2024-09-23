package util

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	serverAddr = "localhost:50051"
)

func GetConn() (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
