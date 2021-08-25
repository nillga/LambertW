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
		w := new(w)
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
	w := new(w)
	w.setup(-1)
	w.x = 5
	got := w.router()
	t.Logf("Result is: %.10f", got)
}

func TestPolynomial_coeff (t *testing.T) {
	t.Run("BranchPoint", func(t *testing.T) {
		p := polynomial{"branchPoint", 0}
		got := p.coeff()

		if got != -1 {
			t.Errorf("Got %f but expected -1", got)
		}
	})
	t.Run("AsymptoticB", func(t *testing.T) {
		p := polynomial{1,0}
		got := p.coeff()

		if got != 0 {
			t.Errorf("Got %f but expected 0", got)
		}
	})
}
func TestPolynomial_coeff2 (t *testing.T) {
	p := polynomial{"AsymptoticPolynomialA",0}
	got := p.coeff2(1)

	if got != -1 {
		t.Errorf("Got %f but expected -1", got)
	}
}

func TestHornerLib (t *testing.T) {
	t.Run("BranchPoint", func(t *testing.T) {
		got := hornerLib("branchPoint",0,0)

		if got != -1 {
			t.Errorf("Got %f but expected -1", got)
		}
	})
	t.Run("AsymptoticA", func(t *testing.T) {
		got := hornerLib("AsymptoticPolynomialA",0,1)

		if got != -1 {
			t.Errorf("Got %f but expected -1", got)
		}
	})
	t.Run("AsymptoticB", func(t *testing.T) {
		got := hornerLib(1,0,0)

		if got != 0 {
			t.Errorf("Got %f but expected 0", got)
		}
	})
}

func TestHalleyStep (t *testing.T) {
	got := halleyStep(1,0)
	
	if got != 0.5 {
		t.Errorf("Got %f but wanted 0.5", got)
	}
}