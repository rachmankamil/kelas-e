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
	record := fromCore(data)

	if err := ar.Conn.Create(&record).Error; err != nil {
		return articles.Core{}, err
	}

	return data, nil
}

func (ar *mysqlArticleRepository) SelectData(title string) (resp []articles.Core) {
	var record []Article
	if err := ar.Conn.Preload("TagArticle").Find(&record).Error; err != nil {
		return []articles.Core{}
	}
	return toCoreList(record)
}
