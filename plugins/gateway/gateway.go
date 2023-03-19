package gateway

import (
	"context"
	"dsl/plugins"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
)

type Gateway struct {
	Opts          Options
	Ctx           context.Context
	GatewayStatus plugins.Status
}

func (g *Gateway) ID() string {
	return "gateway"
}

func (g *Gateway) Name() string {
	return "gateway-dsl-service"
}

func (g *Gateway) Init() error {
	g.GatewayStatus = plugins.INITED
	return g.run(g.Ctx, g.Opts)
}

func (g *Gateway) Status() plugins.Status {
	return g.GatewayStatus
}

func (g *Gateway) Destroy() error {
	g.Ctx.Done()
	fmt.Printf("destroy plugin Name:%s\n", g.Name())
	return nil
}

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

func (g *Gateway) run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)

	//todo 这里要考虑有多个服务端集群的情况
	conn, err := dial(ctx, opts.RPCServer.Network, opts.RPCServer.Addr)
	if err != nil {
		cancel()
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

	gw, err := newGateway(ctx, conn, opts)
	if err != nil {
		cancel()
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    opts.Addr,
		Handler: allowCORS(mux),
	}

	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Printf("Failed to Shutdown http Server:%v\n", err)
		}
	}()

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			cancel()
		}
	}()

	return nil
}

func dial(ctx context.Context, network string, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	case "udp":
		return dialUDP(ctx, addr)
	}
	return nil, fmt.Errorf("unsupported network type %q", network)
}

func dialUDP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr)
}

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

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}
