package api

import (
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	board, err := h.Service.GetBoard(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) GetBoardByUuid(c *gin.Context) {
	boardUuid := c.Param("uuid")
	if _, err := uuid.Parse(boardUuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board UUID"})
		return
	}

	board, err := h.Service.GetBoardByUuid(boardUuid)
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	updatedBoard, err := h.Service.UpdateBoardTitle(uint(id), board.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update board title"})
		return
	}

	c.JSON(http.StatusOK, updatedBoard)
}