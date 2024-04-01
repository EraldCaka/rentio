package services

import (
	"context"
	"github.com/EraldCaka/rentio/internal/types"
)

type ClientService interface {
	CreateClient(ctx context.Context, clientReq *types.ClientRegisterRequest) (string, error)
	GetAllClients(ctx context.Context) ([]*types.Client, error)
	GetClientByID(ctx context.Context, clientID string) (*types.Client, error)
	UpdateClient(ctx context.Context, clientID string, clientReq *types.ClientRequest) error
	DeleteClient(ctx context.Context, clientID string) error
	LoginClient(ctx context.Context, clientReq *types.ClientRequest) error
}

type UserService interface {
	Create(ctx context.Context, userReq *types.UserCreateRequest) (string, error)
	GetAll(ctx context.Context) ([]*types.User, error)
	GetByID(ctx context.Context, userID string) (*types.User, error)
	Update(ctx context.Context, userID string, userReq *types.UserRequest) error
	Delete(ctx context.Context, userID string) error
	Login(ctx context.Context, userReq *types.UserRequest, token string) error
	Logout(ctx context.Context, token string) error
	GetActiveUserContract(ctx context.Context, userID string) (*types.Contract, error)
	GetAllUserContract(ctx context.Context, userID string) (*[]types.Contract, error)
}
