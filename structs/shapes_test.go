package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// These test examples use a helper function
	// checkArea := func(t *testing.T, s Shape, want float64) {
	// 	t.Helper()
	// 	got := s.Area()

	// 	if got != want {
	// 		t.Errorf("got %g, wanted %g", got, want)
	// 	}
	// }

	// t.Run("rectangle", func(t *testing.T) {
	// 	rect := Rectangle{10.0, 10.0}
	// 	want := 100.0
	// 	checkArea(t, rect, want)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	// These tests use a table-driven approach
	// i.e. an array of test cases looped through
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 10.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				// %#v will print out the struct with its fields and values
				// so if it fails we get detail
				t.Errorf("%#v got %g, wanted %g", areaTest.shape, got, areaTest.want)
			}
		})
	}
}
