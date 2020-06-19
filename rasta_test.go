package rasta

import "testing"

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

var triangle = []byte(`
	{
		"name": "triangle",
		"points": [
			{
				"x": 150.0,
				"y": 0.0
			},
			{
				"x": 75.0,
				"y": 200.0
			},
			{
				"x": 225.0,
				"y": 200
			}
		]
	}`,
)
