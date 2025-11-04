package handlers

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex
	log      *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{
		upgrader: websocket.Upgrader{
			CheckOrigin:     func(r *http.Request) bool { return true },
			ReadBufferSize:  4128,
			WriteBufferSize: 4128,
		},
		clients: make(map[*websocket.Conn]bool),
		mu:      sync.Mutex{},
		log:     log,
	}
}

func (h *Handler) HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	}
}

func (h *Handler) BlackjackPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *Handler) BlackjackWS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.log.Warn("failed connection to the client's websocket", slog.String("error", err.Error()))
			ctx.JSON(http.StatusInternalServerError, "failed connect to websocket")
		}
		defer conn.Close()

		h.mu.Lock()
		h.clients[conn] = true
		h.mu.Unlock()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				h.log.Warn("failed to read message from client", slog.String("error", err.Error()))
				break
			}
		}

		h.mu.Lock()
		delete(h.clients, conn)
		h.mu.Unlock()
	}
}

func (h *Handler) DealerPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *Handler) ScannerPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *Handler) DealerWS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.log.Warn("failed connection to the dealer's websocket", slog.String("error", err.Error()))
			ctx.JSON(http.StatusInternalServerError, "failed to connect to websocket")
			return
		}
		defer conn.Close()

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				h.log.Warn("failed to read message from dealer's websocket", slog.String("error", err.Error()))
				break
			}

			h.broadcastMessage(msgType, msg)
		}
	}
}

func (h *Handler) ScannerWS() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *Handler) broadcastMessage(msgType int, msg []byte) {
	for clientConn := range h.clients {
		if err := clientConn.WriteMessage(msgType, msg); err != nil {
			h.log.Warn("failed to send msg to client from broadcast", slog.String("error", err.Error()))
		}
	}
}
