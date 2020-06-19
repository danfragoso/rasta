package rasta

// Polygon - Rasta Polygon
type Polygon struct {
	Name   string   `json:"name"`
	Points []*Point `json:"points"`
}
