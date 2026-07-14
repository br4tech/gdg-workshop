package usecase_test

import (
	"context"
	"gdg-workshop/internal/core/domain"
	"gdg-workshop/internal/core/usecase"
	"testing"
)

// mockAIService simula o comportamento da API do Gemini para testes rápidos
type mockAIService struct{}

func (m *mockAIService) ExplainCode(ctx context.Context, query domain.CodeQuery) (*domain.CodeExplanation, error) {
	return &domain.CodeExplanation{
		Summary:     "Este é um mock didático para testar o fluxo hexagonal",
		Suggestions: []string{"Utilize testes robustos"},
	}, nil
}

func TestExplainCodeUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockAI := &mockAIService{}
	uc := usecase.NewExplainCodeUseCase(mockAI)
	query := domain.CodeQuery{Code: "fmt.Println(\"GDG\")", Language: "go"}

	// Act
	resp, err := uc.Execute(context.Background(), query)

	// Assert
	if err != nil {
		t.Fatalf("esperava erro nulo, retornou: %v", err)
	}
	if resp.Summary != "Este é um mock didático para testar o fluxo hexagonal" {
		t.Errorf("resposta inesperada do resumo: %s", resp.Summary)
	}
}
