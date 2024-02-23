package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func TestRepaet(t *testing.T) {
  t.Run("run X times", func(t *testing.T) {
    got2 := Repeat("b", 2)
    want2 := "bb"

    assetValue(t, got2, want2)

    got7 := Repeat("c", 7)
    want7 := "ccccccc"

    assetValue(t, got7, want7)
  })

  t.Run("run 0 times if count is empty", func(t *testing.T) {
    got := Repeat("a", 0)
    want := ""

    assetValue(t, got, want)
  })
}

func assetValue(t testing.TB, got, want string) {
  t.Helper()

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
