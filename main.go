package main

import (
	"net/http"

	_app "github.com/damascus-mx/kepler/bin"
)

func main() {
	// Init root router
	/*
	*	TODO: Add PosgreSQL (pg) auth with driver
	 */
	http.ListenAndServe(": 3000", _app.InitApplication())
}
