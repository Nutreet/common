package clients

import (
	"context"

	proto "github.com/nutreet/common/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	user proto.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	user := proto.NewUserServiceClient(conn)

	return &UserClient{
		user,
	}, nil
}

func (c *UserClient) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	resp, err := c.user.Register(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *UserClient) GetAuthenticatedUser(ctx context.Context, req *proto.GetAuthenticatedUserRequest) (*proto.GetAuthenticatedUserResponse, error) {
	resp, err := c.user.GetAuthenticatedUser(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
