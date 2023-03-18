package gateway

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestNewServeMux(t *testing.T) {
	ctx := context.Background()
	opts := Options{
		Addr: ":8080",
		RPCServer: Endpoint{
			Network: "tcp",
			Addr:    "localhost:9090",
		},
		OpenAPIDir: "examples/hello",
		Handlers: []ServiceHandler{
			RegisterEchoService,
		},
	}

	Run(ctx, opts)
}

func RegisterEchoService(ctx context.Context, mux *ServeMux, conn *grpc.ClientConn) error {

	return nil
}
