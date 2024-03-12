package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Kacper")

  got := buffer.String()
  want := "Hello Kacper"

  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
