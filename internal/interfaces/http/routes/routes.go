package routes

import (
	"github.com/cnt-777/internal/interfaces/http/handlers"
	"github.com/cnt-777/internal/interfaces/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	// client
	r.GET("/", handlers.HomePage)
	r.GET("/blackjack", handlers.BlackjackPage)

	r.GET("ws/blackjack", handlers.BlackjackWS)

	// staff
	r.GET("/dealer", middlewares.IsStaff, handlers.DealerPage)
	r.GET("/scanner", middlewares.IsStaff, handlers.ScannerPage)

	r.GET("ws/dealer", handlers.DealerWS)
	r.GET("ws/scanner", handlers.ScannerWS)
}
