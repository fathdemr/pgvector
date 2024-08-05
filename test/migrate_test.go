package test

import (
	"pgvector/Config"
	"pgvector/Models"
	"testing"
)

func TestMigrate(t *testing.T) {

	Config.InitDb()

	answer := []Models.Answer{
		{ID: 1, Answer: "Otelin Adı Hilton'dur"},
		{ID: 2, Answer: "Otelde Burger King, Nusret, Starbucks restronları mevcut"},
		{ID: 3, Answer: "Bu otelin şehri Istanbul'dur"},
		{ID: 4, Answer: "Bu otelin ülkesi Türkiye'dir"},
		{ID: 5, Answer: "Otelin yetkili kişisi Fatih'dir"},
		{ID: 6, Answer: "Otelin yetkili telefonu 05555555555"},
		{ID: 7, Answer: "Otelin e postası otel@example.com"},
		{ID: 8, Answer: "Otelin viber numarası 05555555555"},
		{ID: 9, Answer: "Sosyal medya hesaplarımız otel_dalaman"},
		{ID: 10, Answer: "Resepsiyon numaramız 02121111111"},
	}

	Config.Db.AutoMigrate(&Models.Answer{})

	err := Config.Db.Create(&answer).Error
	if err != nil {
		panic(err)
	}

}
