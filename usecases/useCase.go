package usecases

import (
	"context"
	"log"

	repository "github.com/miguelanselmo/my-api-demo/infrastructure/repositories/sqlite"
	"github.com/miguelanselmo/my-api-demo/models"
)

type UseCase struct {
	repo *repository.DbRepository
	log  *log.Logger
	ctx  context.Context
}

func New(ctx context.Context, log *log.Logger, repo *repository.DbRepository) *UseCase {
	return &UseCase{
		repo: repo,
		ctx:  ctx,
		log:  log,
	}
}

type IUserUseCase interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	DeleteUser(int) error
	AddUserToGroup(int, int) error
	RemoveUserFromGroup(int, int) error
}

type IGroupUseCase interface {
	CreateGroup(context.Context, *models.Group) (*models.Group, error)
	GetGroup(context.Context, int) (*models.Group, error)
	GetGroups(context.Context) ([]*models.Group, error)
	UpdateGroup(context.Context, *models.Group) (*models.Group, error)
	DeleteGroup(context.Context, int) error
}
