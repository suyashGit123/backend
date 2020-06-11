package router

import (
	"github.com/gorilla/mux"
	V1Router "github.com/suyashGit123/backend/routes/v1"
)

type Service interface {
	GetRawRouter() *mux.Router
}

func GetRouter() Service {

	r := Router{
		RawRouter: mux.NewRouter().StrictSlash(true),
	}

	for _, route := range GetRoutes() {
		r.RawRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	for name, pack := range V1Router.GetRoutes() {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return r
}
