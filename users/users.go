package users

import (
	"context"
	"sync"

	usersv1 "github.com/johanbrandhorst/grpc-web-go-react-example/gen/users/v1"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usersv1.UnimplementedUserServiceServer

	mu    sync.Mutex
	users map[string]*usersv1.User
}

func (u *UserService) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	if req.GetUserId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing user id")
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[req.GetUserId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &usersv1.GetUserResponse{
		User: user,
	}, nil
}

func (u *UserService) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing name")
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	user := &usersv1.User{
		Id:   ksuid.New().String(),
		Name: req.GetName(),
	}
	if u.users == nil {
		u.users = map[string]*usersv1.User{}
	}
	u.users[user.Id] = user
	return &usersv1.CreateUserResponse{
		User: user,
	}, nil
}
