package rasta

import (
	"encoding/json"
)

func ReadPolygon(polygonData []byte) *Polygon {
	var polygon *Polygon
	err := json.Unmarshal(polygonData, &polygon)

	if err != nil {
		return nil
	}

	return polygon
}
