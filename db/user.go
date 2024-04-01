package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
	"log"
	"time"
)

func (pg *Postgres) GetAllUserContract(ctx context.Context, userID string) (*[]types.Contract, error) {
	query := "SELECT * FROM public.contracts WHERE user_id = $1"

	rows, err := pg.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching contracts for user with ID %s: %v", userID, err)
	}
	defer rows.Close()

	var contracts []types.Contract
	for rows.Next() {
		var contract types.Contract
		err := rows.Scan(&contract.ID, &contract.RoomID, &contract.UserID, &contract.StartDate, &contract.EndDate, &contract.Rent, &contract.Status)
		if err != nil {
			return nil, fmt.Errorf("error scanning contract row: %v", err)
		}
		contracts = append(contracts, contract)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over contract rows: %v", err)
	}

	return &contracts, nil
}

func (pg *Postgres) GetActiveUserContract(ctx context.Context, userID string) (*types.Contract, error) {
	query := "SELECT * FROM public.contracts WHERE user_id = $1 AND status = $2"
	var contract types.Contract
	var activeStatus = 0
	row := pg.db.QueryRow(ctx, query, userID, activeStatus)
	err := row.Scan(&contract.ID, &contract.RoomID, &contract.UserID, &contract.StartDate, &contract.EndDate, &contract.Rent, &contract.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s has no active contract", userID)
		}
		return nil, err
	}
	return &contract, nil
}

func (pg *Postgres) CreateUser(ctx context.Context, u *types.UserCreateRequest) (string, error) {
	query := "INSERT INTO public.users (username, password, role) VALUES ($1, $2, $3) RETURNING id"
	var userID string
	err := pg.db.QueryRow(ctx, query, u.Username, u.Password, u.Role).Scan(&userID)
	if err != nil {
		log.Printf("Unable to insert user: %v\n", err)
		return "", err
	}
	return userID, nil
}

func (pg *Postgres) GetUserByID(ctx context.Context, userID string) (*types.User, error) {
	query := "SELECT * FROM public.users WHERE id = $1"
	row := pg.db.QueryRow(ctx, query, userID)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s not found", userID)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func (pg *Postgres) GetUserByName(ctx context.Context, username string) (*types.User, error) {
	query := "SELECT * FROM public.users WHERE username = $1"
	row := pg.db.QueryRow(ctx, query, username)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s not found", username)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func (pg *Postgres) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	query := "SELECT * FROM public.users"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}
	return users, nil
}

func (pg *Postgres) UpdateUser(ctx context.Context, userID string, u *types.UserRequest) error {
	query := "UPDATE public.users SET username=$1, password=$2 WHERE id=$3"
	_, err := pg.db.Exec(ctx, query, u.Username, u.Password, userID)
	if err != nil {
		log.Printf("Unable to update user: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) DeleteUser(ctx context.Context, userID string) error {
	_, err := pg.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user with ID %s not found", userID)
	}
	query := "DELETE FROM public.users WHERE id=$1"
	_, err = pg.db.Exec(ctx, query, userID)
	if err != nil {
		log.Printf("Unable to delete user: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) Login(ctx context.Context, u *types.UserRequest, token string) error {
	query := "SELECT id, username, password, role FROM public.users WHERE username = $1"
	row := pg.db.QueryRow(ctx, query, u.Username)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user with username %s not found", u.Username)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return err
	}

	if user.ID == "" {
		return fmt.Errorf("user with username %s not found", u.Username)
	}

	if err := util.CheckPassword(u.Password, user.Password); err != nil {
		return fmt.Errorf("invalid password for user %s", u.Username)
	}

	activeUserQuery := "INSERT INTO public.active_users (role, jwt_token, username, expire_time) VALUES ($1, $2, $3, $4)"
	expirationTime := time.Now().Add(time.Hour)
	if _, err := pg.db.Exec(ctx, activeUserQuery, user.Role, token, user.Username, expirationTime); err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) Logout(ctx context.Context, token string) error {
	query := "DELETE FROM public.active_users WHERE jwt_token=$1"
	_, err := pg.db.Exec(ctx, query, token)
	if err != nil {
		log.Printf("Unable to delete user: %v\n", err)
		return err
	}
	return nil
}
