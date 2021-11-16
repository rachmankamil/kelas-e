package routes

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"km-kelas-e/config"
	c_articles "km-kelas-e/controller/articles"
	c_auth "km-kelas-e/controller/auth"

	our_middleware "km-kelas-e/middleware"
)

func New() *echo.Echo {

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	our_middleware.LogMiddleware(e)

	auth := e.Group("auth")
	auth.POST("/login", c_auth.AuthLogin)

	jwtAuthGroup := e.Group("jwt_auth")
	jwtAuthGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS512.Name,
		SigningKey:    []byte(config.JwtSecret),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"status":  "error login",
				"message": e,
			})
		},
	}))

	jwtAuthGroup.GET("/articles", c_articles.GetAllArticle)

	basicAuthGroup := e.Group("basic_auth")
	basicAuthGroup.Use(middleware.BasicAuth(our_middleware.BasicAuth))

	basicAuthGroup.GET("/articles", c_articles.GetAllArticle)

	// router
	// article:= e.Group("article")
	// article.Use()
	// e.GET("/article/:id", getArticle)
	e.GET("", c_articles.GetAllArticle, middleware.BasicAuth(our_middleware.BasicAuth))
	// e.POST("/article", addArticle)
	// e.PUT("/article/:id", updateArticle)

	return e
}
