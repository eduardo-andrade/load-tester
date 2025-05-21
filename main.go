package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eduardo-andrade/load-tester/tester"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requisições")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	result := tester.RunLoadTest(*url, *requests, *concurrency)
	fmt.Printf("Resumo da execução:\n")
	fmt.Printf("➡️  Total de requisições: %d\n", result.TotalRequests)
	fmt.Printf("✅ Requisições bem-sucedidas: %d\n", result.SuccessCount)
	fmt.Printf("❌ Requisições com falha: %d\n", result.FailedRequests)
	fmt.Printf("⏱️  Tempo total: %v\n", result.Duration)
	fmt.Println("📊 Status HTTP por código:")

	for code, count := range result.StatusCodeCounts {
		fmt.Printf("   - %d: %d respostas\n", code, count)
	}

}
