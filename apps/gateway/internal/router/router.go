package router

import (
	"fmt"
	"salesbot-kommo/apps/gateway/internal/code"
	"salesbot-kommo/apps/gateway/internal/types"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	midLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

// Config ...
type Config struct {
	Addr string `yaml:"bind_addr" env:"BIND_ADDR" env-default:"0.0.0.0"`
	Port string `yaml:"bind_port" env:"BIND_PORT" env-default:"80"`
}

// ServiceInterface ...
type ServiceInterface interface {
}

// Router ...
type Router struct {
	config   *Config
	service  ServiceInterface
	router   *fiber.App
	validate *validator.Validate
}

// New ...
func New(config *Config, service ServiceInterface) (*Router, error) {
	r := &Router{
		config:   config,
		service:  service,
		router:   fiber.New(),
		validate: validator.New(),
	}

	r.router.Use(midLogger.New())
	r.router.Use(r.cacheControl)

	return r, nil
}

// Listen ...
func (r *Router) Listen() error {
	return r.router.Listen(fmt.Sprintf("%s:%s", r.config.Addr, r.config.Port))
}

// cacheControl ...
func (r *Router) cacheControl(ctx *fiber.Ctx) error {
	ctx.Response().Header.Add("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response().Header.Add("Pragma", "no-cache")
	ctx.Response().Header.Add("Expires", "0")
	return ctx.Next()
}

// sendResponse ...
func (r *Router) sendResponse(ctx *fiber.Ctx, result any) error {
	return ctx.JSON(fiber.Map{
		"result": result,
	})
}

// sendValidationError ...
func (r *Router) sendValidationError(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"result": map[string]any{
			"data": types.ErrorResponse{
				Code:    code.BadRequestData,
				Message: "Bad request data",
			},
		},
	})
}

// sendProcessError ...
func (r *Router) sendProcessError(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"result": map[string]any{
			"data": types.ErrorResponse{
				Code:    code.ProcessRequestError,
				Message: "Process request error",
			},
		},
	})
}

// shouldBindJSON ...
func (r *Router) shouldBindJSON(ctx *fiber.Ctx, req interface{}) error {
	if err := ctx.BodyParser(req); err != nil {
		log.Error(err)
		return err
	}

	if err := r.validate.Struct(req); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
