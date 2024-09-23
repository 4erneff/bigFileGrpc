# Large File Transfer via gRPC

This project demonstrates a simple gRPC service that:
- **Server**: Generates a large file (1GB) on startup and serves it through a gRPC endpoint.
- **Client**: Shows an example of how to consume the server's file service to download the file in chunks.

### Proto Definitions
All gRPC interfaces are defined in `proto/server.proto`. These definitions specify the gRPC methods available for file metadata and file streaming.

### Server
The server generates a large file (1GB) when started and exposes two gRPC endpoints:
- `GetFileMetadata`: Returns metadata about the file, such as its total size and the number of chunks.
- `GetFileStream`: Streams the file in chunks to the client.

### Client
The client demonstrates how to consume the gRPC service provided by the server. It fetches the file metadata and downloads the file in chunks, validating the data using checksums.

## How to Run

### Prerequisites
Ensure you have the following installed:
- [Go](https://golang.org/dl/)
- gRPC Go libraries
- Protocol Buffers compiler (for generating Go code from `.proto` files)

### 1. Clone the Repository
```bash```
git clone <repository-url>
cd <repository-name>

### 2. Run the server
```bash```
cd server/
go run main.go

### 2. Run the client
```bash```
cd client/
go run main.go

## gRPC Interface

The gRPC service provides the following methods:

### GetFileMetadata
- **Request**: Empty
- **Response**:
  - `TotalSize`: Size of the file in bytes.
  - `TotalChunks`: Number of chunks the file is divided into.

### GetFileStream
- **Request**:
  - `StartChunk`: The chunk number from where the download should start.
- **Response**:
  - `SequenceNumber`: The current chunk number.
  - `ChunkData`: The data of the chunk.
  - `Checksum`: SHA-256 checksum of the chunk data.
  - `TotalSize`: Total size of the file.
  - `TotalChunks`: Total number of chunks in the file.





