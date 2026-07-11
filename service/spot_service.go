package service

import (
	"fishindary/model"
	"fishindary/store"
)

type SpotService struct {
	Store *store.Store
}

func NewSpotService(s *store.Store) *SpotService {
	return &SpotService{Store: s}
}

func (ss *SpotService) CreateSpot(s model.Spot) model.Spot {
	s.ID = ss.Store.SpotNextID
	ss.Store.SpotNextID++

	ss.Store.Spots = append(ss.Store.Spots, s)

	return s
}

func (ss *SpotService) GetSpots() []model.Spot {
	return ss.Store.Spots
}

func (ss *SpotService) SpotExists(id int) bool {
	for _, spot := range ss.Store.Spots {
		if spot.ID == id {
			return true
		}
	}
	return false
}
