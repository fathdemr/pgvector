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
	e := EmbeddingService.New(Config.Db)
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

	// alternative
	//var answerNames []string
	//Config.Db.Model(&Models.Answer{}).Select("answer").Scan(&answerNames)
	//Config.Db.Model(&Models.Answer{}).Pluck("answer", &answerNames)
	// select answer from answers

	var answersStringArray []string
	for _, answer := range answers {
		answersStringArray = append(answersStringArray, answer.Answer)
	}

	embeddings, err := e.FetchEmbeddings(answersStringArray, Config.OpenAIKey)
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
