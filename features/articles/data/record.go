package data

import (
	"fmt"
	"km-kelas-e/features/articles"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Status    bool
	Title     string `gorm:"column:article_title"`
	Content   string `gorm:"column:raw_content"`
	Publisher string `gorm:"column:pub"`
	// CategoryID uint
	// Category   CategoryArticle `gorm:"foreignKey:category_id"`
	// TagArticle []TagArticle    `gorm:"many2many:articles_tag"`
}

type CategoryArticle struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50)"`
	Status bool
}

type TagArticle struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50)"`
	Status bool
}

//DTO

func (a *Article) toCore() articles.Core {
	fmt.Printf("%+v\n\n", a)
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

func fromCore(core articles.Core) Article {
	return Article{}
}
