package users

import (
	"github.com/suyashGit123/backend/models"
)

var mockUsers models.Users

func init() {
	usr, _ := models.NewUser("Suyash", "Sardeshpande", "ss@gmail.com", "123")

	mockUsers = models.Users{*usr}
}
