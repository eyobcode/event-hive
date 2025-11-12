package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.Default()
	// define routes here later
	return router
}
