package go_say_hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, World!"
	if got := SayHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestSayName(t *testing.T) {
	want := "Hello, John!"
	if got := SayName("John"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
