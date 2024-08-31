package rectangle

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{7.7, 12.5}, 96.25},
		{Circle{10}, 314.1592653589793},
	}

	// iterate over the test cases
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
// func TestArea(t *testing.T) {
// 	t.Run("test area of a rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{7.7, 12.5}
// 		got := rectangle.Area()
// 		want := 96.25
	
// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("test the area of a circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})
// }

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.7, 5.8}
	got := rectangle.Perimeter()
	want := 33.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}