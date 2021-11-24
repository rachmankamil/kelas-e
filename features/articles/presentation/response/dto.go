package response

import (
	"time"

	"km-kelas-e/features/articles"
)

type Article struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func FromCore(core articles.Core) Article {
	return Article{
		ID:        core.ID,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Title:     core.Title,
		Content:   core.Content,
	}
}

func FromCoreSlice(core []articles.Core) []Article {
	var artArray []Article
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
