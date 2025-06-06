// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_token.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getUserTokenByRefreshToken = `-- name: GetUserTokenByRefreshToken :one
SELECT id, user_id, role, refresh_token, expires_at, updated_at, created_at, updated_by FROM user_token WHERE refresh_token = $1
`

func (q *Queries) GetUserTokenByRefreshToken(ctx context.Context, refreshToken uuid.UUID) (UserToken, error) {
	row := q.db.QueryRowContext(ctx, getUserTokenByRefreshToken, refreshToken)
	var i UserToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const upsertUserToken = `-- name: UpsertUserToken :one
INSERT INTO
    user_token (
        user_id,
        refresh_token,
        role,
        expires_at,
        updated_by
    )
VALUES (
    $1, -- User Id
    $2, -- Refresh Token
    $3, -- Role
    $4, -- Expires At
    $5  -- Updated By
)
ON CONFLICT (user_id)  -- Specify the unique constraint (e.g., user_id)
DO UPDATE SET
    refresh_token = EXCLUDED.refresh_token,
    expires_at = EXCLUDED.expires_at,
    updated_by = EXCLUDED.updated_by
RETURNING id, user_id, role, refresh_token, expires_at, updated_at, created_at, updated_by
`

type UpsertUserTokenParams struct {
	UserID       uuid.UUID
	RefreshToken uuid.UUID
	Role         UserType
	ExpiresAt    time.Time
	UpdatedBy    uuid.NullUUID
}

func (q *Queries) UpsertUserToken(ctx context.Context, arg UpsertUserTokenParams) (UserToken, error) {
	row := q.db.QueryRowContext(ctx, upsertUserToken,
		arg.UserID,
		arg.RefreshToken,
		arg.Role,
		arg.ExpiresAt,
		arg.UpdatedBy,
	)
	var i UserToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Role,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return i, err
}
