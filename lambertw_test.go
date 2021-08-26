package lambertw

import (
	"math"
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
		t.Run("order == 0 / 1", func(t *testing.T) {
			got := hornerLib("AsymptoticPolynomialA",0,1)

			if got != -1 {
				t.Errorf("Got %f but expected -1", got)
			}
		})
		t.Run("order == 2 ... 5", func(t *testing.T) {
			got := hornerLib("AsymptoticPolynomialA",2,1)
			h := horner{2,2}
			want := h.eval(1)

			if got != want {
				t.Errorf("Got %f but expected %f", got, want)
			}
		})
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

func TestI_recurse (t *testing.T) {
	t.Run("Depth == 0", func(t *testing.T) {
		i := iterator{halleyStep,0}

		got := i.recurse(1,0)
		if got != 0 {
			t.Errorf("Got %f but wanted 1", got)
		}
	})
	t.Run("Depth == 1", func(t *testing.T) {
		i := iterator{halleyStep,1}

		got := i.recurse(1,0)
		if got != 0.5 {
			t.Errorf("Got %f but wanted 0.5", got)
		}
	})
	t.Run("Depth == 2", func(t *testing.T) {
		i := iterator{halleyStep,2}

		got := i.recurse(1,0)
		want := halleyStep(1,0.5)
		if got != want {
			t.Errorf("Got %f, but wanted %f", got, want)
		}
	})
}

func TestBranch_branchPointExpansion (t *testing.T) {
	t.Run("Upper Branch", func(t *testing.T) {
		b := branch{branch: 0, order: 1, sgn: 1}
		h := horner{"branchPoint", 1}
		got := b.branchPointExpansion(0)
		want := h.eval(math.Sqrt(2))

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
	t.Run("Lower Branch", func(t *testing.T) {
		b := branch{branch: -1, order: 1, sgn: -1}
		h := horner{"branchPoint", 1}
		got := b.branchPointExpansion(0)
		want := h.eval(-math.Sqrt(2))

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
}
func TestBranch_asymptoticExpansion (t *testing.T) {
	t.Run("Upper Branch", func(t *testing.T) {
		b := branch{branch: 0, order: 1, sgn: 1}
		got := b.asymptoticExpansion(math.E)
		want := asymptoticExpansionImpl(1, 0, 1)

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
	t.Run("Lower Branch", func(t *testing.T) {
		b := branch{branch: -1, order: 1, sgn: -1}
		got := b.asymptoticExpansion(-math.Exp(-1))
		want := asymptoticExpansionImpl(-1,0,1)

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
}

func TestAsymptoticExpansionImpl (t *testing.T) {
	got := asymptoticExpansionImpl(1,0,1)

	h := horner{"AsymptoticPolynomialA",1}
	want := h.eval2(1,0) + 1

	if got != want {
		t.Errorf("Got %f but wanted %f", got, want)
	}
}

func TestBranch_logRecursion (t *testing.T) {
	b := branch{sgn:1, order: 1}
	got := b.logRecursion(math.E)

	if got != 1 {
		t.Errorf("Got %f but wanted 1", got)
	}
}

func TestLogRecursionImpl_step (t *testing.T) {
	t.Run("order == 0 --> stop case", func(t *testing.T) {
		l := logRecursionImpl{order: 0}
		got := l.step(0)

		if got != 0 {
			t.Errorf("Got %f but wanted 0", got)
		}
	})
	t.Run("order == 1 --> recursion", func(t *testing.T) {
		l := logRecursionImpl{order: 1, sgn: 1}
		got := l.step(1)

		if got != 1 {
			t.Errorf("Got %f but wanted 0", got)
		}
	})
}