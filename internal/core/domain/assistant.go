package domain

// CodeQuery representa a dúvida de código do aluno
type CodeQuery struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// CodeExplanation representa a resposta tratada da IA
type CodeExplanation struct {
	Summary     string   `json:"summary"`
	Suggestions []string `json:"suggestions"`
}


