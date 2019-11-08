package user

import (
	"github.com/go-chi/chi"

	user "github.com/damascus-mx/kepler/user/actions"
)

// Routes Exports all user routing
type Routes struct {
}

// SetRoutes Sets user routes into the sent router
func (r Routes) SetRoutes(router *chi.Mux) {
	actions := user.Actions{}
	router.Route("/user", func(router chi.Router) {
		router.Get("/", actions.GetUser)
	})
}
