package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pgvector/pgvector-go"
	pgxvector "github.com/pgvector/pgvector-go/pgx"
	"os"
	"pgvector/Config"
	"pgvector/services/EmbeddingService"
)

func main() {
	Config.InitDb()

	e := EmbeddingService.New(Config.Db)

	// Veritabanı bağlantısını kur
	ctx := context.Background()
	cnnString, _ := Config.GetCnnString()
	conn, err := pgx.Connect(ctx, cnnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	err = pgxvector.RegisterTypes(ctx, conn)
	if err != nil {
		panic(err)
	}

	// Soru vektörünü al
	query := "restoran ne?"
	questionEmbedding, err := e.FetchEmbeddings([]string{query}, Config.OpenAIKey)
	if err != nil {
		panic(err)
	}

	// En yakın cevabı sorgula
	var closestAnswer struct {
		ID     int
		Answer string
		Score  float64
	}
	err = conn.QueryRow(ctx, `
		SELECT id, answer, answer_embedding <-> $1 AS score
		FROM answers
		ORDER BY score
		LIMIT 1
	`, pgvector.NewVector(questionEmbedding[0])).Scan(&closestAnswer.ID, &closestAnswer.Answer, &closestAnswer.Score)
	if err != nil {
		panic(err)
	}

	fmt.Printf("En yakın cevap: %s (Score: %f)\n", closestAnswer.Answer, closestAnswer.Score)
}
