package userhandler

import (
	"net/http"

	"github.com/easc01/mindo-server/internal/middleware"
	"github.com/easc01/mindo-server/internal/models"
	userservice "github.com/easc01/mindo-server/internal/services/user_service"
	"github.com/easc01/mindo-server/pkg/logger"
	"github.com/easc01/mindo-server/pkg/utils/constant"
	"github.com/easc01/mindo-server/pkg/utils/message"
	networkutil "github.com/easc01/mindo-server/pkg/utils/network_util"
	"github.com/easc01/mindo-server/pkg/utils/route"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterAppUserRoutes(rg *gin.RouterGroup) {
	userRg := rg.Group(route.User, middleware.RequireRole(models.UserTypeAppUser))

	{
		userRg.GET(constant.Blank, getAppUser)
		userRg.GET(constant.IdParam, getAppUserByID)
	}

}

func getAppUser(c *gin.Context) {
	user, ok := middleware.GetUser(c)

	if user.AppUser == nil || !ok {
		logger.Log.Errorf(message.NullAppUserContext)
		networkutil.NewErrorResponse(
			http.StatusInternalServerError,
			message.SomethingWentWrong,
			message.NullAppUserContext,
		).Send(c)
	}

	networkutil.NewResponse(
		http.StatusAccepted,
		user.AppUser,
	).Send(c)
}

func getAppUserByID(c *gin.Context) {
	paramId := c.Param("id")

	parsedId, parseErr := uuid.Parse(paramId)
	if parseErr != nil {
		networkutil.NewErrorResponse(
			http.StatusBadRequest,
			message.InvalidUserID,
			parseErr.Error(),
		).Send(c)
		return
	}

	user, statusCode, userErr := userservice.GetAppUserByUserID(parsedId)

	if userErr != nil {
		logger.Log.Errorf("failed to get user %s userID: %s", userErr, parsedId)
		networkutil.NewErrorResponse(
			statusCode,
			message.SomethingWentWrong,
			userErr.Error(),
		).Send(c)
		return
	}

	networkutil.NewResponse(
		http.StatusAccepted,
		user,
	).Send(c)
}
