package routes

import (
	"analytics-layer/pkg"

	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/gofiber/fiber/v2"
)

func Revenue() fiber.Handler{
	return http_server.NewHandler[any,any,any]().
	SetRouteAccess(http_server.Public).
	SetHandler(
		func (data http_server.RequestData[any, any, any]) (any, error){
			return pkg.Revenue_distribution(),nil

		}).Build()
	
}


