package model

type Weather struct {
	Temperature       float64 `json:"temperature"`
	AirPressure       float64 `json:"air_pressure"`
	AirPressureRising bool    `json:"air_pressure_rising"`
	Humidity          float64 `json:"humidity"`
	WindSpeed         float64 `json:"wind_speed"`
	WindDirection     string  `json:"wind_direction"`
	CloudCover        float64 `json:"cloud_cover"`
	Precipitation     float64 `json:"precipitation"`
}

//find nearest water hydromolygy station and get data from it if applicable
