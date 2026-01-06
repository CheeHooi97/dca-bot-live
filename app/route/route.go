package route

import (
	"dca-bot-live/app/handler"
	"dca-bot-live/app/service"
	"dca-bot-live/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mileusna/useragent"
)

func NewRouter() *echo.Echo {
	server := echo.New()
	server.Validator = utils.NewValidator()
	server.Use(
		middleware.Recover(),
		middleware.RequestLogger(),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				agent := c.Request().UserAgent()
				ua := useragent.Parse(agent)
				if ua.Bot {
					return c.NoContent(http.StatusNoContent)
				}
				return next(c)
			}
		},
	)

	return server
}

func SetupRoutes(server *echo.Echo, services *service.Services) *echo.Echo {
	// e.Validator = utils.NewValidator()

	server.Use(
		middleware.Recover(),
		middleware.Logger(),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				agent := c.Request().UserAgent()
				ua := useragent.Parse(agent)
				if ua.Bot {
					return c.NoContent(http.StatusNoContent)
				}
				return next(c)
			}
		},
	)

	server.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Response().Header().Set("X-Content-Type-Options", "nosniff")
				c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000")
				c.Response().Header().Set("X-Frame-Options", "DENY")
				c.Response().Header().Set("Cache-control", "no-store")
				c.Response().Header().Set("Pragma", "no-store")
				return next(c)
			}
		},
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			MaxAge:           86400,
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
			ExposeHeaders: []string{
				"x-request-id",
				"content-type",
				"vary",
			},
		}),
	)

	h := handler.NewHandler(services)

	trade := server.Group("/trade")
	trade.POST("/create", h.CreateBot)

	return server
}
