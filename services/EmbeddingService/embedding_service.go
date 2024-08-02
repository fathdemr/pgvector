package EmbeddingService

import "gorm.io/gorm"

type EmbeddingService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EmbeddingService {
	return &EmbeddingService{
		db: db,
	}
}
