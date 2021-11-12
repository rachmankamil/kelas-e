package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {

	var err error
	db, err := gorm.Open(mysql.Open("root:masukaja@/kelas-e?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&Article{})
}

func main() {
	//initiateDB
	initDB()

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	//router
	// e.GET("/article/:id", getArticle)
	e.GET("/articles", GetAllArticle)
	e.POST("/article", addArticle)
	e.PUT("/article/:id", updateArticle)

	//starting the server
	e.Start(":1234")
}

type Article struct {
	gorm.Model
	Status  bool   `json:"status"`
	Title   string `json:"title" gorm:"column:article_title"`
	Content string `json:"content" gorm:"column:raw_content"`
}

// func getArticle(c echo.Context) error {

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	if id <= len(data) && id > 0 {
// 		return c.JSON(http.StatusOK, data[id-1])
// 	} else {
// 		return c.JSON(http.StatusOK, data)
// 	}
// }

func GetAllArticle(c echo.Context) error {
	var articles []Article

	if err := DB.Find(&articles).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    articles,
	})
}

func addArticle(c echo.Context) error {
	newArticle := Article{}
	c.Bind(&newArticle)

	if err := DB.Save(&newArticle).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    newArticle,
	})
}

func updateArticle(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	reqArticle := Article{}
	c.Bind(&reqArticle)

	// var dataArticle Article

	// DB.Where("id = ?", id).First(&dataArticle)

	// if reqArticle.Title != "" {
	// 	dataArticle.Title = reqArticle.Title
	// }

	if err := DB.Where("id = ?", id).Updates(&reqArticle).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var dataArticle Article
	DB.Raw("SELECT * FROM articles", id).First(&dataArticle)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    dataArticle,
	})

}
