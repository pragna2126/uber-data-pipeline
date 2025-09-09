package pkg

import(
	"analytics-layer/config"
)
type Result struct {
	RatingsBy string
    VehicleType string
    Avg float64
}

func Ratings_summary()map[string]float64{
	ratings_summary :=make(map[string]float64)
	var avgDriverRating float64

	config.DB.Model(&Ride{}).
		Select("AVG(driver_rating)").
		Scan(&avgDriverRating)
	var avgCustomerRating float64

	config.DB.Model(&Ride{}).
		Select("AVG(customer_rating)").
		Scan(&avgCustomerRating)

	ratings_summary["average driver rating"]=float64(avgDriverRating)
	ratings_summary["average customer rating"]=float64(avgCustomerRating)
	return ratings_summary
}

func RatingsByVehicleType() []Result {
	var results []Result

	var highestCustomerRated Result

	config.DB.Model(&Ride{}).
		Select("vehicle_type, AVG(customer_rating) as avg").
		Group("vehicle_type").
		Order("avg DESC").
		Limit(1).
		Scan(&highestCustomerRated)

	highestCustomerRated.RatingsBy="Customer"
	results = append(results, highestCustomerRated)

	var highestDriverRated Result

	config.DB.Model(&Ride{}).
		Select("vehicle_type, AVG(driver_rating) as avg").
		Group("vehicle_type").
		Order("avg DESC").
		Limit(1).
		Scan(&highestDriverRated)

	highestDriverRated.RatingsBy="Driver"
	results = append(results, highestDriverRated)

	return results
}


//git