package routes

import (
	"github.com/go-chi/chi"

	"github.com/damascus-mx/kepler/src/core"
	"github.com/damascus-mx/kepler/src/usecases"
)

// UserRoutes Exports all user routing
type UserRoutes struct {
}

// SetRoutes Sets user routes into the sent router
func (r UserRoutes) SetRoutes(router *chi.Mux) {
	var _userActions core.IUsecase = usecases.UserActions{}
	router.Route("/user", func(router chi.Router) {
		router.Get("/", _userActions.GetAll)
		router.Post("/", _userActions.Create)

		router.Route("/{userID}", func(router chi.Router) {
			// router.Use(_userActions.Context)
			router.Get("/", _userActions.GetByID)
		})
	})
}
