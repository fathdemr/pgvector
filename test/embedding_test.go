package test

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pgvector/pgvector-go"
	pgxvector "github.com/pgvector/pgvector-go/pgx"
	"pgvector/Config"
	"pgvector/Models"
	"pgvector/services/EmbeddingService"
	"testing"
)

func TestEmbedding(t *testing.T) {

	Config.InitDb()
	ctx := context.Background()
	cnnstring, _ := Config.GetCnnString()

	conn, err := pgx.Connect(ctx, cnnstring)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	err = pgxvector.RegisterTypes(ctx, conn)
	if err != nil {
		panic(err)
	}

	var answers []Models.Answer
	err = Config.Db.Find(&answers).Error

	var answersStringArray []string
	for _, answer := range answers {
		answersStringArray = append(answersStringArray, answer.Answer)
	}

	embeddings, err := EmbeddingService.FetchEmbeddings(answersStringArray, Config.OpenAIKey)
	if err != nil {
		panic(err)
	}

	for i, answer := range answers {
		_, err = conn.Exec(ctx, "UPDATE answers SET answer_embedding = $1 WHERE id = $2", pgvector.NewVector(embeddings[i]), answer.ID)
		if err != nil {
			panic(err)
		}
	}

}
