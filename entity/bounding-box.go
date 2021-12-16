package entity

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BoundingBox struct {
	NorthEast Coordinate `json:"northeast"`
	NorthWest Coordinate `json:"northwest"`
	SouthEast Coordinate `json:"southeast"`
	SouthWest Coordinate `json:"southwest"`
}
