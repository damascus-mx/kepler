package routes

import (
	"github.com/go-chi/chi"

	"github.com/damascus-mx/kepler/src/usecases"
)

// UserRoutes Exports all user routing
type UserRoutes struct {
}

// SetRoutes Sets user routes into the sent router
func (r UserRoutes) SetRoutes(router *chi.Mux) {
	_userActions := usecases.UserActions{}
	router.Route("/user", func(router chi.Router) {
		router.Get("/", _userActions.GetUser)
	})
}
