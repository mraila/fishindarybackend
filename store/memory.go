package store

import "fishindary/model"

type Store struct {
	Catches     []model.Catch
	Spots       []model.Spot
	CatchNextID int
	SpotNextID  int
}

func NewStore() *Store {
	return &Store{
		Catches:     []model.Catch{},
		Spots:       []model.Spot{},
		CatchNextID: 1,
		SpotNextID:  1,
	}
}
