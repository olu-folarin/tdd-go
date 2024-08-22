package repeat

import "testing"

func TestRepeat(t *testing.T) {
	got, _ := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


func TestAdder(t *testing.T) {
	expected := Adder("b")
	want := "bbbbbbb"

	if expected != want {
		t.Errorf("got %q want %q", expected, want)
	}
}