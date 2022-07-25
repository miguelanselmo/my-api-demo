package repositories

import (
	"time"

	"github.com/miguelanselmo/my-api-demo/models"
)

func (repo *DbRepository) GetGroups() ([]*models.Group, error) {
	stmt, err := repo.db.Prepare("SELECT ID, NAME FROM GROUPS")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Groups []*models.Group
	for rows.Next() {
		var Group models.Group
		err = rows.Scan(&Group.Id, &Group.Name)
		if err != nil {
			return nil, err
		}
		Groups = append(Groups, &Group)
	}

	return Groups, nil
}

func (repo *DbRepository) GetGroup(id int) (*models.Group, error) {
	stmt, err := repo.db.Prepare("SELECT ID, NAME FROM GROUPS WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var group models.Group
	err = stmt.QueryRow(id).Scan(&group.Id, &group.Name)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (repo *DbRepository) CreateGroup(group *models.Group) (*models.Group, error) {
	stmt, err := repo.db.Prepare("SELECT IIF(MAX(ID) IS NULL, 0, MAX(ID)) ID FROM GROUPS")
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
	stmt, err = repo.db.Prepare("INSERT INTO GROUPS (ID, NAME, CREATEDAT) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	//defer stmt.Close()

	group.Id = int(id + 1)
	_, err = stmt.Exec(group.Id, group.Name, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return group, nil
}
