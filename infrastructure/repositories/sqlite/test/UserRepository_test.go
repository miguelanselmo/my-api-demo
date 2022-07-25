package repositories

import (
	"context"
	"log"
	"os"
	"testing"

	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	"github.com/miguelanselmo/my-api-demo/models"

	_ "github.com/mattn/go-sqlite3"
)

var (
//sqliteHandler = &SQLiteHandler{}
)

func TestUserRepository(t *testing.T) {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	//t.Skip("Skipping UserRepository test")
	t.Log("UserRepository test")
	ctx := context.Background()
	//sqlConn, err := SetupDatabase(t)
	userRepository, err := repository.New(ctx, &log, "file:../../database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1")
	if err == nil {
		//userRepository := NewDbRepository(ctx, sqlConn) //(sqliteHandler.Conn)
		t.Log("UserRepository CreateUser")
		user, err := userRepository.CreateUser(&models.User{Name: "Test User", Email: "test@email.com"})
		if err != nil {
			t.Error(err)
		} else {
			t.Log("User created, id:", user.Id)
		}
		t.Log("UserRepository GetUsers")
		users, err := userRepository.GetUsers()
		if err != nil {
			t.Error(err)
		} else {
			if len(users) != 0 {
				t.Log("Users: ", users)
				for _, user := range users {
					t.Log("User: ", user)
				}
			}
		}
	}
}
