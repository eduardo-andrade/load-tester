package tester

import (
	"fmt"
)

type TestResult struct {
	Duration      string
	TotalRequests int
	SuccessCount  int
	StatusCodes   map[int]int
}

func (r TestResult) GenerateReport() string {
	report := fmt.Sprintf(
		"\n\n===== Relatório de Teste =====\n"+
			"Tempo total: %v\n"+
			"Total de requisições: %d\n"+
			"Respostas 200 OK: %d\n"+
			"Distribuição de status:\n",
		r.Duration,
		r.TotalRequests,
		r.SuccessCount,
	)

	for code, count := range r.StatusCodes {
		report += fmt.Sprintf("  - %d: %d\n", code, count)
	}
	report += "==============================\n"
	return report
}
