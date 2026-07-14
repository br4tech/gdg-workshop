package http

import (
	"encoding/json"
	"gdg-workshop/internal/core/domain"
	"gdg-workshop/internal/core/ports"
	"log/slog"
	"net/http"
)

type HTTPHandler struct {
	useCase ports.AssistantUseCase
}

func NewHTTPHandler(uc ports.AssistantUseCase) *HTTPHandler {
	return &HTTPHandler{useCase: uc}
}

func (h *HTTPHandler) AnalyzeCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "chamada HTTP recebida", slog.String("method", r.Method))

	var query domain.CodeQuery
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		slog.ErrorContext(ctx, "falha ao decodificar payload", slog.Any("error", err))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	explanation, err := h.useCase.Execute(ctx, query)
	if err != nil {
		slog.ErrorContext(ctx, "erro na execução do caso de uso", slog.Any("error", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(explanation)
}
