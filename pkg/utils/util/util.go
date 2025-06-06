package util

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/easc01/mindo-server/pkg/utils/constant"
	"github.com/google/uuid"
)

var adjectives = []string{
	"Swift", "Clever", "Brave", "Witty", "Calm", "Mighty", "Happy", "Silent", "Lucky", "Bold",
}

var nouns = []string{
	"Fox", "Eagle", "Tiger", "Panda", "Wolf", "Hawk", "Bear", "Lion", "Otter", "Shark",
}

// Returns sql.NullString for a string
func GetSQLNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != constant.Blank}
}

// Returns sql.NullString for a string
func GetNullUUID(s uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: s, Valid: true}
}

// ConvertStringToUUID converts a string to a uuid.UUID
func ConvertStringToUUID(id string) uuid.UUID {
	parsedId, _ := uuid.Parse(id)
	return parsedId
}

// Generate a random username with UUID suffix
func GenerateUsername() string {
	uid := uuid.New()
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	// Combine random word with UUID suffix (first 8 characters)
	username := fmt.Sprintf("%s%s-%s", adjective, noun, uid.String()[:8])

	return username
}

func GenerateHexCode(num int) string {
	// Format the integer as a hexadecimal string with padding to ensure it's 6 characters long
	hexCode := fmt.Sprintf("%06X", num)
	return hexCode
}

func GetUUIDFromString(s string) uuid.NullUUID {
	if s != constant.Blank {
		parsedUUID, err := uuid.Parse(s)

		if err != nil {
			return uuid.NullUUID{
				Valid: false,
			}
		}

		return uuid.NullUUID{
			UUID:  parsedUUID,
			Valid: true,
		}
	}

	return uuid.NullUUID{
		Valid: false,
	}
}

func ReverseSlice[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func GetGrade(totalAns, correctAns int) string {
	if totalAns == 0 {
		return "N/A"
	}

	percentage := float64(correctAns) / float64(totalAns) * 100

	switch {
	case percentage >= 90:
		return "A+"
	case percentage >= 80:
		return "A"
	case percentage >= 70:
		return "B"
	case percentage >= 60:
		return "C"
	case percentage >= 50:
		return "D"
	default:
		return "F"
	}
}
