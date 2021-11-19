package factory

import (
	"km-kelas-e/config"
	_article_bussiness "km-kelas-e/features/articles/bussiness"
	_article_data "km-kelas-e/features/articles/data"
	_article_presentation "km-kelas-e/features/articles/presentation"
)

type Presenter struct {
	ArticlePresentation *_article_presentation.ArticlesHandler
}

func Init() Presenter {

	articleData := _article_data.NewArticleRepository(config.DB)
	articleBussiness := _article_bussiness.NewAricleBussiness(articleData)
	articlePresentation := _article_presentation.NewArticleHandler(articleBussiness)

	return Presenter{
		ArticlePresentation: articlePresentation,
	}
}
