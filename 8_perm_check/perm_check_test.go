package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input []int
    output int
	}{
    {[]int{1}, 1},
    {[]int{3}, 0},
    {[]int{1,2,4}, 0},
    {[]int{1,2,3}, 1},
    {[]int{1,2,2}, 0},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
