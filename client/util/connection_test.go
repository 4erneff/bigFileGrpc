package util

import (
	"context"
	"net"
	"testing"

	"io/ioutil"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"
)

// bufconnDialer for an in-memory gRPC server.
var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

// TestGetConn_Success tests the successful gRPC connection case.
func TestGetConn_Success(t *testing.T) {
	// Assume that server.crt is located in the same directory as the test file.
	certFile := "../server.crt"

	// Load the TLS credentials from the server.crt file.
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		t.Fatalf("Failed to load TLS credentials from server.crt: %v", err)
	}

	// Attempt to connect to the gRPC server at localhost:50051.
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}

	// Ensure that the connection is not nil.
	if conn == nil {
		t.Errorf("Expected non-nil connection, got nil")
	}

	// Close the connection when done.
	conn.Close()
}

// TestGetConn_CertificateError tests the case when the certificate file is missing.
func TestGetConn_CertificateError(t *testing.T) {
	// Simulate missing certificate file
	_, err := GetConn()

	// We expect an error because the certificate file doesn't exist
	if err == nil {
		t.Error("Expected error due to missing certificate, got nil")
	}
}

// TestGetConn_ConnectionError tests the case when the gRPC server is unavailable.
func TestGetConn_ConnectionError(t *testing.T) {
	// Mock the certificate file
	certFile := "server.crt"
	err := ioutil.WriteFile(certFile, []byte("mock certificate data"), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock certificate file: %v", err)
	}
	defer os.Remove(certFile) // Clean up the mock cert file after test

	// Temporarily change the server address to a non-existent one
	invalidServerAddr := "invalid"

	// Attempt to get connection
	conn, err := grpc.Dial(invalidServerAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err == nil && conn.GetState().String() != "CONNECTING" {
		t.Errorf("Expected error due to invalid server address, got nil")
	}

	if conn != nil {
		conn.Close()
	}
}
