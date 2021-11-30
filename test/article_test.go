package test

import (
	"fmt"
	"km-kelas-e/config"
	m_articles "km-kelas-e/model/articles"
	"net/http"
	"testing"
)

func init() {
	config.InitDB("kelas_e_test")

	article := &m_articles.Article{
		Title: "Test1",
	}
	if err := config.DB.Create(&article).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetArticle(t *testing.T) {
	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)
	
	result := e.GET("/articles").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("hope")
	result.Value("data").Array().NotEmpty()
}
