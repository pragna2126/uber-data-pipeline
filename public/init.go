package routes

import (

	"github.com/HOA-Share/hoashare-go-common/http_server"
	
)

func MountRoutes(app *http_server.Router){
	revenue:=app.Group("/revenue")
	revenue.Get("/distribution",Revenue())

	cancellations:=app.Group("/cancellations")
	cancellations.Get("/customer",Customer_reason())
	cancellations.Get("/top-cancellations",top_3_reasons())

	ratings:=app.Group("/ratings")
	ratings.Get("/summary",Ratings_summary())
	ratings.Get("/vehicle-type",Ratings())
}