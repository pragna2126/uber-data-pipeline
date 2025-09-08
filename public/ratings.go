package routes

import (
	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/gofiber/fiber/v2"
	"analytics-layer/pkg"
)

func Ratings_summary() fiber.Handler{
	return http_server.NewHandler[any,any,any]().
	SetRouteAccess(http_server.Public).
	SetHandler(
		func (data http_server.RequestData[any, any, any]) (any,error){
			return pkg.Ratings_summary(),nil
		}).Build()
}
func Ratings() fiber.Handler{
	return http_server.NewHandler[any,any,any]().
	SetRouteAccess(http_server.Public).
	SetHandler(
		func (data http_server.RequestData[any, any, any]) (any,error){
			return pkg.RatingsByVehicleType(),nil
		}).Build()
}