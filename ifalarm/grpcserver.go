package ifalarm

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	*grpc.Server
	addr string
}

func (g *GrpcServer) Close() {
	g.Server.GracefulStop()
}

func NewGrpcServer(addr string) *GrpcServer {
	svc := grpc.NewServer()
	server := &GrpcServer{
		Server: svc,
		addr:   addr,
	}
	globalClosers = append(globalClosers, server)
	return server

}

func (g *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", g.addr)
	if err != nil {
		return err
	}
	go func() {
		if err := g.Server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return nil
}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
