package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", handleCotacao)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := fetchCotacaoFromAPI(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar cotação: %v", err), http.StatusInternalServerError)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	err = saveCotacaoToDB(ctxDB, cotacao)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao salvar cotação no banco de dados: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cotacao)
}

func fetchCotacaoFromAPI(ctx context.Context) (*Cotacao, error) {
	const urlCotacao = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", urlCotacao, nil)
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

	var result map[string]Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	cotacao, ok := result["USDBRL"]
	if !ok {
		return nil, fmt.Errorf("cotação não encontrada na resposta")
	}

	return &cotacao, nil
}

func saveCotacaoToDB(ctx context.Context, cotacao *Cotacao) error {
	db, err := sql.Open("sqlite", "cotacoes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, bid TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)"
	if _, err := db.ExecContext(ctx, query); err != nil {
		return err
	}

	query = "INSERT INTO cotacoes (bid) VALUES (?)"
	if _, err := db.ExecContext(ctx, query, cotacao.Bid); err != nil {
		return err
	}

	return nil
}
