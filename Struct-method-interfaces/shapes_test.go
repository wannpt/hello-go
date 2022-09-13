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

	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area();
		expected := 100.0
	
		if(got != expected) {
			t.Errorf("got %g expected %g", got, expected)
		}
	})

	t.Run("Circle Area", func(t *testing.T){
		circle := Circle{10}
		got := circle.Area();
		expected := 314.1592653589793

		if(got != expected) {
			t.Errorf("got %g expected %g", got, expected)
		}
	})

}