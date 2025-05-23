package ifapm

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	*grpc.ClientConn
}

func NewGrpcClient(addr string) (*GrpcClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithUnaryInterceptor(unaryInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{conn}, nil
}

func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
