package lambertw

import (
	"math"
	"testing"
)

func TestFukushima(t *testing.T) {
	_ = math.E
	branch, x := -1, -0.000001

	t.Errorf("\nFukushima: %.16f\nVeberic: %.16f",Fukushima(branch,x),W(branch,x))
}