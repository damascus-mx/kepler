package user

import (
	"github.com/go-chi/chi"

	_userActions "github.com/damascus-mx/kepler/user/actions"
)

// SetRoutes Sets user routes into the sent router
func SetRoutes(router *chi.Mux) {
	router.Route("/user", func(router chi.Router) {
		router.Get("/", _userActions.GetUser)
	})
}
