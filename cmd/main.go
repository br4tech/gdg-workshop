package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"gdg-workshop/internal/adapters/gemini"
	httpAdapter "gdg-workshop/internal/adapters/http"
	"gdg-workshop/internal/core/usecase"
)

func main() {
	// Configurando o Logger Estruturado Nativo (slog) formatado em JSON para Cloud Logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	// 1. Instancia o Adapter do Gemini
	geminiClient, err := gemini.NewGeminiClient(ctx)
	if err != nil {
		slog.Error("falha catastrófica ao inicializar adapter Gemini", slog.Any("error", err))
		os.Exit(1)
	}

	// 2. Instancia o Caso de Uso injetando o Adapter do Gemini
	explainUC := usecase.NewExplainCodeUseCase(geminiClient)

	// 3. Instancia o Handler HTTP injetando o Caso de Uso
	handler := httpAdapter.NewHTTPHandler(explainUC)

	// 4. Configura as Rotas Modernas (Go 1.22+)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/analyze", handler.AnalyzeCode)

	slog.Info("servidor pronto e escutando na porta :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("erro ao iniciar o servidor", slog.Any("error", err))
	}
}
