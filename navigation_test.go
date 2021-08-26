package lambertw

import (
	"testing"
	"math"
)	

func TestW_router_integration(t *testing.T) {
	t.Run("Upper", func(t *testing.T) {
		got := W(0, math.E)
		if got != 1 {
			t.Errorf("Got %f but expected 1", got)
		}
	})
	t.Run("Lower", func(t *testing.T) {
		got := W(-1, -0.1)
		w := w{-1,-0.1}
		want := w.router1()
		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
}

func TestW_router0(t *testing.T) {
	t.Run("Case a", func(t *testing.T) {
		w := w{0, -0.3676791}
		got := w.router0()

		b := branch{branch: 0, sgn: 1, order: 8}
		want := b.branchPointExpansion(w.x)

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case b", func(t *testing.T) {
		w := w{0, -0.367679}
		got := w.router0()

		b := branch{branch: 0, sgn: 1, order: 10}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, b.branchPointExpansion(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case c", func(t *testing.T) {
		w := w{0, -0.311}
		got := w.router0()

		p := pade{branch: 0, N: 1}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case d", func(t *testing.T) {
		w := w{0, 235}
		got := w.router0()

		p := pade{branch: 0, N: 2}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case e", func(t *testing.T) {
		w := w{0, 236}
		got := w.router0()

		b := branch{branch: 0, sgn: 1, order: 5}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, b.asymptoticExpansion(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
}

func TestW_router1(t *testing.T) {
	t.Run("Case a", func(t *testing.T) {
		w := w{-1, -0.3675791}
		got := w.router1()

		b := branch{branch: -1, sgn: -1, order: 8}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, b.branchPointExpansion(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case b", func(t *testing.T) {
		w := w{-1, -0.367579}
		got := w.router1()

		b := branch{branch: -1, sgn: -1, order: 4}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, b.branchPointExpansion(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case c", func(t *testing.T) {
		w := w{-1, -0.2893791}
		got := w.router1()

		i := iterator{halleyStep, 1}
		p := pade{branch: -1, N: 7}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case d", func(t *testing.T) {
		w := w{-1, -0.289379}
		got := w.router1()

		i := iterator{halleyStep, 1}
		p := pade{branch: -1, N: 4}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case e", func(t *testing.T) {
		w := w{-1, -0.0001318261}
		got := w.router1()

		i := iterator{halleyStep, 1}
		p := pade{branch: -1, N: 5}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case f", func(t *testing.T) {
		w := w{-1, -6.30957e-30}
		got := w.router1()

		i := iterator{halleyStep, 1}
		p := pade{branch: -1, N: 6}
		want := i.recurse(w.x, p.approximation(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
	t.Run("Case g", func(t *testing.T) {
		w := w{-1, -6.30957e-32}
		got := w.router1()

		b := branch{branch: -1, sgn: -1, order: 3}
		i := iterator{halleyStep, 1}
		want := i.recurse(w.x, b.logRecursion(w.x))

		if got != want {
			t.Errorf("Got %f but expected %f", got, want)
		}
	})
}
