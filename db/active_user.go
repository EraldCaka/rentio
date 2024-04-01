package db

import (
	"context"
	"github.com/EraldCaka/rentio/internal/types"
	"log"
)

func (pg *Postgres) NewLoggedInUser(ctx context.Context, user *types.ActiveUsersRequest) error {
	query := "INSERT INTO public.active_users (role, jwt_token,username,expire_time) VALUES ($1, $2, $3, $4)"
	err := pg.db.QueryRow(ctx, query, user.Role, user.JwtToken, user.Username, user.ExpireTime).Scan()
	if err != nil {
		log.Printf("Unable to insert active_user: %v\n", err)
		return err
	}
	return nil
}
