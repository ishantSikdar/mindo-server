package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ishantSikdar/mindo-server/internal/config"
	"github.com/ishantSikdar/mindo-server/internal/models"
	"github.com/ishantSikdar/mindo-server/pkg/db"
	"github.com/ishantSikdar/mindo-server/pkg/logger"
	"github.com/ishantSikdar/mindo-server/pkg/structs"
	"github.com/ishantSikdar/mindo-server/pkg/utils"
	"google.golang.org/api/idtoken"
)

func GoogleAuthService(c context.Context, googleReq structs.GoogleLoginRequest) (structs.AppUserDataDTO, error) {

	payload, payloadErr := idtoken.Validate(c, googleReq.IDToken, config.GetConfig().GoogleClientId)
	if payloadErr != nil {
		logger.Log.Errorf("invalidate app user token %s for token %s", payloadErr, googleReq.IDToken)
		return structs.AppUserDataDTO{}, payloadErr
	}

	name, _ := payload.Claims["name"].(string)
	email, _ := payload.Claims["email"].(string)

	appUserParams := structs.NewAppUserParams{
		Name:     name,
		Email:    email,
		Username: payload.Subject,
		Mobile:   "",
	}

	// Check if appUser exists by username/subject
	appUser, appUserErr := db.Queries.UpdateUserLastLoginAtByUsername(c, utils.GetSQLNullString(appUserParams.Username))

	if appUserErr != nil {
		if errors.Is(appUserErr, sql.ErrNoRows) {
			// Create new user
			newAppUser, newAppUserErr := CreateNewAppUser(appUserParams)
			if newAppUserErr != nil {
				logger.Log.Errorf("failed to create app user %s for email %s username %s", newAppUserErr, appUserParams.Email, appUserParams.Username)
				return structs.AppUserDataDTO{}, newAppUserErr
			}

			logger.Log.Info("new app user created", newAppUser.UserID)
			return newAppUser, nil
		}

		// Log unexpected DB errors
		logger.Log.Errorf("failed to update last login %s for username %s", appUserErr, appUserParams.Username)
		return structs.AppUserDataDTO{}, appUserErr
	}

	return structs.AppUserDataDTO{
		UserID:            appUser.UserID,
		Username:          appUser.Username,
		ProfilePictureUrl: appUser.ProfilePictureUrl,
		Bio:               appUser.Bio,
		Name:              appUser.Name,
		Mobile:            appUser.Mobile,
		Email:             appUser.Email,
		LastLoginAt:       appUser.LastLoginAt,
		UpdatedAt:         appUser.UpdatedAt,
		CreatedAt:         appUser.CreatedAt,
		UpdatedBy:         appUser.UpdatedBy,
		UserType:          models.UserTypeAppUser,
	}, nil
}
