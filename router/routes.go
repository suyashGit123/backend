package router

import (
	"github.com/suyashGit123/backend/models"
	HomeHandler "github.com/suyashGit123/backend/routes/home"
	StatusHandler "github.com/suyashGit123/backend/routes/status"
)

func GetRoutes() models.Routes {
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
	}
}
