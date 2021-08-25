package lambertw

func (w *w) router() (res float64) {
	if w.branch == 0 {
		res = w.router0()
	}
	if w.branch == -1 {
		res = w.router1()
	}
	return res
}
func (w *w) router0() float64 {
	i := iterator{halleyStep, 1}
	b := branch{branch: 0, sgn: 1}
	if w.x < 1.38 {
		if w.x < -0.311 {
			if w.x < -0.367679 {
				b.order = 8
				return b.branchPointExpansion(w.x)
			}
			b.order = 10
			return i.recurse(w.x, b.branchPointExpansion(w.x))
		}
		p := pade{branch: 0, N: 1}
		return i.recurse(w.x, p.approximation(w.x))
	}
	if w.x < 236 {
		p := pade{branch: 0, N: 2}
		return i.recurse(w.x, p.approximation(w.x))
	}
	b.order = 6 - 1
	return i.recurse(w.x, b.asymptoticExpansion(w.x))
}
func (w *w) router1() float64 {
	i := iterator{halleyStep, 1}
	b := branch{branch: -1, sgn: -1}
	if w.x < -0.0509 {
		if w.x < -0.366079 {
			if w.x < -0.367579 {
				b.order = 8
				return b.branchPointExpansion(w.x)
			}
			b.order = 4
			return i.recurse(w.x, b.branchPointExpansion(w.x))
		}
		if w.x < -0.289379 {
			p := pade{branch: -1, N: 7}
			return i.recurse(w.x, p.approximation(w.x))
		}
		p := pade{branch: -1, N: 4}
		return i.recurse(w.x, p.approximation(w.x))
	}
	if w.x < -0.000131826 {
		p := pade{branch: -1, N: 5}
		return i.recurse(w.x, p.approximation(w.x))
	}
	if w.x < -6.30957e-31 {
		p := pade{branch: -1, N: 6}
		return i.recurse(w.x, p.approximation(w.x))
	}
	b.order = 3
	return i.recurse(w.x, b.logRecursion(w.x))
}