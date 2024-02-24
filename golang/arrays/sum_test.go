package arrays

import (
	"testing"
  "reflect"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("want %d, got %d | %v", want, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
  got := SumAll([]int{1, 2}, []int{0, 9})
  want := []int{3, 9}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("want %v, got %v", want, got)
  }
}

func TestSumAllTails(t *testing.T) {
  got := SumAllTails([]int{1, 2}, []int{0, 9})
  want := []int{2, 9}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("want %v, got %v", want, got)
  }
}
