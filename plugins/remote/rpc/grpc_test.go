package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v\n", network, address, err)
		}
	}()

	s := grpc.NewServer()
	RegisterGreeterServer(s, new(EchoServer))
	s.GetServiceInfo()

	return nil
}

type EchoServer struct {
}

func (e EchoServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	//TODO implement me
	panic("implement me")
}

func (e EchoServer) mustEmbedUnimplementedGreeterServer() {
	//TODO implement me
	panic("implement me")
}
