package usecase

import (
	"context"
	"errors"
	"gdg-workshop/internal/core/domain"
	"gdg-workshop/internal/core/ports"
	"log/slog"
)

type explainCodeUseCase struct {
	aiService ports.AIService
}

// NewExplainCodeUseCase injeta a dependência via Interface
func NewExplainCodeUseCase(ai ports.AIService) ports.AssistantUseCase {
	return &explainCodeUseCase{aiService: ai}
}

func (u *explainCodeUseCase) Execute(ctx context.Context, query domain.CodeQuery) (*domain.CodeExplanation, error) {
	if query.Code == "" {
		slog.WarnContext(ctx, "tentativa de requisição sem código")
		return nil, errors.New("o código fornecido não pode ser vazio")
	}

	slog.InfoContext(ctx, "iniciando análise de código", slog.String("lang", query.Language))
	return u.aiService.ExplainCode(ctx, query)
}
