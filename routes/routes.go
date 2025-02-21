package routes

import (
	"os"

	"github.com/Todari/pin-to-gather-server/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(boardHandler *api.BoardHandler, websocketHandler *api.WebSocketHandler, naverLocalSearchHandler *api.NaverLocalSearchHandler) *gin.Engine {
	r := gin.Default()

	// CORS 설정 추가
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ALLOWED_ORIGIN")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"baggage", "content-type", "sentry-trace"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	boardRoutes := r.Group("/board")
	{
		boardRoutes.POST("", boardHandler.RegisterBoard)
		boardRoutes.GET("/:uuid", boardHandler.GetBoard)
		boardRoutes.PUT("/:uuid", boardHandler.UpdateBoardTitle)
	}

	naverLocalSearchRoutes := r.Group("/naver")
	{
		naverLocalSearchRoutes.GET("/local-search", naverLocalSearchHandler.SearchLocal)
	}

	wsRoutes := r.Group("/ws")
	{
		wsRoutes.GET("/:boardUuid", websocketHandler.HandleWebSocket)
	}

	return r
}
