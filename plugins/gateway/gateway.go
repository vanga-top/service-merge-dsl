package gateway

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Endpoint struct {
	Network, Addr string
}

type Options struct {
	Addr       string
	RPCServer  Endpoint
	OpenAPIDir string
	Mux        []ServeMuxOption
	Handlers   []ServiceHandler // register http request to grpc
}

func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//todo 这里要考虑有多个服务端集群的情况
	conn, err := dial(ctx, opts.RPCServer.Network, opts.RPCServer.Addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close a client Connection to gRpc Server:%v\n", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/openapi/", openAPIServer(opts.OpenAPIDir))

	return nil
}

func dial(ctx context.Context, network string, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	}
	return nil, fmt.Errorf("unsupported network type %q", network)
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

type ServiceHandler func(context.Context, *ServeMux, *grpc.ClientConn) error

func newGateway(ctx context.Context, conn *grpc.ClientConn, opt Options) (http.Handler, error) {
	mux := NewServeMux(opt.Mux...)
	for _, f := range opt.Handlers {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
