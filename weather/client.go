package weather

import (
	"fishindary/model"
	"time"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetWeatherData(lat, long float64, catchTime time.Time) (*model.Weather, error) {
	return &model.Weather{
		Temperature:       25,
		WindSpeed:         5,
		WindDirection:     "West",
		CloudCover:        0.5,
		Precipitation:     0.1,
		AirPressure:       1013.25,
		Humidity:          70,
		AirPressureRising: false,
	}, nil
}
