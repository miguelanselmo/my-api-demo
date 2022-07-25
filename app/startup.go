package app

import (
	"context"
	"log"
	"os"

	ctrl "github.com/miguelanselmo/my-api-demo/api/controllers"
	"github.com/miguelanselmo/my-api-demo/config"
	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	usecase "github.com/miguelanselmo/my-api-demo/usecases"

	_ "github.com/mattn/go-sqlite3"
)

func StartApplication() {
	ServiceContainer()
}

/*
type IServiceContainer interface {
	InjectDependencies()
}
*/
func ServiceContainer() error {
	log := log.Logger{}
	log.SetOutput(os.Stdout)

	config, err := config.LoadConfig(".") //(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	ctx := context.Background()
	//ctx.WithTimeout(time.Second * 10)
	//repository
	repo, err := repository.New(ctx, &log, config.DBSource)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = repo.Open()
	if err != nil {
		return err
	}
	defer repo.Close()

	//use cases
	uc := usecase.New(ctx, &log, repo)
	//controller
	ctrl := ctrl.New(ctx, &log, &config, uc)
	ctrl.Start(config.ServerAddress)
	return nil
}
