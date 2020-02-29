package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{10, 10}, 100},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{8, 4}, 16},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("want %v got %v", tt.want, got)
			}
		})
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()

// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("want %v got %v", want, got)
// 		}
// 	}
// 	t.Run("Rectangle", func(t *testing.T) {

// 		rectangle := Rectangle{10, 10}

// 		checkArea(t, rectangle, 100)
// 	})

// 	t.Run("Circle", func(t *testing.T) {
// 		circle := Circle{10}

// 		want := 314.1592653589793

// 		checkArea(t, circle, want)
// 	})
// }
