package repositories

import (
	"context"
	"log"
	"os"
	"testing"

	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	"github.com/miguelanselmo/my-api-demo/models"
	"github.com/miguelanselmo/my-api-demo/utils"
)

func TestDbRepository_GetGroups(t *testing.T) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	ctx := context.Background()
	dbRepository, _ := repository.New(ctx, &log, "file:../../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	tests := []struct {
		name    string
		repo    *repository.DbRepository
		want    []*models.Group
		wantErr bool
	}{
		{
			"GetGroups",
			dbRepository,
			[]*models.Group{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.repo.GetGroups()
			if (err != nil) != tt.wantErr {
				t.Errorf("DbRepository.GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbRepository.GetGroups() = %v, want %v", got, tt.want)
			}*/
		})
	}
}

func TestDbRepository_GetGroup(t *testing.T) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	ctx := context.Background()
	dbRepository, _ := repository.New(ctx, &log, "file:../../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		repo    *repository.DbRepository
		args    args
		want    *models.Group
		wantErr bool
	}{
		{
			"GetGroup",
			dbRepository,
			args{1},
			&models.Group{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.repo.GetGroup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DbRepository.GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbRepository.GetGroup() = %v, want %v", got, tt.want)
			}*/
		})
	}
}

func TestDbRepository_CreateGroup(t *testing.T) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	ctx := context.Background()
	dbRepository, _ := repository.New(ctx, &log, "file:../../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	type args struct {
		group *models.Group
	}
	tests := []struct {
		name    string
		repo    *repository.DbRepository
		args    args
		want    *models.Group
		wantErr bool
	}{
		{
			"CreateGroup",
			dbRepository,
			args{&models.Group{Name: utils.RandomString(20)}},
			&models.Group{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.repo.CreateGroup(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("DbRepository.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbRepository.CreateGroup() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
