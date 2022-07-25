package controllers

import (
	"context"
	"log"

	"github.com/miguelanselmo/my-api-demo/api/auth/token"
	"github.com/miguelanselmo/my-api-demo/api/middleware"
	"github.com/miguelanselmo/my-api-demo/config"
	usecase "github.com/miguelanselmo/my-api-demo/usecases"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc     *usecase.UseCase
	ctx    context.Context
	log    *log.Logger
	router *gin.Engine
	token  token.IToken
	config *config.Config
}

type IController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}

func New(ctx context.Context, log *log.Logger, config *config.Config, uc *usecase.UseCase) *Controller {
	router := gin.Default()
	token, err := token.New(config.TokenKey)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Controller{
		uc:     uc,
		ctx:    ctx,
		log:    log,
		router: router,
		token:  token,
		config: config,
	}

}
func (ctrl *Controller) Start(addr string) error {
	ctrl.router.GET("/ping", ctrl.Ping)
	ctrl.router.POST("/login", ctrl.Auth)
	ctrl.router.POST("/users", ctrl.CreateUser)

	authRouter := ctrl.router.Group("/").Use(middleware.AuthMiddleware(ctrl.token))
	authRouter.GET("/users", ctrl.GetUserAll)
	authRouter.GET("/users/:id", ctrl.GetUserById)
	authRouter.PUT("/users/:id", ctrl.UpdateUser)

	err := ctrl.router.Run(addr)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
