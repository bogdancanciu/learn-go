package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name:"Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name:"Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name:"Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T){
			got := tc.shape.Area()
			if got != tc.hasArea {
				t.Errorf("got %g want %g", got, tc.hasArea)
			}
		})
	}
}
