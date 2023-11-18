package services

import (
	"context"

	"github.com/andriyskachko/zephyr-api/repositories"
)

type TextService struct {
	ctx            context.Context
	textRepository repositories.TextRepository
}

func NewTextService(ctx context.Context, textRepository repositories.TextRepository) *TextService {
	return &TextService{textRepository: textRepository, ctx: ctx}
}

func (s *TextService) GetTexts() ([]repositories.Text, error) {
	texts, err := s.textRepository.All(s.ctx)
	if err != nil {
		return nil, err
	}

	return texts, nil
}

func (s *TextService) CreateText(title, content string) (*repositories.Text, error) {
	text := repositories.NewText(title, content)

	return text, nil
}
