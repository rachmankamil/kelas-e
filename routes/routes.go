package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	c_articles "km-kelas-e/controller/articles"
)

func New() *echo.Echo {

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	//router
	// e.GET("/article/:id", getArticle)
	e.GET("/articles", c_articles.GetAllArticle)
	// e.POST("/article", addArticle)
	// e.PUT("/article/:id", updateArticle)

	return e
}
