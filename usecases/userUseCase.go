package usecases

import (
	"errors"

	"github.com/miguelanselmo/my-api-demo/models"
	"github.com/miguelanselmo/my-api-demo/utils"
)

func (uc *UseCase) CreateUser(user *models.User) (*models.User, error) {
	u, err := uc.repo.GetUserByName(user.Name)
	if err != nil {
		return nil, err
	} else if u != nil {
		return nil, errors.New("user already exists")
	}
	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		return nil, err
	} else {
		user.Password = hashedPassword
	}
	return uc.repo.CreateUser(user)
}

func (uc *UseCase) GetUsersAll() ([]*models.User, error) {
	return uc.repo.GetUsers()
}

func (uc *UseCase) GetUserById(id int) (*models.User, error) {
	return uc.repo.GetUserById(id)
}

func (uc *UseCase) GetUserAuth(name string) (*models.Authentication, error) {
	user, err := uc.repo.GetUserByName(name)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.New("user not found")
	}
	return &models.Authentication{
		UserId:   user.Id,
		UserName: user.Name,
		Password: user.Password,
	}, nil
}

func (uc *UseCase) UpdateUser(user *models.User) (*models.User, error) {
	return uc.repo.UpdateUser(user)
}

func (uc *UseCase) DeleteUser(id int) error {
	return uc.repo.DeleteUser(id)
}

func (uc *UseCase) AddUserToGroup(userID int, groupID int) error {
	return errors.New("not implemented")
	//return uc.repo.AddUserToGroup(userID, groupID)
}

func (uc *UseCase) RemoveUserFromGroup(userID int, groupID int) error {
	return errors.New("not implemented")
	//return uc.repo.RemoveUserFromGroup(userID, groupID)
}
