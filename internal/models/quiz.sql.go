// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quiz.sql

package models

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getQuestionsByQuizId = `-- name: GetQuestionsByQuizId :many
SELECT id, quiz_id, question, options, correct_option, updated_at, created_at, updated_by FROM "quiz_question" WHERE quiz_id = $1
`

func (q *Queries) GetQuestionsByQuizId(ctx context.Context, quizID uuid.UUID) ([]QuizQuestion, error) {
	rows, err := q.db.QueryContext(ctx, getQuestionsByQuizId, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuizQuestion
	for rows.Next() {
		var i QuizQuestion
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Question,
			pq.Array(&i.Options),
			&i.CorrectOption,
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

const saveQuiz = `-- name: SaveQuiz :one
INSERT INTO
    "quiz" (
        name,
        thumbnail_url,
        play_count,
        updated_by
    )
VALUES ($1, $2, $3, $4) RETURNING id, name, thumbnail_url, play_count, updated_at, created_at, updated_by
`

type SaveQuizParams struct {
	Name         sql.NullString
	ThumbnailUrl sql.NullString
	PlayCount    sql.NullInt32
	UpdatedBy    uuid.NullUUID
}

func (q *Queries) SaveQuiz(ctx context.Context, arg SaveQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, saveQuiz,
		arg.Name,
		arg.ThumbnailUrl,
		arg.PlayCount,
		arg.UpdatedBy,
	)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ThumbnailUrl,
		&i.PlayCount,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const saveQuizQuestion = `-- name: SaveQuizQuestion :one
INSERT INTO
    "quiz_question" (
        quiz_id,
        question,
        options,
        correct_option,
        updated_by
    )
VALUES ($1, $2, $3, $4, $5) RETURNING id, quiz_id, question, options, correct_option, updated_at, created_at, updated_by
`

type SaveQuizQuestionParams struct {
	QuizID        uuid.UUID
	Question      sql.NullString
	Options       []string
	CorrectOption sql.NullInt32
	UpdatedBy     uuid.NullUUID
}

func (q *Queries) SaveQuizQuestion(ctx context.Context, arg SaveQuizQuestionParams) (QuizQuestion, error) {
	row := q.db.QueryRowContext(ctx, saveQuizQuestion,
		arg.QuizID,
		arg.Question,
		pq.Array(arg.Options),
		arg.CorrectOption,
		arg.UpdatedBy,
	)
	var i QuizQuestion
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Question,
		pq.Array(&i.Options),
		&i.CorrectOption,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return i, err
}
