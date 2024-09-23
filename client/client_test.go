package main

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"testing"

	"io/ioutil"
	"os"

	pb "github.com/4erneff/alcatraz/pb/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MockFileServiceClient struct {
	mock.Mock
}

func (m *MockFileServiceClient) GetFileStream(ctx context.Context, in *pb.FileRequest, opts ...grpc.CallOption) (pb.FileService_GetFileStreamClient, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(pb.FileService_GetFileStreamClient), args.Error(1)
}

func (m *MockFileServiceClient) GetFileMetadata(ctx context.Context, in *pb.FileMetadataRequest, opts ...grpc.CallOption) (*pb.FileMetadataResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.FileMetadataResponse), args.Error(1)
}

type MockFileService_GetFileStreamClient struct {
	mock.Mock
}

func (m *MockFileService_GetFileStreamClient) Recv() (*pb.FileChunk, error) {
	args := m.Called()
	return args.Get(0).(*pb.FileChunk), args.Error(1)
}

func (m *MockFileService_GetFileStreamClient) CloseSend() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFileService_GetFileStreamClient) Context() context.Context {
	args := m.Called()
	return args.Get(0).(context.Context)
}

func (m *MockFileService_GetFileStreamClient) Header() (metadata.MD, error) {
	args := m.Called()
	return nil, args.Error(0)
}

func (m *MockFileService_GetFileStreamClient) Trailer() metadata.MD {
	args := m.Called()
	return args.Get(0).(map[string][]string)
}

func (m *MockFileService_GetFileStreamClient) SendMsg(msg interface{}) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockFileService_GetFileStreamClient) RecvMsg(msg interface{}) error {
	args := m.Called(msg)
	return args.Error(0)
}

func TestDownloadFile_Success(t *testing.T) {
	// Create a temporary file for download output.
	tmpFile, err := ioutil.TempFile("", "downloaded_file_parallel.mov")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Cleanup the temporary file after the test

	mockClient := new(MockFileServiceClient)
	mockStream := new(MockFileService_GetFileStreamClient)
	for i := 0; i < 10; i++ {
		data := make([]byte, fileChunkSize)
		chunk := &pb.FileChunk{
			SequenceNumber: int32(i),
			ChunkData:      data,
			Checksum:       fmt.Sprintf("%x", sha256.Sum256(data)), // Assume checksum verification passes
		}
		mockStream.On("Recv").Return(chunk, nil).Once()
	}
	mockStream.On("Recv").Return(&pb.FileChunk{}, io.EOF)

	// Mock the file stream
	mockClient.On("GetFileStream", mock.Anything, mock.Anything).Return(mockStream, nil)

	startChunk := 0
	lc, err := downloadFile(mockClient, startChunk, 10, int64(10*fileChunkSize))
	if err != nil && lc != 10 {
		t.Fatalf("downloadFile failed: %v", err)
	}

	mockClient.AssertExpectations(t)
	mockStream.AssertExpectations(t)
}

func TestDownloadFile_ConnectionDrop(t *testing.T) {
	// Create a temporary file for download output.
	tmpFile, err := ioutil.TempFile("", "downloaded_file_parallel.mov")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Cleanup the temporary file after the test

	// Mock the gRPC client and stream
	mockClient := new(MockFileServiceClient)
	mockStream := new(MockFileService_GetFileStreamClient)

	for i := 0; i < 6; i++ {
		data := make([]byte, fileChunkSize)
		chunk := &pb.FileChunk{
			SequenceNumber: int32(i),
			ChunkData:      data,
			Checksum:       fmt.Sprintf("%x", sha256.Sum256(data)), // Assume checksum verification passes
		}
		mockStream.On("Recv").Return(chunk, nil).Once()
	}

	// Simulate a connection drop (error) after receiving 5 chunks
	mockStream.On("Recv").Return(&pb.FileChunk{}, errors.New("connection dropped")).Once()
	mockClient.On("GetFileStream", mock.Anything, mock.Anything).Return(mockStream, nil)

	startChunk := 0
	_, err = downloadFile(mockClient, startChunk, 10, int64(10*fileChunkSize))

	if err == nil {
		t.Fatalf("Expected connection drop error, but got nil")
	}

	mockClient.AssertExpectations(t)
	mockStream.AssertExpectations(t)
}
