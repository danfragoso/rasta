package rasta

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadPolygon(t *testing.T) {
	triangle, _ := ioutil.ReadFile("assets/triangle.json")
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
	star, _ := ioutil.ReadFile("assets/star.json")
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{300, 300},
	})

	polygon := ReadPolygon(star)
	Rasterize(polygon, img, image.Point{0, 0})

	f, _ := os.Create("out.png")
	png.Encode(f, img)
}
