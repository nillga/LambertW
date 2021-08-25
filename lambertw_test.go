package lambertw

import (
	"testing"
)

func TestW_setup (t *testing.T) {
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
			w.setup(test.input)
			got := w.branch

			if got != test.want {
				t.Errorf("Error: Got %d but wanted %d", got, test.want)
			}
		})
	}
}

func TestW_router(t *testing.T) {
	w := new(W)
	w.setup(-1)
	w.x = -0.1
	got := w.router()
	t.Logf("Result is: %.10f", got)
}