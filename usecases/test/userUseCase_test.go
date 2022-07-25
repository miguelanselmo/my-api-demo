package usecases

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	"github.com/miguelanselmo/my-api-demo/models"
	"github.com/miguelanselmo/my-api-demo/usecases"

	"github.com/miguelanselmo/my-api-demo/utils"

	_ "github.com/mattn/go-sqlite3"
)

var uc *usecases.UseCase

//func initialize() (*usecases.UseCase, error) {
func TestMain(m *testing.M) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	ctx := context.Background()
	repo, err := repository.New(ctx, &log, "file:../../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	if err != nil {
		log.Fatalln(err)
		return //nil, err
	}
	//defer repo.Close()
	//use cases
	uc := usecases.New(ctx, &log, repo)
	uc.GetUserById(1)
	//return //uc, err
}

func TestUseCase_CreateUser(t *testing.T) {
	//uc, _ := initialize()
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		uc      *usecases.UseCase
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			"OK",
			uc,
			args{&models.User{Name: utils.RandomString(20), Email: utils.RandomEmail(), Password: utils.RandomString(6)}},
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.args.user) {
					t.Errorf("UseCase.CreateUser() = %v, want %v", got, tt.want)
				}
			} else if got == nil {
				t.Errorf("UseCase.CreateUser() = %v, want %v", got, "not created")
			}
		})
	}
}

/*
func TestUserUseCase(t *testing.T) {
	t.Log("UserUseCase test")
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	ctx := context.Background()
	repo, err := repository.New(ctx, &log, "file:../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	//assert.Nil(t, err)
	require.NoError(t, err)
	if err != nil {
		t.Error(err)
		return
	}
	defer repo.Close()
	//use cases
	uc := usecases.New(ctx, &log, repo)
	//assert.Nil(t, err)
	require.NoError(t, err)
	if err == nil {
		i := 0
		for ; i < 5; i++ {
			t.Log("UserUseCase CreateUser")
			user, err := uc.CreateUser(&models.User{Name: utils.RandomString(20), Email: utils.RandomEmail()})
			assert.Nil(t, err)
			if err != nil {
				t.Error(err)
			} else {
				t.Log("User created, id:", user.Id)
			}
			t.Log("UserUseCase GetUsers")
			user, err = uc.GetUserById(user.Id)
			//assert.Nil(t, err)
			//assert.Greater(t, user.Id, 0)
			require.NoError(t, err)
			require.Greater(t, user.Id, 0)
			if err != nil {
				t.Error(err)
			} else {
				t.Log("User found, id:", user.Id)
			}
		}
	}
}
*/
func TestUseCase_GetUsersAll(t *testing.T) {
	//uc, _ := initialize()
	tests := []struct {
		name    string
		uc      *usecases.UseCase
		want    *models.User
		wantErr bool
	}{
		{
			"OK",
			uc,
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetUsersAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetUsersAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UseCase.GetUsersAll() = %v, want %v", got, tt.want)
				}
			} else if len(got) == 0 {
				t.Errorf("UseCase.GetUsersAll() = %v, want %v", got, "more than 0 users")
			}
		})
	}
}

func TestUseCase_GetUserById(t *testing.T) {
	//uc, _ := initialize()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		uc      *usecases.UseCase
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			"OK",
			uc,
			args{id: 1},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UseCase.GetUserById() = %v, want %v", got, tt.want)
				}
			} else if got == nil {
				t.Errorf("UseCase.GetUserById() = %v, want %v", got, "want 1 user")
			}
		})
	}
}

func TestUseCase_GetUserAuth(t *testing.T) {
	//uc, _ := initialize()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		uc      *usecases.UseCase
		args    args
		want    *models.Authentication
		wantErr bool
	}{
		{
			"OK",
			uc,
			args{name: "teste1"},
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetUserAuth(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetUserAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UseCase.GetUserAuth() = %v, want %v", got, tt.want)
				}
			} else if got == nil {
				t.Errorf("UseCase.GetUserAuth() = %v, want %v", got, "want 1 user")
			}

		})
	}
}
