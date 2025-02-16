package services

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID        string
	BoardUuid string
	Conn      *websocket.Conn
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Bounds struct {
	Min Coordinate `json:"min"`
	Max Coordinate `json:"max"`
}

type CursorMessage struct {
	UserID    string     `json:"userId"`
	BoardUuid string     `json:"boardUuid"`
	Center    Coordinate `json:"center"`
	Bounds    Bounds     `json:"bounds"`
	Zoom      float64    `json:"zoom"`
}

type WebSocketService struct {
	clients    map[*Client]bool
	clientsMux sync.RWMutex
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{
		clients: make(map[*Client]bool),
	}
}

func (s *WebSocketService) NewClient(userID string, boardUuid string, conn *websocket.Conn) *Client {
	return &Client{
		ID:        userID,
		BoardUuid: boardUuid,
		Conn:      conn,
	}
}

func (s *WebSocketService) AddClient(client *Client) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()
	s.clients[client] = true
}

func (s *WebSocketService) RemoveClient(client *Client) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()
	delete(s.clients, client)
}

func (s *WebSocketService) BroadcastMessage(msg CursorMessage) {
	s.clientsMux.RLock()
	defer s.clientsMux.RUnlock()

	for client := range s.clients {
		log.Printf("Sending message to client: %+v\n", client)
		if client.BoardUuid == msg.BoardUuid && client.ID != msg.UserID {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				continue
			}
		}
	}
}
