package users

import "github.com/suyashGit123/backend/models"

func Index() (users *models.Users, err error) {
	users = &mockUsers

	return
}
