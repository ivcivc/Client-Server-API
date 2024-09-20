package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotacao, err := fetchCotacao(ctx)
	if err != nil {
		log.Printf("Erro ao buscar cotação: %v", err)
		os.Exit(1)
	}

	err = saveCotacao(cotacao)
	if err != nil {
		log.Printf("Erro ao salvar cotação: %v", err)
		os.Exit(1)
	}

	fmt.Println("Cotação salva com sucesso!")
}

func fetchCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha na requisição: status %d", resp.StatusCode)
	}

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func saveCotacao(cotacao *Cotacao) error {
	content := fmt.Sprintf("Dólar: %s", cotacao.Bid)
	return os.WriteFile("cotacao.txt", []byte(content), 0644)
}
