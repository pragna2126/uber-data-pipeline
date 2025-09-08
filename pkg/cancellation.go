package pkg

import(
	"analytics-layer/config"
)
type Reason struct{
	Reason string
}
func Customer_reason( ) []Reason {
	var customer_reason []Reason
						
		config.DB.Model(&Ride{}).
			Select("cancellation_reason as reason, COUNT(*) as count").
			Group("cancellation_reason").Scan(&customer_reason)
		return customer_reason
}

func Top_3_reasons()[]Reason{
	var top_cancellations []Reason 
	 
		
	config.DB.Model(&Ride{}).
		Select("cancellation_reason as reason, COUNT(*) as count").
		Where("cancellation_reason IS NOT NULL").
		Group("cancellation_reason").
		Order("count DESC").
		Limit(3).
		Scan(&top_cancellations)
	return top_cancellations
}