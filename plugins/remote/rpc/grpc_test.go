package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
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

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

type EchoServer struct {
}

func (e EchoServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	result := &HelloReply{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Message:       "",
	}
	return result, nil
}

func (e EchoServer) mustEmbedUnimplementedGreeterServer() {
	//todo
}
