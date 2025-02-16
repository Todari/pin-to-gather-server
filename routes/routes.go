package routes

import (
	"github.com/Todari/pin-to-gather-server/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter(boardHandler *api.BoardHandler, websocketHandler *api.WebSocketHandler) *gin.Engine {
    r := gin.Default()

    boardRoutes := r.Group("/board")
    {
        boardRoutes.POST("/", boardHandler.RegisterBoard)
        boardRoutes.GET("/:id", boardHandler.GetBoard)
        boardRoutes.GET("/uuid/:uuid", boardHandler.GetBoardByUuid)
        boardRoutes.PUT("/:id", boardHandler.UpdateBoardTitle)
    }

    wsRoutes := r.Group("/ws")
    {
        wsRoutes.GET("/uuid/:boardUuid", websocketHandler.HandleWebSocket)
    }
    

    return r
}