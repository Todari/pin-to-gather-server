package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Todari/pin-to-gather-server/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Service *services.WebSocketService
}

func NewWebSocketHandler(service *services.WebSocketService) *WebSocketHandler {
	return &WebSocketHandler{Service: service}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == os.Getenv("ALLOWED_ORIGIN")
	},
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	userID := c.Query("userId")
	boardUuid := c.Param("boardUuid")

	log.Printf("New client connected: %s\n", boardUuid)

	client := h.Service.NewClient(userID, boardUuid, conn)
	h.Service.AddClient(client)

	for {
		var msg services.CursorMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Read Error:", err)
			}
			h.Service.RemoveClient(client)
			break
		}
		if msg.UserID == "" || msg.BoardUuid == "" {
			log.Println("Invalid message received")
			continue
		}
		h.Service.BroadcastMessage(msg)
	}
}
