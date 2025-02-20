package api

import (
	"net/http"

	"github.com/Todari/pin-to-gather-server/services"
	"github.com/gin-gonic/gin"
)

type NaverLocalSearchHandler struct {
	Service *services.LocalSearchService
}

//TODO: naverLocalSearch Cache 추가
func NewNaverLocalSearchHandler(service *services.LocalSearchService) *NaverLocalSearchHandler {
	return &NaverLocalSearchHandler{Service: service}
}

func (h *NaverLocalSearchHandler) SearchLocal(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "검색어를 입력해주세요"})
		return
	}

	searchResult, err := h.Service.SearchLocal(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchResult)
}