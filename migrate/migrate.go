package migrate

import (
	"km-kelas-e/config"
	m_articles "km-kelas-e/features/articles/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_articles.Article{})
}
