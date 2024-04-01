package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
	"log"
)

func (pg *Postgres) CreateClient(ctx context.Context, client *types.ClientRegisterRequest) (string, error) {
	query := "INSERT INTO public.clients (username, password) VALUES ($1, $2) RETURNING id"
	var clientID string
	err := pg.db.QueryRow(ctx, query, client.Username, client.Password).Scan(&clientID)
	if err != nil {
		log.Printf("Unable to insert client: %v\n", err)
		return "", err
	}
	return clientID, nil
}

func (pg *Postgres) GetClientByID(ctx context.Context, clientID string) (*types.Client, error) {
	query := "SELECT * FROM public.clients WHERE id = $1"
	row := pg.db.QueryRow(ctx, query, clientID)

	var client types.Client
	err := row.Scan(&client.ID, &client.Username, &client.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("client with ID %s not found", clientID)
		}
		log.Printf("Error scanning client data: %v\n", err)
		return nil, err
	}
	return &client, nil
}

func (pg *Postgres) GetClient(ctx context.Context) ([]*types.Client, error) {
	var clients []*types.Client

	query := "SELECT * FROM public.clients"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying clients: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client types.Client
		err := rows.Scan(&client.ID, &client.Username, &client.Password)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		clients = append(clients, &client)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over client rows: %v\n", err)
		return nil, err
	}
	return clients, nil
}

func (pg *Postgres) UpdateClient(ctx context.Context, clientID string, u *types.ClientRequest) error {
	query := "UPDATE public.clients SET username=$1, password=$2 WHERE id=$3"
	_, err := pg.db.Exec(ctx, query, u.Username, u.Password, clientID)
	if err != nil {
		log.Printf("unable to update client: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) DeleteClient(ctx context.Context, clientID string) error {
	_, err := pg.GetClientByID(ctx, clientID)
	if err != nil {
		return fmt.Errorf("client with ID %s not found", clientID)
	}
	query := "DELETE FROM public.clients WHERE id=$1"
	_, err = pg.db.Exec(ctx, query, clientID)
	if err != nil {
		log.Printf("unable to delete client: %v\n", err)
		return err
	}
	return nil
}
func (pg *Postgres) LoginClient(ctx context.Context, c *types.ClientRequest) error {
	query := "SELECT * FROM public.clients WHERE username = $1"
	row := pg.db.QueryRow(ctx, query, c.Username)
	var client types.Client
	err := row.Scan(&client.ID, &client.Username, &client.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("client with username %s not found", c.Username)
		}
		log.Printf("Error scanning client data: %v\n", err)
		return err
	}
	if client.ID == "" {
		return fmt.Errorf("wrong credentials for client %s", c.Username)
	}
	if err := util.CheckPassword(c.Password, client.Password); err != nil {
		return fmt.Errorf("invalid password for client %s", c.Username)
	}
	return nil
}
