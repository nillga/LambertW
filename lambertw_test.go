package lambertw

import (
	"testing"
)

func TestW_Setup (t *testing.T) {
	tests := []struct{
		name string
		input int
		want int
	}{
		{"Zero Branch", 0, 0}, {"-1 Branch", -1, -1}, {"Conversion", -100, -1}, {"Rejection", 12, 1},
	}
	for _,test := range tests {
		w := new(W)
		t.Run(test.name, func(t *testing.T) {
			w.Setup(test.input)
			got := w.Branch

			if got != test.want {
				t.Errorf("Error: Got %d but wanted %d", got, test.want)
			}
		})
	}
}

func TestW_Router0(t *testing.T) {
	w := new(W)
	w.Setup(0)
	w.X = 3
	got := w.Router0()
	t.Logf("Result is: %.10f", got)
}