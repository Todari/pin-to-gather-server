package routes

import (
	"github.com/Todari/pin-to-gather-server/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter(boardHandler *api.BoardHandler) *gin.Engine {
    r := gin.Default()

    boardRoutes := r.Group("/board")
    {
        boardRoutes.POST("/", boardHandler.RegisterBoard)
        boardRoutes.GET("/:id", boardHandler.GetBoard)
        boardRoutes.GET("/uuid/:uuid", boardHandler.GetBoardByUuid)
        boardRoutes.PUT("/:id", boardHandler.UpdateBoardTitle)
    }

    return r
}