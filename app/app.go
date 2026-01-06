package app

import (
	"dca-bot-live/app/config"
	"dca-bot-live/app/migrate"
	"dca-bot-live/app/repository"
	"dca-bot-live/app/route"
	service "dca-bot-live/app/service"
	"dca-bot-live/app/utils"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Start(portNo string) {
	hosts := make(map[string]*Host)

	// Init
	config.LoadConfig()
	utils.NewLang()

	// Init DB
	db, err := migrate.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Migrate
	if err := migrate.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// Init repo and services
	repos := repository.InitializeRepository(db)
	services := service.InitializeService(repos)

	e := echo.New()
	e.Use(middleware.Recover())

	// route
	api := route.NewRouter()
	route.SetupRoutes(api, services)

	hosts[clearProtocol(config.GatewayApi)] = &Host{api}
	if config.IsLocal() {
		hosts["localhost:2001"] = &Host{api}
	}

	server := echo.New()
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		if !c.Response().Committed {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": map[string]interface{}{
					"code":    "INTERNAL_ERROR",
					"message": "Internal error",
					"debug":   err.Error(),
				},
			})
		}
	}

	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}
		return
	})

	server.Logger.Fatal(server.Start(":" + portNo))
}

func clearProtocol(host string) string {
	host = strings.ReplaceAll(host, "https://", "")
	host = strings.ReplaceAll(host, "http://", "")
	return host
}
