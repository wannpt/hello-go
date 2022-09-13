package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if(got != expected) {
		t.Errorf("got %f expected %f", got, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72},
		{name: "Circle",shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle",shape: Triangle{12, 6}, want: 36.0},
  }

		for _, tt := range areaTests {
			got := tt.shape.Area()
				if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}