package service

import (
	"time"

	"fishindary/model"
	"fishindary/store"
)

type CatchService struct {
	Store       *store.Store
	SpotService *SpotService
}

func NewCatchService(st *store.Store, ss *SpotService) *CatchService {
	return &CatchService{
		Store:       st,
		SpotService: ss,
	}
}

func (cs *CatchService) CreateCatch(c model.Catch) model.Catch {
	c.ID = cs.Store.CatchNextID
	cs.Store.CatchNextID++

	c.CatchTime = time.Now()

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
