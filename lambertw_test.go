package lambertw

import (
	"fmt"
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

func TestPade_approximation (t *testing.T) {
	branch := 0
	x := 1.0
	wants := []float64{
		horner4(1, 0.07066247420543414, 2.4326814530577687, 6.39672835731526, 4.663365025836821, 0.99999908757381) / 
			horner4(1, 1.2906660139511692, 7.164571775410987, 10.559985088953114, 5.66336307375819, 1),
		2 + (math.Log(0.5)-2) * horner3(math.Log(0.5)-2, 0.00006979269679670452, 0.017110368846615806, 0.19338607770900237, 0.6666648896499793) / 
			horner2(math.Log(0.5)-2, 0.0188060684652668, 0.23451269827133317, 1),
		0,
		horner4(1, -2793.4565508841197, -1987.3632221106518, 385.7992853617571, 277.2362778379572, -7.840776922133643) /
            horner4(1, 280.6156995997829, 941.9414019982657, 190.64429338894644, -63.93540494358966, 1),
		-math.Exp(
				horner3(1, 0.16415668298255184, -3.334873920301941, 2.4831415860003747, 4.173424474574879) /
        		horner3(1, 0.031239411487374164, -1.2961659693400076, 4.517178492772906, 1),
		),
		-math.Exp(
			horner4(1, 0.000026987243254533254, -0.007692106448267341, 0.28793461719300206, -1.5267058884647018, -0.5370669268991288) /
			horner4(1, 3.6006502104930343e-6, -0.0015552463555591487, 0.08801194682489769, -0.8973922357575583, 1),
		),
		-1 -math.Sqrt(
			horner4(-0.1, 988.0070769375508, 1619.8111957356814, 989.2017745708083, 266.9332506485452, 26.875022558546036) /
			horner4(-0.1, -205.50469464210596, -270.0440832897079, -109.554245632316, -11.275355431307334, 1),
		),
	}
	for i := 1; i < 8; i++ {
		if i == 3 {
			branch = -1
			continue
		}
		if i == 5 {
			x = -math.E
		}
		if i == 7 {
			x = -0.1
		}
		t.Run(fmt.Sprintf("Test Branch %d ## n=%d", branch, i), func(t *testing.T) {
			p := pade{branch: branch, N: i}

			got := p.approximation(x)
			want := wants[i-1]

			if got != want {
				t.Errorf("Got %f but wanted %f", got, want)
			}
		})
	}
}