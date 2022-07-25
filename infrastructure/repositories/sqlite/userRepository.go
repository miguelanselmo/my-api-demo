package repositories

import (
	"time"

	"github.com/miguelanselmo/my-api-demo/models"
)

func (repo *DbRepository) GetUsers() ([]*models.User, error) {
	stmt, err := repo.db.Prepare("SELECT ID, NAME, EMAIL, IIF(GROUPID IS NULL, -1, GROUPID) GROUPID FROM USERS LIMIT 10 OFFSET 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.GroupName)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repo *DbRepository) GetUserById(id int) (*models.User, error) {
	stmt, err := repo.db.Prepare("SELECT ID, NAME, EMAIL, IIF(GROUPID IS NULL, -1, GROUPID) GROUPID FROM USERS WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.GroupName)
	if err != nil {
		if err.Error() == Error_NoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (repo *DbRepository) GetUserByName(name string) (*models.User, error) {
	stmt, err := repo.db.Prepare("SELECT ID, NAME, PASSWORD FROM USERS WHERE NAME = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(name).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		if err.Error() == Error_NoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (repo *DbRepository) CreateUser(user *models.User) (*models.User, error) {
	stmt, err := repo.db.Prepare("SELECT IIF(MAX(ID) IS NULL, 0, MAX(ID)) ID FROM USERS")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id = 0
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	}
	stmt, err = repo.db.Prepare("INSERT INTO USERS (ID, NAME, PASSWORD, EMAIL, CREATEDAT) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	user.Id = int(id + 1)
	_, err = stmt.Exec(user.Id, user.Name, user.Password, user.Email, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *DbRepository) UpdateUser(user *models.User) (*models.User, error) {
	stmt, err := repo.db.Prepare("UPDATE USERS SET NAME = ?, EMAIL = ? WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *DbRepository) DeleteUser(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM USERS WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
