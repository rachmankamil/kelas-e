package response

import (
	"time"

	m_articles "km-kelas-e/model/articles"
)

type Article struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func FromModel(model m_articles.Article) Article {
	return Article{
		ID:        int(model.ID),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Status:    model.Status,
		Title:     model.Title,
		Content:   model.Content,
	}
}

func FromModelSlice(model []m_articles.Article) []Article {
	var artArray []Article
	for key := range model {
		artArray = append(artArray, FromModel(model[key]))
	}
	return artArray
}
