package core

import "github.com/go-chi/chi"

// IRoute Route contract
type IRoute interface {
	SetRoutes(router *chi.Mux)
}
