package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/miguelanselmo/my-api-demo/api"
	"github.com/miguelanselmo/my-api-demo/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/techschool/simplebank/util"
)

type authResponse struct {
	SessionID             uuid.UUID   `json:"session_id"`
	AccessToken           string      `json:"access_token"`
	AccessTokenExpiresAt  time.Time   `json:"access_token_expires_at"`
	RefreshToken          string      `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time   `json:"refresh_token_expires_at"`
	User                  models.User `json:"user"`
}

func (ctrl *Controller) Auth(ctx *gin.Context) {
	var req models.Authentication
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}

	auth, err := ctrl.uc.GetUserAuth(req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, api.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, auth.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, api.ErrorResponse(err))
		return
	}

	accessToken, accessPayload, err := ctrl.token.CreateToken(
		auth.UserName,
		ctrl.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := ctrl.token.CreateToken(
		auth.UserName,
		ctrl.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}
	/*
		session, err := ctrl.store.CreateSession(ctx, db.CreateSessionParams{
			ID:           refreshPayload.ID,
			Username:     user.Username,
			RefreshToken: refreshToken,
			UserAgent:    ctx.Request.UserAgent(),
			ClientIp:     ctx.ClientIP(),
			IsBlocked:    false,
			ExpiresAt:    refreshPayload.ExpiredAt,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	*/
	user, _ := ctrl.uc.GetUserById(auth.UserId)
	rsp := authResponse{
		SessionID:             uuid.New(), //session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  *user,
	}
	ctx.JSON(http.StatusOK, rsp)
}
