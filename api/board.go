package api

import (
	"net/http"

	"github.com/Todari/pin-to-gather-server/models"
	"github.com/Todari/pin-to-gather-server/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BoardHandler struct {
	Service *services.BoardService
}

func NewBoardHandler(service *services.BoardService) *BoardHandler {
	return &BoardHandler{Service: service}
}

func (h *BoardHandler) RegisterBoard(c *gin.Context) {
	var board models.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.RegisterBoard(&board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	c.JSON(http.StatusCreated, board)
}

func (h *BoardHandler) GetBoard(c *gin.Context) {
	boardUuid := c.Param("uuid")
	if _, err := uuid.Parse(boardUuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board UUID"})
		return
	}

	board, err := h.Service.GetBoard(boardUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) UpdateBoardTitle(c *gin.Context) {
	var board models.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	boardUuid := c.Param("uuid")
	if _, err := uuid.Parse(boardUuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board UUID"})
		return
	}

	updatedBoard, err := h.Service.UpdateBoardTitle(boardUuid, board.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update board title"})
		return
	}

	c.JSON(http.StatusOK, updatedBoard)
}
