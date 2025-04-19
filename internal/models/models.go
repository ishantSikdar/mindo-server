// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type UserType string

const (
	UserTypeAppUser UserType = "app_user"
	UserTypeAdmin   UserType = "admin"
)

func (e *UserType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserType(s)
	case string:
		*e = UserType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserType: %T", src)
	}
	return nil
}

type NullUserType struct {
	UserType UserType
	Valid    bool // Valid is true if UserType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserType) Scan(value interface{}) error {
	if value == nil {
		ns.UserType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserType), nil
}

type AdminUser struct {
	UserID       uuid.UUID
	Name         sql.NullString
	Email        sql.NullString
	PasswordHash sql.NullString
	LastLoginAt  sql.NullTime
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedBy    uuid.NullUUID
}

type AppUser struct {
	UserID            uuid.UUID
	Username          sql.NullString
	ProfilePictureUrl sql.NullString
	Bio               sql.NullString
	Name              sql.NullString
	Mobile            sql.NullString
	Email             sql.NullString
	PasswordHash      sql.NullString
	LastLoginAt       sql.NullTime
	UpdatedAt         sql.NullTime
	CreatedAt         sql.NullTime
	UpdatedBy         uuid.NullUUID
}

type AppUserInterest struct {
	ID         uuid.UUID
	AppUserID  uuid.UUID
	InterestID uuid.NullUUID
	Name       sql.NullString
	UpdatedAt  sql.NullTime
	CreatedAt  sql.NullTime
	UpdatedBy  uuid.NullUUID
}

type Community struct {
	ID           uuid.UUID
	Title        sql.NullString
	About        sql.NullString
	ThumbnailUrl sql.NullString
	LogoUrl      sql.NullString
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedBy    uuid.NullUUID
}

type Interest struct {
	ID        uuid.UUID
	Name      sql.NullString
	UpdatedAt sql.NullTime
	CreatedAt sql.NullTime
	UpdatedBy uuid.NullUUID
}

type JoinedCommunity struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	CommunityID uuid.UUID
	UpdatedAt   sql.NullTime
	CreatedAt   sql.NullTime
	UpdatedBy   uuid.NullUUID
}

type Message struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	CommunityID uuid.UUID
	Content     sql.NullString
	UpdatedAt   sql.NullTime
	CreatedAt   sql.NullTime
	UpdatedBy   uuid.NullUUID
}

type Playlist struct {
	ID           uuid.UUID
	Name         sql.NullString
	Description  sql.NullString
	ThumbnailUrl sql.NullString
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedBy    uuid.NullUUID
}

type Quiz struct {
	ID           uuid.UUID
	Name         sql.NullString
	ThumbnailUrl sql.NullString
	PlayCount    sql.NullInt32
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedBy    uuid.NullUUID
}

type QuizQuestion struct {
	ID            uuid.UUID
	QuizID        uuid.UUID
	Question      sql.NullString
	Options       []string
	CorrectOption sql.NullInt32
	UpdatedAt     sql.NullTime
	CreatedAt     sql.NullTime
	UpdatedBy     uuid.NullUUID
}

type QuizResult struct {
	ID        uuid.UUID
	QuizID    uuid.UUID
	UserID    uuid.UUID
	Score     sql.NullString
	UpdatedAt sql.NullTime
	CreatedAt sql.NullTime
	UpdatedBy uuid.NullUUID
}

type QuizResultQuestion struct {
	ID              uuid.UUID
	QuizResultID    uuid.UUID
	QuizQuestionID  uuid.UUID
	AttemptedOption sql.NullInt32
	UpdatedAt       sql.NullTime
	CreatedAt       sql.NullTime
	UpdatedBy       uuid.NullUUID
}

type SavedStudyMaterial struct {
	ID              uuid.UUID
	StudyMaterialID uuid.UUID
	UserID          uuid.UUID
	UpdatedAt       sql.NullTime
	CreatedAt       sql.NullTime
	UpdatedBy       uuid.NullUUID
}

type StudyMaterial struct {
	ID        uuid.UUID
	TopicID   uuid.UUID
	Title     sql.NullString
	Content   sql.NullString
	UpdatedAt sql.NullTime
	CreatedAt sql.NullTime
	UpdatedBy uuid.NullUUID
}

type Topic struct {
	ID         uuid.UUID
	Number     sql.NullInt32
	Name       sql.NullString
	PlaylistID uuid.UUID
	UpdatedAt  sql.NullTime
	CreatedAt  sql.NullTime
	UpdatedBy  uuid.NullUUID
}

type User struct {
	ID        uuid.UUID
	UserType  NullUserType
	UpdatedAt sql.NullTime
	CreatedAt sql.NullTime
	UpdatedBy uuid.NullUUID
}

type WatchedPlaylist struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	PlaylistID uuid.UUID
	Progress   sql.NullInt32
	UpdatedAt  sql.NullTime
	CreatedAt  sql.NullTime
	UpdatedBy  uuid.NullUUID
}

type WatchedVideo struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	YoutubeVideoID uuid.UUID
	UpdatedAt      sql.NullTime
	CreatedAt      sql.NullTime
	UpdatedBy      uuid.NullUUID
}

type YoutubeVideo struct {
	ID           uuid.UUID
	TopicID      uuid.UUID
	Title        sql.NullString
	VideoDate    sql.NullTime
	VideoViews   sql.NullString
	ThumbnailUrl sql.NullString
	ExpiryAt     sql.NullTime
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedBy    uuid.NullUUID
}
