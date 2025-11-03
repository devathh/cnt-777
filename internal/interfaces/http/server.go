package server

import (
	"github.com/cnt-777/internal/interfaces/http/handlers"
	"github.com/cnt-777/internal/interfaces/http/routes"
	"github.com/gin-gonic/gin"
)

func New(handler *handlers.Handler) *gin.Engine {
	r := gin.Default()

	// TODO: uncomment
	// r.LoadHTMLGlob("./web/templates/*")
	r.Static("css", "./web/static/css/")
	r.Static("js", "./web/static/js/")

	routes.Setup(r, handler)

	return r
}
