package services

import (
	"context"
	"github.com/EraldCaka/rentio/db"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
)

type ClientStore struct {
	db *db.Postgres
}

func NewClientService(db *db.Postgres) *ClientStore {
	return &ClientStore{db: db}
}

func (s *ClientStore) CreateClient(ctx context.Context, clientReq *types.ClientRegisterRequest) (string, error) {
	clientReq.Password, _ = util.HashPassword(clientReq.Password)
	return s.db.CreateClient(ctx, clientReq)
}

func (s *ClientStore) GetAllClients(ctx context.Context) ([]*types.Client, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetClient(ctx)
}

func (s *ClientStore) GetClientByID(ctx context.Context, clientID string) (*types.Client, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetClientByID(ctx, clientID)
}

func (s *ClientStore) UpdateClient(ctx context.Context, clientID string, clientReq *types.ClientRequest) error {
	// TODO : IMPLEMENT error handling LOGIC
	if err := s.db.UpdateClient(ctx, clientID, clientReq); err != nil {
		return err
	}
	return nil
}

func (s *ClientStore) DeleteClient(ctx context.Context, clientID string) error {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.DeleteClient(ctx, clientID)
}

func (s *ClientStore) LoginClient(ctx context.Context, userReq *types.ClientRequest) error {
	return s.db.LoginClient(ctx, userReq)
}
