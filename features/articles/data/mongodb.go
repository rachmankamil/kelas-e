package data

import (
	"km-kelas-e/features/articles"

	"gorm.io/gorm"
)

type mongoArticleRepository struct {
	Conn *gorm.DB
}

func NewMongoArticleRepository(conn *gorm.DB) articles.Data {
	return &mongoArticleRepository{
		Conn: conn,
	}
}

func (ar *mongoArticleRepository) InsertData(data articles.Core) (resp articles.Core, err error) {
	return articles.Core{}, nil
}

func (ar *mongoArticleRepository) SelectData(title string) (resp []articles.Core) {
	record := []Article{}
	if err := ar.Conn.Find(&record).Error; err != nil {
		return []articles.Core{}
	}
	return toCoreList(record)
}
