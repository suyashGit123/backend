package users

import (
	"errors"

	"github.com/suyashGit123/backend/models"
)

func Update(id string, usr *models.User) (*models.User, error) {

	var exists bool
	for index, user := range mockUsers {
		if user.ID == id {
			usr.ID = mockUsers[index].ID
			usr.Password = mockUsers[index].Password
			usr.Email = mockUsers[index].Email

			mockUsers[index] = *usr
			exists = true
		}

	}

	if !exists {
		return nil, errors.New("unable to update user because user was not found")
	}

	return usr, nil
}
