package internal

type Trip struct {
	UUID                string  `json:"uuid" ddb:"key=hash"`
	VendorID            int     `json:"VendorID"`
	TpepPickupDatetime  string  `json:"tpep_pickup_datetime"`
	TpepDropoffDatetime string  `json:"tpep_dropoff_datetime"`
	PassengerCount      int     `json:"passenger_count"`
	TripDistance        float64 `json:"trip_distance"`
	TotalAmount         float64 `json:"total_amount"`
}
