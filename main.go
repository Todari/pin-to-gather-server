package main

import (
	"log"

	"github.com/Todari/pin-to-gather-server/api"
	"github.com/Todari/pin-to-gather-server/config"
	"github.com/Todari/pin-to-gather-server/database"
	"github.com/Todari/pin-to-gather-server/repository"
	"github.com/Todari/pin-to-gather-server/routes"
	"github.com/Todari/pin-to-gather-server/services"
)

func main() {
    config.LoadConfig()    // 환경 변수 로드
    database.ConnectDatabase() // PostgreSQL 연결

    boardRepository := repository.NewBoardRepository(database.DB)
    boardService := services.NewBoardService(boardRepository)
    boardHandler := api.NewBoardHandler(boardService)

    websocketService := services.NewWebSocketService()
    websocketHandler := api.NewWebSocketHandler(websocketService)

    r := routes.SetupRouter(boardHandler, websocketHandler) // *gin.Engine 반환

    port := config.AppConfig.ServerPort
    log.Println("🚀 Server running on port:", port)
    r.Run(":" + port)
}
