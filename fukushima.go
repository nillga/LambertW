package lambertw

import "math"

var q = []float64{
	-1,
	+1,
	-0.333333333333333333,
	+0.152777777777777778,
	-0.0796296296296296296,
	+0.0445023148148148148,
	-0.0259847148736037625,
	+0.0156356325323339212,
	-0.00961689202429943171,
	+0.00601454325295611786,
	-0.00381129803489199923,
	+0.00244087799114398267,
	-0.00157693034468678425,
	+0.00102626332050760715,
	-0.000672061631156136204,
	+0.000442473061814620910,
	-0.000292677224729627445,
	+0.000194387276054539318,
	-0.000129574266852748819,
	+0.0000866503580520812717,
	-0.0000581136075044138168,
  }

func fukushimaW0 (x float64) float64 {
	e := make([]float64, 66)
	g := make([]float64, 65)
	a := make([]float64, 12)
	b := make([]float64, 12)

	if e[0] == 0 {
		e1 := math.Exp(-1)
		ej := 1.0
		e[0] = math.E
		e[1] = 1
		g[0] = 0

		for j,jj := 1,2; jj < 66; jj++ {
			ej *= math.E
			e[jj] = e[j] * e1
			g[j] = float64(j) * ej
			j = jj
		}

		a[0] = math.Sqrt(e1)
		b[0] = 0.5

		for j,jj := 0,1; jj < 12; jj++ {
			a[jj] = math.Sqrt(a[j])
			b[jj] = b[j] * 0.5
			j = jj
		}
	}

	if math.Abs(x) < 0.05 {
		return lambertW0ZeroSeries(x)
	}
	if x < -0.35 {
		p2 := 2 * (math.E * x + 1)

		if p2 > 0 {
			return lambertWSeries(math.Sqrt(p2))
		}
		if p2 == 0 {
			return -1
		}
		return math.NaN()
	}

	var n int

	for n = 0; n <= 2; n++ {
		if g[n] > x {
//			goto line1	
		}
	}

	n = 2

	for j := 1; j <= 5; j++ {
		n *= 2
		if g[n] > x {
//			goto line2
		}
	}

	return math.NaN()
}

func fukushimaWm1 (x float64) float64 {
	e := make([]float64, 64)
	g := make([]float64, 64)
	a := make([]float64, 12)
	b := make([]float64, 12)

	if e[0] == 0 {
		e1 := math.Exp(-1)
		ej := e1
		e[0] = math.E
		g[0] = -e1

		for j,jj := 0,1; jj < 64; jj++ {
			ej *= e1
			e[jj] = e[j] * math.E
			g[j] = float64(-(jj+1)) * ej
			j = jj
		}

		a[0] = math.Sqrt(math.E)
		b[0] = 0.5

		for j,jj := 0,1; jj < 12; jj++ {
			a[jj] = math.Sqrt(a[j])
			b[jj] = b[j] * 0.5
			j = jj
		}
	}
}

func lambertW0ZeroSeries(x float64) float64 {
	return x*(1 -
    	x*(1 -
    	x*(1.5 -
    	x*(2.6666666666666666667 -
    	x*(5.2083333333333333333 -
    	x*(10.8 -
    	x*(23.343055555555555556 -
    	x*(52.012698412698412698 -
    	x*(118.62522321428571429 -
    	x*(275.57319223985890653 -
    	x*(649.78717234347442681 -
    	x*(1551.1605194805194805 -
    	x*(3741.4497029592385495 -
    	x*(9104.5002411580189358 -
    	x*(22324.308512706601434 -
    	x*(55103.621972903835338 -
    	x*136808.86090394293563))))))))))))))))
}

func lambertWSeries(x float64) float64 {
	ax := math.Abs(x)

	if ax < 0.01159 {
		// HORNER THIS
		return -1 + x*(1 + x*(q[2] + x*(q[3] + x*(q[4] + x*(q[5] + x*q[6])))))
	}
	if ax < 0.0766 {
		// HORNER THIS
		return -1 + x*(1 + x*(q[2] + x*(q[3] + x*(q[4] + x*(q[5] + x*(q[6] + x*(q[7] + x*(q[8] + x*(q[9] + x*q[10])))))))))
	}
	return -1 +
	x*(1 +
	x*(q[2] +
	x*(q[3] +
	x*(q[4] +
	x*(q[5] +
	x*(q[6] +
	x*(q[7] +
	x*(q[8] +
	x*(q[9] +
	x*(q[10] +
	x*(q[11] +
	x*(q[12] +
	x*(q[13] +
	x*(q[14] +
	x*(q[15] +
	x*(q[16] +
	x*(q[17] +
	x*(q[18] +
	x*(q[19] +
	x*q[20])))))))))))))))))))
}

func finalResult(w, y float64) float64 {
	f0 := w - y
	f1 := 1 + y
	f00 := f0 * f0
	f11 := f1 * f1
	f0y := f0 * y

	return w - 4 * f0 * (6 * f1 * (f11 + f0y) + f00 * y) / (f11 * (24 * f11 + 36 * f0y) + f00 * (6 * y * y + 8 * f1 * y + f0y))
}