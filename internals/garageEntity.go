package internals

type GarageEntity struct {
	Name      string `json:"name"`
	Current   int    `json:"current"`
	OpenSpots int    `json:"open_spots"`
	Total     int    `json:"total"`
	Percent   int    `json:"percent"`
}
