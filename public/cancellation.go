package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HOA-Share/hoashare-go-common/http_server"
	"analytics-layer/pkg"

)

func Customer_reason() fiber.Handler{
	return http_server.NewHandler[any,any,any]().
	SetRouteAccess(http_server.Public).
	SetHandler(
		func (data http_server.RequestData[any, any, any]) (any,error){
			
			return pkg.Customer_reason(),nil
		}).Build()
}

func top_3_reasons() fiber.Handler{
	return http_server.NewHandler[any,any,any]().
	SetRouteAccess(http_server.Public).
	SetHandler(
		func (data http_server.RequestData[any, any, any]) (any,error){
			
			return pkg.Top_3_reasons(),nil
		}).Build()
}