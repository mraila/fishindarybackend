package model

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Spot struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	BodyOfWater string   `json:"body_of_water"`
	Visited     bool     `json:"visited"`
	Location    Location `json:"location"`
	//add picture url in future
}
