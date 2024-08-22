package integers

import "testing"

// func TestAdd(t *testing.T) {
// 	result := Add(4, 3)
// 	expected := 7

// 	if result != expected {
// 		t.Errorf("Expected '%d' but got '%d'", expected, result)
// 	}
// }

func TestTryAdd(t *testing.T) {
	outcome := Add(21, 14)
	want := 35
	
	if outcome != want {
		t.Errorf("Expected '%d' but got '%d'", outcome, want)
	}
}