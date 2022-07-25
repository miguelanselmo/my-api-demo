package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/miguelanselmo/my-api-demo/models"
)

//DbRepositoryImpl is the implementation of the Repository interface
type DbRepository struct {
	db      *sql.DB
	ctx     context.Context
	log     *log.Logger
	connStr string
}

type IDbRepository interface {
	//New(context.Context, string) (*DbRepositoryImpl, error)
	Open() (*DbRepository, error)
	Close() error
}

const Error_NoRows = "sql: no rows in result set"

//creates a new DbRepository
func New(ctx context.Context, log *log.Logger, connStr string) (*DbRepository, error) {
	db, err := sql.Open("sqlite3", connStr)
	return &DbRepository{
		db:      db,
		ctx:     ctx,
		log:     log,
		connStr: connStr,
	}, err
}

func (repo *DbRepository) Open() (*DbRepository, error) {
	err := repo.db.Ping()
	if err != nil {
		log.Println("Opening database connection")
		db, err := sql.Open("sqlite3", repo.connStr)
		repo.db = db
		return repo, err
	} else {
		return repo, err
	}

}

func (repo *DbRepository) Close() error {
	return repo.db.Close()
}

//Database accessor for the User table
type IUserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int) error
	AddUserToGroup(userID int, groupID int) error
	RemoveUserFromGroup(userID int, groupID int) error
}

//Database accessor for the Group table
type IGroupRepository interface {
	CreateGroup(group *models.Group) (*models.Group, error)
	GetGroup(id int) (*models.Group, error)
	GetGroups() ([]*models.Group, error)
	UpdateGroup(group *models.Group) (*models.Group, error)
	DeleteGroup(id int) error
}
