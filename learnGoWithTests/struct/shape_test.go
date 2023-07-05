package _struct

import "testing"

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{2, 3}, 6},
		{"triangle", Triangle{2, 3}, 5},
	}

	for _, area := range areaTest {
		got := area.shape.Area()
		t.Run(area.name, func(t *testing.T) {
			if got != area.want {
				t.Errorf("got %v want %v", got, area.want)
			}
		})
	}
}
