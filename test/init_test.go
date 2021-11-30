package test

import (
	"km-kelas-e/config"
	"km-kelas-e/migrate"
	"km-kelas-e/routes"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	config.InitDB("kelas_e_test")
	migrate.AutoMigrate()

	echoHandler = routes.New()
	server = httptest.NewServer(echoHandler)

}

func TearUp() func() {
	return func() {
		config.InitDB("kelas_e_test")
		config.DB.Exec("TRUNCATE TABLE articles;")
	}
}

// GetHttpExpect Get cached expect, create new if nil
func GetHTTPExpect(t *testing.T) *httpexpect.Expect {
	if server == nil {
		server = httptest.NewServer(echoHandler)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpect(t *testing.T) *httpexpect.Expect {
	// TODO : printer set, only if the testing is failed
	// https://github.com/gavv/httpexpect/issues/69
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
