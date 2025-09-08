package pkg

import(
	"analytics-layer/config"
)
type Ride struct{
	Timestamp          string  `json:"timestamp"`
	BookingStatus      string  `json:"booking_status"`
	VehicleType        string  `json:"vehicle_type"`
	CancellationReason string  `json:"cancellation_reason"`
	DriverRating       float64 `json:"driver_rating"`
	CustomerRating     float64 `json:"customer_rating"`
	PaymentMethod      string  `json:"payment_method"`
}

type PaymentDistribution struct {
    PaymentMethod string
    Count         int

}
func Revenue_distribution() map[string]int {
    var results []PaymentDistribution
    paymentMethods := make(map[string]int)
    config.DB.Model(&Ride{}).
        Select("payment_method, COUNT(*) as count").
        Group("payment_method").
        Scan(&results)
    for _, r := range results {
        paymentMethods[r.PaymentMethod] = r.Count
    }

    return paymentMethods
}

