package migrate

import (
	"km-kelas-e/config"
	m_articles "km-kelas-e/features/articles/data"
	m_user "km-kelas-e/features/user/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_articles.Article{},
		&m_user.User{},
		&m_articles.CategoryArticle{},
	)
}
