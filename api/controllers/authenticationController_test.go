package controllers

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/miguelanselmo/my-api-demo/config"
	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	usecase "github.com/miguelanselmo/my-api-demo/usecases"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

func ServiceContainerTest() (*Controller, error) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	viper.SetConfigName(".env")
	viper.AddConfigPath("../../")
	config, err := config.LoadConfig(os.Getenv("../../"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	ctx := context.Background()
	//ctx.WithTimeout(time.Second * 10)
	//repository
	repo, err := repository.New(ctx, &log, config.DBSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = repo.Open()
	if err != nil {
		return nil, err
	}
	defer repo.Close()

	//use cases
	uc := usecase.New(ctx, &log, repo)
	//controller
	ctrl := New(ctx, &log, &config, uc)
	//ctrl.Start(config.ServerAddress)
	return ctrl, nil
}

func TestController_Auth(t *testing.T) {
	c, _ := ServiceContainerTest()
	type args struct {
		ctx *gin.Context
	}
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBufferString(`{"name": "admin", "password": "admin"}`))

	tests := []struct {
		name string
		ctrl *Controller
		args args
	}{
		{
			"TestController_Auth",
			c,
			args{ctx},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctrl.Auth(tt.args.ctx)
		})
	}
}
