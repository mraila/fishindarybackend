package service

import (
	"time"

	"fishindary/model"
	"fishindary/store"
	"fishindary/weather"
)

type CatchService struct {
	Store         *store.Store
	SpotService   *SpotService
	WeatherClient *weather.Client
}

func NewCatchService(st *store.Store, ss *SpotService, wc *weather.Client) *CatchService {
	return &CatchService{
		Store:         st,
		SpotService:   ss,
		WeatherClient: wc,
	}
}

func (cs *CatchService) CreateCatch(c model.Catch) model.Catch {
	c.ID = cs.Store.CatchNextID
	cs.Store.CatchNextID++

	c.CatchTime = time.Now()

	weatherData, err := cs.WeatherClient.GetWeatherData(
		c.Location.Latitude,
		c.Location.Longitude,
		c.CatchTime,
	)

	if err == nil {
		c.Weather = weatherData
	}

	cs.Store.Catches = append(cs.Store.Catches, c)

	return c
}

func (cs *CatchService) GetCatches() []model.Catch {
	return cs.Store.Catches
}

func (cs *CatchService) GetCatchesBySpotID(spotID int) []model.Catch {
	var result []model.Catch

	for _, catch := range cs.Store.Catches {
		if catch.SpotID != nil && *catch.SpotID == spotID {
			result = append(result, catch)
		}
	}
	return result
}
