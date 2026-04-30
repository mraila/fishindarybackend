package service

import (
	"time"

	"fishindary/model"
	"fishindary/store"
)

func CreateCatch(c model.Catch) model.Catch {
	c.ID = store.NextID
	store.NextID++

	c.CatchTime = time.Now()

	store.Catches = append(store.Catches, c)

	return c
}

func GetCatches() []model.Catch {
	return store.Catches
}
