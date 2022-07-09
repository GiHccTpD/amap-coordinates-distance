package amap_coordinates_distance

import (
	"testing"
)

func TestCoordinates_Distance(t *testing.T) {
	// 116.368904,39.923423,116.387271,39.922501
	pointA := Coordinates{116.368904, 39.923423}
	pointB := Coordinates{116.387271, 39.922501}

	distance1 := pointA.Distance(pointB)
	t.Logf("The distance from point A to point B is %.2f meters.\n", distance1)

	// 116.373528,39.922288,116.387271,39.922501
	pointC := Coordinates{116.373528, 39.922288}
	pointD := Coordinates{116.387271, 39.922501}
	distance2 := pointC.Distance(pointD)
	t.Logf("The distance from point C to point D is %.2f meters.\n", distance2)

	pointE := Coordinates{117.129181, 36.662213}
	pointF := Coordinates{117.129181, 36.662213}
	distance3 := pointE.Distance(pointF)
	t.Logf("The distance from point E to point F is %.2f meters.\n", distance3)
}
