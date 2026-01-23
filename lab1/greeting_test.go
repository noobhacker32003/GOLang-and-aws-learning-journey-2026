package greeting

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("World", "normal")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func TestGreetWithName(t *testing.T) {
	got := Greet("Karim", "normal")
	want := "Hello, Karim!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func TestGreetWithEmptyName(t *testing.T) {
	got := Greet("", "normal")
	want := "Hello, Poridhian!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func TestGreetCasual(t *testing.T) {
	got := Greet("Karim", "casual")
	want := "Hey, Karim!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
func TestGreetFormal(t *testing.T) {
	got := Greet("Rahim uncle", "formal")
	want := "Good day, Rahim uncle!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
