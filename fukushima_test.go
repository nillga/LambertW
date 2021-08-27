package lambertw

import (
	"fmt"
	"math"
	"testing"
)

func TestFukushima(t *testing.T) {
	var u float64
	jend := 256
	wmin,wmax:=-64.0,64.0
	dw:=(wmax-wmin)/float64(jend)

	res := fmt.Sprintf("\n%20s%25s%10s%10s\n","W","z","dW","dz")
	for j :=0; j < jend; j++{
		w := wmin + dw*float64(j)
		z := w*math.Exp(w)
		if w < -1 {
			u = Fukushima(-1,z)
		} else {
			u = Fukushima(0,z)
		}
		x := u * math.Exp(u)
		du := (u-w)/(1e-16+math.Abs(w))
		dx := (x-z)/(1e-16+math.Abs(z))
		res += fmt.Sprintf("%20.15f%25.15e%10.2e%10.2e%25.15f\n",w,z,du,dx,u)
	}
	t.Logf(res)
}

func TestLambertW0Series(t *testing.T) {
	t.Run("Insert Zero", func(t *testing.T) {
		got := lambertW0ZeroSeries(0)

		if got != 0 {
			t.Errorf("Got %f but wanted 0", got)
		}
	})
	t.Run("Insert One", func(t *testing.T) {
		got := lambertW0ZeroSeries(1)
		want := 97572.747613193801953457295894622802734375

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
}
func TestLambertSeries(t *testing.T) {
	t.Run("1st Bracket", func(t *testing.T) {
		got := lambertWSeries(0)
		want := -1.0

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
	t.Run("2nd Bracket", func(t *testing.T) {
		got := lambertWSeries(-0.2)
		want := -1.2146990950079346038847916133818216621875762939453125

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
	t.Run("3rd Bracket", func(t *testing.T) {
		got := lambertWSeries(1)
		want := -0.23198430458523233710366184823215007781982421875

		if got != want {
			t.Errorf("Got %f but wanted %f", got, want)
		}
	})
}