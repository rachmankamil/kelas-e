package request

import (
	"km-kelas-e/features/articles"
)

type Article struct {
	Status  bool   `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToCore(req Article) articles.Core {
	return articles.Core{
		Status:  req.Status,
		Title:   req.Title,
		Content: req.Content,
	}
}
