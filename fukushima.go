package lambertw

import "math"

type fukushima struct {
	n int
	branch int
	sgn int

	a []float64
	b []float64
	e []float64
	g []float64
}

func Fukushima(branch int, x float64) float64 {
	f := fukushima{branch: branch, sgn: -1}
	switch {
	case branch == 0:
		f.sgn = 1
		return f.w0(x)
	case branch < 0:
		f.branch = -1
		return f.wm1(x)
	default:
		return math.NaN()
	}
}

func (f *fukushima) w0(x float64) float64 {
	f.e = make([]float64, 66)
	f.g = make([]float64, 65)
	f.a = make([]float64, 12)
	f.b = make([]float64, 12)

	if f.e[0] == 0 {
		e1 := math.Exp(-1)
		ej := 1.0
		f.e[0] = math.E
		f.e[1] = 1
		f.g[0] = 0

		for j,jj := 1,2; jj < 66; jj++ {
			ej *= math.E
			f.e[jj] = f.e[j] * e1
			f.g[j] = float64(j) * ej
			j = jj
		}

		f.a[0] = math.Sqrt(e1)
		f.b[0] = 0.5

		for j,jj := 0,1; jj < 12; jj++ {
			f.a[jj] = math.Sqrt(f.a[j])
			f.b[jj] = f.b[j] * 0.5
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

	for f.n = 0; f.n <= 2; f.n++ {
		if f.g[f.n] > x {
			return f.step1(x)
		}
	}

	f.n = 2

	for j := 1; j <= 5; j++ {
		f.n *= 2
		if f.g[f.n] > x {
			return f.step2(x)
		}
	}

	return math.NaN()
}

func (f *fukushima) step2(x float64) float64 {
	nh := f.n / 2

	for j := 1; j <= 5; j++ {
		nh /= 2
		if nh <= 0 {
			return f.step1(x)
		}
		if f.g[f.n-nh + f.branch] > x {
			f.n -= nh
		}
	}
	return f.step1(x)
}

func (f *fukushima) step1(x float64) float64 {
	f.n--
	jmax := 8 - 3 * f.branch

	if f.branch == 0 {
		switch {
		case x <= -0.36:
			jmax = 12
		case x <= -0.3:
			jmax = 11
		case f.n <= 0:
			jmax = 10
		case f.n <= 1:
			jmax = 9
		}
	} else {
		switch {
		case f.n >= 8:
			jmax = 8
		case f.n >= 3:
			jmax = 9
		case f.n >= 2:
			jmax = 10
		}
	}

	y := x * f.e[f.n + f.sgn]
	w := float64(f.n * f.sgn)

	for j := 0; j < jmax; j++ {
		wj := w + float64(f.sgn) * f.b[j]
		yj := y * f.a[j]

		if wj < yj {
			w = wj
			y = yj
		}
	}
	return finalResult(w,y)
}

func (f *fukushima) wm1(x float64) float64 {
	f.e = make([]float64, 64)
	f.g = make([]float64, 64)
	f.a = make([]float64, 12)
	f.b = make([]float64, 12)

	if f.e[0] == 0 {
		e1 := math.Exp(-1)
		ej := e1
		f.e[0] = math.E
		f.g[0] = -e1

		for j,jj := 0,1; jj < 64; jj++ {
			ej *= e1
			f.e[jj] = f.e[j] * math.E
			f.g[j] = float64(-(jj+1)) * ej
			j = jj
		}

		f.a[0] = math.Sqrt(math.E)
		f.b[0] = 0.5

		for j,jj := 0,1; jj < 12; jj++ {
			f.a[jj] = math.Sqrt(f.a[j])
			f.b[jj] = f.b[j] * 0.5
			j = jj
		}
	}

	if x >= 0 {
		return math.NaN()
	}

	if x < -0.35 {
		p2 := 2 * (math.E * x + 1)
		if p2 > 0 {
			return lambertWSeries(-math.Sqrt(p2))
		}
		if p2 == 0 {
			return -1
		}
		return math.NaN()
	}
	f.n = 2

	if f.g[f.n-1] > x {
		return f.step1(x)
	}
	for j := 1; j <= 5; j++ {
		f.n *= 2
		if f.g[f.n-1] > x {
			return f.step2(x)
		}
	}
	return math.NaN()
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