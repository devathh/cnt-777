package routes

import (
	"github.com/cnt-777/internal/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, handler *handlers.Handler) {
	// client
	r.GET("/", handler.HomePage())
	r.GET("/blackjack", handler.BlackjackPage())

	r.GET("ws/blackjack", handler.BlackjackWS())

	// staff
	r.GET("/dealer", handler.DealerPage())
	r.GET("/scanner", handler.ScannerPage())

	r.GET("ws/dealer", handler.DealerWS())
	r.GET("ws/scanner", handler.ScannerWS())
}
