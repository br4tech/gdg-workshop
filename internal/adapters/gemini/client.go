package gemini

import (
	"context"
	"fmt"
	"os"

	"gdg-workshop/internal/core/domain"
	"gdg-workshop/internal/core/ports"

	"google.golang.org/genai"
)

type geminiClient struct {
	client *genai.Client
}

func NewGeminiClient(ctx context.Context) (ports.AIService, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("a variável GEMINI_API_KEY não foi configurada")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("falha ao instanciar genai client: %w", err)
	}

	return &geminiClient{client: client}, nil
}

func (g *geminiClient) ExplainCode(ctx context.Context, query domain.CodeQuery) (*domain.CodeExplanation, error) {
	prompt := fmt.Sprintf(
		"Como um tutor de programação didático, analise este código em %s e forneça uma explicação rápida "+
			"de como ele funciona e uma melhoria de performance/segurança.\n\nCódigo:\n```%s\n%s\n```",
		query.Language, query.Language, query.Code,
	)

	// Utilizando o modelo gemini-2.5-flash recomendável para velocidade e baixo custo
	resp, err := g.client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(prompt), nil)
	if err != nil {
		return nil, fmt.Errorf("erro de comunicação com o modelo Gemini: %w", err)
	}

	// Por simplicidade do tempo do workshop, estruturamos de forma fixa.
	return &domain.CodeExplanation{
		Summary: resp.Text(),
		Suggestions: []string{
			"Considere documentar assinaturas complexas",
			"Evite aninhamento excessivo em blocos lógicos",
		},
	}, nil
}
