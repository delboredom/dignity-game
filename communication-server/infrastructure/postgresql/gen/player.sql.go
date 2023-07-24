// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: player.sql

package gen

import (
	"context"

	"github.com/google/uuid"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
    username, email, password, updated_at
) VALUES (
    $1, $2, $3, NOW()
)
RETURNING id, username, email, password, active, email_verified_at, created_at, updated_at
`

type CreatePlayerParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.queryRow(ctx, q.createPlayerStmt, createPlayer, arg.Username, arg.Email, arg.Password)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const playerByEmailOrUsername = `-- name: PlayerByEmailOrUsername :one
SELECT id, username, email, password, active, email_verified_at, created_at, updated_at FROM players
WHERE email = $1 OR
    username = $1 LIMIT 1
`

func (q *Queries) PlayerByEmailOrUsername(ctx context.Context, emailOrUsername string) (Player, error) {
	row := q.queryRow(ctx, q.playerByEmailOrUsernameStmt, playerByEmailOrUsername, emailOrUsername)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const playerByID = `-- name: PlayerByID :one
SELECT id, username, email, password, active, email_verified_at, created_at, updated_at FROM players
WHERE id = $1 LIMIT 1
`

func (q *Queries) PlayerByID(ctx context.Context, id uuid.UUID) (Player, error) {
	row := q.queryRow(ctx, q.playerByIDStmt, playerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyPlayerEmail = `-- name: VerifyPlayerEmail :exec
UPDATE players SET email_verified_at = NOW()
WHERE id = $1
`

func (q *Queries) VerifyPlayerEmail(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.verifyPlayerEmailStmt, verifyPlayerEmail, id)
	return err
}
