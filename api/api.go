package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// New configures application resources and routes.
func New(enableCORS bool) (*echo.Echo, error) {
	//logger := logging.NewLogger()

	e := echo.New()

	if enableCORS {
		e.Use(middleware.CORSWithConfig(corsConfig()))
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e, nil
}

func corsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
}
