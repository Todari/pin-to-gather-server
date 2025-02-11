package services

import (
	"time"

	"github.com/Todari/pin-to-gather-server/models"
	"github.com/Todari/pin-to-gather-server/repository"
	"github.com/google/uuid"
)

type BoardService struct {
    Repo *repository.BoardRepository
}

func NewBoardService(repo *repository.BoardRepository) *BoardService {
    return &BoardService{Repo: repo}
}

func (s *BoardService) RegisterBoard(board *models.Board) error {
		board.Uuid = uuid.New().String()
		now := time.Now()
		board.CreatedAt = now
		board.UpdatedAt = now
    return s.Repo.CreateBoard(board)
}

func (s *BoardService) GetBoard(id uint) (*models.Board, error) {
    return s.Repo.GetBoard(id)
}

func (s *BoardService) GetBoardByUuid(uuid string) (*models.Board, error) {
    return s.Repo.GetBoardByUuid(uuid)
}

func (s *BoardService) UpdateBoardTitle(id uint, newTitle string) (*models.Board, error) {
    board, err := s.Repo.GetBoard(id)
    if err != nil {
        return nil, err
    }

    board.Title = newTitle
    if err := s.Repo.UpdateBoard(board); err != nil {
        return nil, err
    }

    return board, nil
}