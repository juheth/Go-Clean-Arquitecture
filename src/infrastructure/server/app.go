package infraestructure

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	config "github.com/juheth/Go-Clean-Arquitecture/src/common/config"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func setRoutesByModule(app *fiber.App, h *types.HandlersStore) {
	log.Infof("Total modules registered: %d", len(h.Handlers))
	for i, module := range h.Handlers {
		log.Infof("Module %d: Prefix '%s' with %d routes", i+1, module.Prefix, len(module.Routes))
	}
	for _, handlerModule := range h.Handlers {
		route := app.Group("/" + handlerModule.Prefix)
		for _, routeItem := range handlerModule.Routes {
			log.Infof("%v %v%v", routeItem.Method, handlerModule.Prefix, routeItem.Route)
			if routeItem.Method == http.MethodGet {
				route.Get(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodPost {
				route.Post(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodPut {
				route.Put(routeItem.Route, routeItem.Handler)
			}
			if routeItem.Method == http.MethodDelete {
				route.Delete(routeItem.Route, routeItem.Handler)
			}
		}
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"isError": true,
		"message": err.Error(),
	})
}

func NewHttpFiberServer(lc fx.Lifecycle, h *types.HandlersStore, cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	setRoutesByModule(app, h)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port " + cfg.App.Port)
			go app.Listen(":" + cfg.App.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
