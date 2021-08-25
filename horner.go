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

func (h *horner) eval2(x,y float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return p.coeff2(y)
	}
	horner := horner{h.tag,h.order-1}
	return horner.recurse2(p.coeff2(y),x,y)
}

func (h *horner) recurse(term,x float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return term*x+p.coeff()
	}
	horner := horner{h.tag, h.order-1}
	return horner.recurse(term*x+p.coeff(), x)
}

func (h *horner) recurse2(term,x,y float64) float64 {
	p := polynomial{h.tag,h.order}
	if h.order == 0 {
		return term*x+p.coeff2(y)
	}
	horner := horner{h.tag, h.order-1}
	return horner.recurse2(term*x+p.coeff2(y), x, y)
}