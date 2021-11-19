package data

import (
	"km-kelas-e/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	Conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) articles.Data {
	return &mysqlArticleRepository{
		Conn: conn,
	}
}

func (ar *mysqlArticleRepository) InsertData(data articles.Core) (resp articles.Core, err error) {
	return articles.Core{}, nil
}

func (ar *mysqlArticleRepository) SelectData(title string) (resp []articles.Core) {
	var record []Article
	if err := ar.Conn.Preload("TagArticle").Find(&record).Error; err != nil {
		return []articles.Core{}
	}
	return toCoreList(record)
}
