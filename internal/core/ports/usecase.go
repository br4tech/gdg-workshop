package ports

import (
	"context"
	"gdg-workshop/internal/core/domain"
)

// AssistantUseCase é nossa Port de Entrada (Input Port)
type AssistantUseCase interface {
	Execute(ctx context.Context, query domain.CodeQuery) (*domain.CodeExplanation, error)
}

// AIService é nossa Port de Saída (Output Port)
// Qualquer motor de IA (Gemini, Mock, etc) deve satisfazer esta interface
type AIService interface {
	ExplainCode(ctx context.Context, query domain.CodeQuery) (*domain.CodeExplanation, error)
}
