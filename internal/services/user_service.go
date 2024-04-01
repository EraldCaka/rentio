package services

import (
	"context"
	"github.com/EraldCaka/rentio/db"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
)

type UserStore struct {
	db *db.Postgres
}

func NewUserService(db *db.Postgres) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(ctx context.Context, userReq *types.UserCreateRequest) (string, error) {
	userReq.Password, _ = util.HashPassword(userReq.Password)
	return s.db.CreateUser(ctx, userReq)
}

func (s *UserStore) GetAll(ctx context.Context) ([]*types.User, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetUsers(ctx)
}

func (s *UserStore) GetByID(ctx context.Context, userID string) (*types.User, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetUserByID(ctx, userID)
}

func (s *UserStore) Update(ctx context.Context, userID string, userReq *types.UserRequest) error {
	// TODO : IMPLEMENT error handling LOGIC
	if err := s.db.UpdateUser(ctx, userID, userReq); err != nil {
		return err
	}
	return nil
}

func (s *UserStore) Delete(ctx context.Context, userID string) error {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.DeleteUser(ctx, userID)
}

func (s *UserStore) Login(ctx context.Context, userReq *types.UserRequest, token string) error {
	return s.db.Login(ctx, userReq, token)
}

func (s *UserStore) Logout(ctx context.Context, token string) error {
	return s.db.Logout(ctx, token)
}

func (s *UserStore) GetActiveUserContract(ctx context.Context, userID string) (*types.Contract, error) {
	return s.db.GetActiveUserContract(ctx, userID)
}

func (s *UserStore) GetAllUserContract(ctx context.Context, userID string) (*[]types.Contract, error) {
	return s.db.GetAllUserContract(ctx, userID)
}
