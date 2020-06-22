package rasta

import (
	"encoding/json"
	"image"
)

func ReadPolygon(polygonData []byte) *Polygon {
	var polygon *Polygon
	err := json.Unmarshal(polygonData, &polygon)

	if err != nil {
		return nil
	}

	return polygon
}

func Rasterize(polygon *Polygon, img *image.RGBA, coords image.Point) {
	//Stroke
	for idx, point := range polygon.Points {
		if idx == len(polygon.Points)-1 {
			firstPoint := polygon.Points[0]
			PlotLine(img, point.X, point.Y, firstPoint.X, firstPoint.Y)
		} else {
			nextPoint := polygon.Points[idx+1]
			PlotLine(img, point.X, point.Y, nextPoint.X, nextPoint.Y)
		}
	}

}
