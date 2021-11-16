package data

import (
	"km-kelas-e/features/articles"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Status  bool
	Title   string `gorm:"column:article_title"`
	Content string `gorm:"column:raw_content"`
}

//DTO

func (a *Article) toCore() articles.Core {
	return articles.Core{
		ID:        int(a.ID),
		Title:     a.Title,
		Status:    a.Status,
		Content:   a.Content,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []Article) []articles.Core {
	a := []articles.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}
