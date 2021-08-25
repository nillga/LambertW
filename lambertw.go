package lambertw

import (
	"fmt"
	"math"
)

type W struct {
	Branch  int
	X float64
	results [2]float64
}

type Branch struct {
	Branch int
	Order int

}

type Horner struct {
	Tag interface{}
	Order int
}

type Polynomial struct {
	Tag interface{}
	Order int
}

type Iterator struct {
	IterationStep IterationStep
	Depth int
}

type Pade struct {
	Branch int
	N int
}

type HORNERCOEFF func(struct{Name string},float64) float64
type IterationStep func(float64,float64) float64

var Branchpoints = map[int]float64 {
	0:-1, 1:1, 2:-0.333333333333333333e0, 3:0.152777777777777777e0, 4: -0.79629629629629630e-1,
	5: 0.44502314814814814e-1, 6:-0.25984714873603760e-1, 7:0.15635632532333920e-1, 8:-0.96168920242994320e-2,
	9:  0.60145432529561180e-2, 10: -0.38112980348919993e-2, 11:  0.24408779911439826e-2, 12: -0.15769303446867841e-2,
	13: 0.10262633205076071e-2, 14: -0.67206163115613620e-3, 15:  0.44247306181462090e-3, 16: -0.29267722472962746e-3,
	17:  0.19438727605453930e-3, 18: -0.12957426685274883e-3, 19:  0.86650358052081260e-4,
}

func HornerLib(t interface{},order int) float64 {
	switch t {
	case "BranchPoint":
		return Branchpoints[order]
	default:
		return 0
	}
}

func (w *W) Setup(branch int) {
	if branch < 0 {
		w.Branch = -1
		return
	}
	if branch > 0 {
		w.Branch = 1
		fmt.Printf("Input %d is invalid, valid branches are 0 and -1", branch)
		return
	}
}

func (w *W) Router() {
	if w.X < 1.38 {
		if w.X < -0.311 {
			if w.X < -0.367679 {
				b := Branch{Order: 8, Branch: 0}
				w.results[0] = b.BranchPointExpansion(w.X)
			}
			i := Iterator{HalleyStep,1}
			b := Branch{Order: 10, Branch: 0}
			i.Recurse(w.X,b.BranchPointExpansion(w.X))
		}
		//Iterator<double, HalleyStep<double> >::Depth<1>::Recurse(x, Pade<double, 0, 1>::Approximation(x))
	} else {
		if w.X < 236 { 
			//Iterator<double, HalleyStep<double> >::Depth<1>::Recurse(x, Pade<double, 0, 2>::Approximation(x)) 
		}
		//Iterator<double, HalleyStep<double> >::Depth<1>::Recurse(x, Branch<double, 0>::AsymptoticExpansion<6-1>(x))
	}
}

func (b *Branch) BranchPointExpansion(x float64) float64 {
	sgn := float64(2 * b.Branch + 1)
	h := Horner{"BranchPoint", b.Order}
	return h.Eval(sgn * math.Sqrt(2.0 * (math.E * x + 1)))
}

func (h *Horner) Eval(x float64) float64 {
	p := Polynomial{h.Tag, h.Order}
	if h.Order == 0 {
		p.Coeff()
	}
	horner := Horner{h.Tag, h.Order-1}
	return horner.Recurse(p.Coeff(),x)
}

func (h *Horner) Recurse(term,x float64) float64 {
	p := Polynomial{h.Tag,h.Order}
	if h.Order == 0 {
		return term*x+p.Coeff()
	}
	horner := Horner{h.Tag, h.Order-1}
	return horner.Recurse(term*x+p.Coeff(), x)
}

func (p *Polynomial) Coeff() float64 {
	return HornerLib(p.Tag, p.Order)
}

func (i *Iterator) Recurse(x,w float64) float64 {
	if i.Depth == 1 {
		return i.IterationStep(x,w)
	}
	if i.Depth == 0 {
		return w
	}
	i.Depth--
	return i.Recurse(x, i.IterationStep(x,w))
}

func HalleyStep (x,w float64) float64 {
	ew := math.Exp(w)
	wew := w * ew
	wewx := wew-x
	w1 := w + 1
	return w - wewx / (ew * w1 - (w + 2) * wewx/(2*w1))
}

func (p *Pade) Approximation(x float64) {
	if p.Branch == 0 {
		if p.N == 1 {

		}
	}
}

func Horner0 (x, c0 float64) float64 {
	return c0
}
func Horner1 (x, c1, c0 float64) float64 {
	return Horner0(x, c1*x+c0)
}
func Horner2 (x, c2, c1, c0 float64) float64 {
	return Horner1(x, c2*x+c1, c0)
}
func Horner3 (x, c3,c2,c1, c0 float64) float64 {
	return Horner2(x, c3*x+c2,c1,c0)
}
func Horner4 (x, c4,c3,c2,c1, c0 float64) float64 {
	return Horner3(x, c4*x+c3,c2,c1,c0)
}
func Horner5 (x, c5,c4,c3,c2,c1, c0 float64) float64 {
	return Horner4(x, c5*x+c4,c3,c2,c1,c0)
}
func Horner6 (x, c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return Horner5(x, c6*x+c5,c4,c3,c2,c1,c0)
}
func Horner7 (x, c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return Horner6(x, c7*x+c6,c5,c4,c3,c2,c1,c0)
}
func Horner8 (x, c8,c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return Horner7(x, c8*x+c7,c6,c5,c4,c3,c2,c1,c0)
}
func Horner9 (x, c9,c8,c7,c6,c5,c4,c3,c2,c1, c0 float64) float64 {
	return Horner8(x, c9*x+c8,c7,c6,c5,c4,c3,c2,c1,c0)
}