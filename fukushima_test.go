package lambertw

import (
	"fmt"
	"math"
	"testing"
)

func TestFukushima(t *testing.T) {
	var u float64
	jend := 40
	wmin,wmax:=-10.0,10.0
	dw:=(wmax-wmin)/float64(jend)

	res := fmt.Sprintf("\n%20s%25s%10s%10s\n","W","z","dW","dz")
	for j :=0; j <= jend; j++{
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
		res += fmt.Sprintf("%20.15f%25.15e%10.2e%10.2e\n",w,z,du,dx)
	}
//	t.Logf(res)
	t.Errorf(res)
}