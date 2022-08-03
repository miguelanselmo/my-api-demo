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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
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
