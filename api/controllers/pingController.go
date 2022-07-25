package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
