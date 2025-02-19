package repository

import (
	"github.com/Todari/pin-to-gather-server/models"
	"gorm.io/gorm"
)

type BoardRepository struct {
    DB *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepository {
    return &BoardRepository{DB: db}
}

func (r *BoardRepository) CreateBoard(board *models.Board) error {
    return r.DB.Create(board).Error
}

func (r *BoardRepository) GetBoard(uuid string) (*models.Board, error) {
    var board models.Board
    if err := r.DB.Where("uuid = ?", uuid).First(&board).Error; err != nil {
        return nil, err
    }
    return &board, nil
}

func (r *BoardRepository) UpdateBoard(board *models.Board) error {
    return r.DB.Save(board).Error
}