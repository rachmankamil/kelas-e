package c_articles

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"km-kelas-e/features/articles"
	presentation_response "km-kelas-e/features/articles/presentation/response"
)

type ArticlesHandler struct {
	articleBussiness articles.Bussiness
}

func NewArticleHandler(ab articles.Bussiness) *ArticlesHandler {
	return &ArticlesHandler{
		articleBussiness: ab,
	}
}

func (ah *ArticlesHandler) GetAllArticle(c echo.Context) error {
	result := ah.articleBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

// func getArticle(c echo.Context) error {

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	if id <= len(data) && id > 0 {
// 		return c.JSON(http.StatusOK, data[id-1])
// 	} else {
// 		return c.JSON(http.StatusOK, data)
// 	}
// }

// func GetAllArticle(c echo.Context) error {
// 	result, err := m_articles.SelectAll()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"claims":  middleware.ExtractClaim(c),
// 		"message": "hope all feeling well",
// 		"data":    respArticle.FromModelSlice(result),
// 	})
// }

// func addArticle(c echo.Context) error {
// 	newArticle := Article{}
// 	c.Bind(&newArticle)

// 	if err := DB.Save(&newArticle).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "hope all feeling well",
// 		"data":    newArticle,
// 	})
// }

// func updateArticle(c echo.Context) error {

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	reqArticle := Article{}
// 	c.Bind(&reqArticle)

// 	// var dataArticle Article

// 	// DB.Where("id = ?", id).First(&dataArticle)

// 	// if reqArticle.Title != "" {
// 	// 	dataArticle.Title = reqArticle.Title
// 	// }

// 	if err := DB.Where("id = ?", id).Updates(&reqArticle).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	var dataArticle Article
// 	DB.Raw("SELECT * FROM articles", id).First(&dataArticle)

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "hope all feeling well",
// 		"data":    dataArticle,
// 	})
// }
