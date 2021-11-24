package bussiness

import (
	"km-kelas-e/features/articles"

	"github.com/go-playground/validator/v10"
)

type articlesUsecase struct {
	articleData articles.Data
	validate    *validator.Validate
}

func NewAricleBussiness(artData articles.Data) articles.Bussiness {

	return &articlesUsecase{
		articleData: artData,
		validate:    validator.New(),
	}
}

func (au *articlesUsecase) CreateData(data articles.Core) (resp articles.Core, err error) {
	if err := au.validate.Struct(data); err != nil {
		return articles.Core{}, err
	}

	resp, err = au.articleData.InsertData(data)
	if err != nil {
		return articles.Core{}, err
	}
	return articles.Core{}, nil
}

func (au *articlesUsecase) GetAllData(search string) (resp []articles.Core) {
	resp = au.articleData.SelectData(search)
	return
}
