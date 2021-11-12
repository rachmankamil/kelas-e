package m_articles

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Status  bool   `json:"status"`
	Title   string `json:"title" gorm:"column:article_title"`
	Content string `json:"content" gorm:"column:raw_content"`
}
