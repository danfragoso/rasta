package rasta

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestReadPolygon(t *testing.T) {
	polygon := ReadPolygon(triangle)

	if polygon == nil {
		t.Errorf("Error reading polygon data")
	} else {
		pointsLen := len(polygon.Points)
		if pointsLen != 3 {
			t.Errorf("Expected 3 points, got %d", pointsLen)
		}

		if polygon.Name != "triangle" {
			t.Errorf("Expected \"triangle\", got \"" + polygon.Name + "\"")
		}
	}

}

func TestRasterize(t *testing.T) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{300, 300},
	})

	polygon := ReadPolygon(triangle)
	Rasterize(polygon, img, image.Point{0, 0})

	f, _ := os.Create("out.png")
	png.Encode(f, img)
}

var triangle = []byte(`
	{
		"name": "triangle",
		"points": [
			{
				"x": 150,
				"y": 0
			},
			{
				"x": 75,
				"y": 200
			},
			{
				"x": 225,
				"y": 200
			}
		]
	}`,
)
