package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	//router
	e.GET("/article/:id", getArticle)
	e.POST("/article", addArticle)

	//starting the server
	e.Start(":1234")
}

type article struct {
	ID      int    `json:"id"`
	Status  bool   `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var data = []article{
	{1, true, "Article 1", "Article 1 pro content"},
	{2, true, "Article 2", "Article 2 pro content"},
}

func getArticle(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	if id <= len(data) && id > 0 {
		return c.JSON(http.StatusOK, data[id-1])
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func addArticle(c echo.Context) error {
	newArticle := article{}
	c.Bind(&newArticle)

	data = append(data, newArticle)

	return c.JSON(http.StatusAccepted, newArticle)
}
