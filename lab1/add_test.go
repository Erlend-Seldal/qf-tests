package lab1

import "testing"

var addTests = []struct{ a, b, want int }{
	{1, 2, 3},
	{0, 0, 0},
	{-1, 1, 0},
	{10, 5, 15},
	{-3, -4, -7},
}

func init() {
	// One point per test case, weight 1.
	scores.Add(TestAdd, len(addTests), 1)
}

func TestAdd(t *testing.T) {
	sc := scores.Max(t)
	defer sc.Recover()
	for _, tt := range addTests {
		if got := Add(tt.a, tt.b); got != tt.want {
			sc.Dec()
			sc.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}