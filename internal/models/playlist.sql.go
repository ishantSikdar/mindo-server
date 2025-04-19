// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: playlist.sql

package models

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPlaylist = `-- name: CreatePlaylist :one
INSERT INTO
    playlist (
        name,
        description,
        thumbnail_url,
        updated_by
    )
VALUES (
        $1, -- Name
        $2, -- Description
        $3, -- Thumbnail URL
        $4 -- Updated By
    ) RETURNING id, name, description, thumbnail_url, updated_at, created_at, updated_by
`

type CreatePlaylistParams struct {
	Name         sql.NullString
	Description  sql.NullString
	ThumbnailUrl sql.NullString
	UpdatedBy    uuid.NullUUID
}

// Create a new playlist
func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, createPlaylist,
		arg.Name,
		arg.Description,
		arg.ThumbnailUrl,
		arg.UpdatedBy,
	)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const getAllPlaylists = `-- name: GetAllPlaylists :many
SELECT id, name, description, thumbnail_url, updated_at, created_at, updated_by FROM playlist
`

// Fetch all playlists
func (q *Queries) GetAllPlaylists(ctx context.Context) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, getAllPlaylists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playlist
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ThumbnailUrl,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlaylistByID = `-- name: GetPlaylistByID :one
SELECT
    id,
    name,
    description,
    thumbnail_url,
    updated_by,
    created_at,
    updated_at
FROM playlist
WHERE
    id = $1
`

type GetPlaylistByIDRow struct {
	ID           uuid.UUID
	Name         sql.NullString
	Description  sql.NullString
	ThumbnailUrl sql.NullString
	UpdatedBy    uuid.NullUUID
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

// Get a playlist by ID
func (q *Queries) GetPlaylistByID(ctx context.Context, id uuid.UUID) (GetPlaylistByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getPlaylistByID, id)
	var i GetPlaylistByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
