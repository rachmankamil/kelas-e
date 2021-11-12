package migrate

import (
	"km-kelas-e/config"
	m_articles "km-kelas-e/model/articles"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_articles.Article{})
}
