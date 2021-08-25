package lambertw

type horner struct {
	tag interface{}
	order int
}

func (h *horner) eval(x float64) float64 {
	p := polynomial{h.tag, h.order}
	if h.order == 0 {
		return p.coeff()
	}
	horner := horner{h.tag, h.order-1}
	return horner.recurse(p.coeff(),x)
}
func (h *horner) recurse(term,x float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return term*x+p.coeff()
	}
	horner := horner{h.tag, h.order-1}
	return horner.recurse(term*x+p.coeff(), x)
}

func (h *horner) eval2(x,y float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return p.coeff2(y)
	}
	horner := horner{h.tag,h.order-1}
	return horner.recurse2(p.coeff2(y),x,y)
}
func (h *horner) recurse2(term,x,y float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return term*x+p.coeff2(y)
	}
	horner := horner{h.tag, h.order-1}
	return horner.recurse2(term*x+p.coeff2(y), x, y)
}

func horner0 (x, c0 float64) float64 {
	return c0
}
func horner1 (x, c1, c0 float64) float64 {
	return horner0(x, c1*x+c0)
}
func horner2 (x, c2, c1, c0 float64) float64 {
	return horner1(x, c2*x+c1, c0)
}
func horner3 (x, c3,c2,c1, c0 float64) float64 {
	return horner2(x, c3*x+c2,c1,c0)
}
func horner4 (x, c4,c3,c2,c1, c0 float64) float64 {
	return horner3(x, c4*x+c3,c2,c1,c0)
}
func horner5 (x, c5,c4,c3,c2,c1, c0 float64) float64 {
	return horner4(x, c5*x+c4,c3,c2,c1,c0)
}
func horner6 (x, c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return horner5(x, c6*x+c5,c4,c3,c2,c1,c0)
}
func horner7 (x, c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return horner6(x, c7*x+c6,c5,c4,c3,c2,c1,c0)
}
func horner8 (x, c8,c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return horner7(x, c8*x+c7,c6,c5,c4,c3,c2,c1,c0)
}
func horner9 (x, c9,c8,c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return horner8(x, c9*x+c8,c7,c6,c5,c4,c3,c2,c1,c0)
}