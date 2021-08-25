package lambertw

import (
	"fmt"
	"math"
)

type W struct {
	Branch  int
	X float64
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

type AsymptoticPolynomialB struct {
	Order int
}

type HORNERFUNC func(float64) float64
type IterationStep func(float64,float64) float64

var Branchpoints = map[int]float64 {
	0:-1, 1:1, 2:-0.333333333333333333e0, 3:0.152777777777777777e0, 4: -0.79629629629629630e-1,
	5: 0.44502314814814814e-1, 6:-0.25984714873603760e-1, 7:0.15635632532333920e-1, 8:-0.96168920242994320e-2,
	9:  0.60145432529561180e-2, 10: -0.38112980348919993e-2, 11:  0.24408779911439826e-2, 12: -0.15769303446867841e-2,
	13: 0.10262633205076071e-2, 14: -0.67206163115613620e-3, 15:  0.44247306181462090e-3, 16: -0.29267722472962746e-3,
	17:  0.19438727605453930e-3, 18: -0.12957426685274883e-3, 19:  0.86650358052081260e-4,
}
var AsymptoticBs = [][]float64{
	{0,-1},{0,1},{0,-1,0.5},{0,1,-3.0/2.0,1.0/3.0},{0,-1,3,-11.0/6.0,0.25},{0,1,-5,35.0/6.0,-25.0/12.0,0.2},
}

func HornerLib(t interface{},order int, x float64) float64 {
	switch param := t.(type) {
	case string:
		switch t {
		case "BranchPoint":
			return Branchpoints[order]
		case "AsymptoticPolynomialA":
			switch order {
			case 0:
				return -x
			case 1:
				return x
			case 2:
			case 3:
			case 4:
			case 5:
				h := Horner{AsymptoticPolynomialB{order},order}
				return h.Eval(x)
			}
		}	
	case int:
		return AsymptoticBs[param][order]
	}
	return 0
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

func (w *W) Router() (res float64) {
	if w.X == 0 {
		res = w.Router0()
	}
	if w.X == -1 {
		res = w.Router1()
	}
	return res
}

func (w *W) Router0() float64{
	if w.X < 1.38 {
		if w.X < -0.311 {
			if w.X < -0.367679 {
				b := Branch{Order: 8, Branch: 0}
				return b.BranchPointExpansion(w.X)
			}
			i := Iterator{HalleyStep,1}
			b := Branch{Order: 10, Branch: 0}
			return i.Recurse(w.X,b.BranchPointExpansion(w.X))
		}
		i := Iterator{HalleyStep,1}
		p := Pade{Branch: 0, N: 1}
		return i.Recurse(w.X,p.Approximation(w.X))
	}
	if w.X < 236 { 
		i := Iterator{HalleyStep,1}
		p := Pade{Branch: 0, N: 2}
		return i.Recurse(w.X,p.Approximation(w.X)) 
	}
	i := Iterator{HalleyStep,1}
	b := Branch{Order: 6-1, Branch: 0}
	return i.Recurse(w.X, b.AsymptoticExpansion(w.X))
}

func (w *W) Router1() float64 {
	if w.X < -0.0509 {
		if w.X < -0.366079 {
			if w.X < -0.367579 {
			//	(Branch<d, -1>::BranchPointExpansion<8>(x))
			}
		//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Branch<d, -1>::BranchPointExpansion<4>(x)))
		}
		if w.X < -0.289379 {
		//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Pade<d, -1, 7>::Approximation(x)))
		}
	//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Pade<d, -1, 4>::Approximation(x)))
	}
	if w.X < -0.000131826 {
	//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Pade<d, -1, 5>::Approximation(x)))
	}
	if w.X < -6.30957e-31 {
	//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Pade<d, -1, 6>::Approximation(x)))
	}
//	(Iterator<d, HalleyStep<d> >::Depth<1>::Recurse(x, Branch<d, -1>::LogRecursion<3>(x)))
	return 0
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

func (h *Horner) Eval2(x,y float64) float64 {
	p := Polynomial{h.Tag,h.Order}
	if h.Order == 0 {
		return p.Coeff2(y)
	}
	horner := Horner{h.Tag,h.Order-1}
	return horner.Recurse2(p.Coeff2(y),x,y)
}

func (h *Horner) Recurse(term,x float64) float64 {
	p := Polynomial{h.Tag,h.Order}
	if h.Order == 0 {
		return term*x+p.Coeff()
	}
	horner := Horner{h.Tag, h.Order-1}
	return horner.Recurse(term*x+p.Coeff(), x)
}

func (h *Horner) Recurse2(term,x,y float64) float64 {
	p := Polynomial{h.Tag,h.Order}
	if h.Order == 0 {
		return term*x+p.Coeff2(y)
	}
	horner := Horner{h.Tag, h.Order-1}
	return horner.Recurse2(term*x+p.Coeff2(y), x, y)
}

func (p *Polynomial) Coeff() float64 {
	return HornerLib(p.Tag, p.Order,0)
}

func (p *Polynomial) Coeff2(x float64) float64 {
	return HornerLib(p.Tag, p.Order,x)
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

func (p *Pade) Approximation(x float64) float64 {
	if p.Branch == 0 {
		if p.N == 1 {
			return x * Horner4(x, 0.07066247420543414, 2.4326814530577687, 6.39672835731526, 4.663365025836821, 0.99999908757381) / 
			Horner4(x, 1.2906660139511692, 7.164571775410987, 10.559985088953114, 5.66336307375819, 1)
		}
		if p.N == 2 {
			y := math.Log(0.5*x)-2
			return 2 + y * Horner3(y, 0.00006979269679670452, 0.017110368846615806, 0.19338607770900237, 0.6666648896499793) / 
			Horner2(y, 0.0188060684652668, 0.23451269827133317, 1)
		}
	}
	if p.Branch == -1 {
		switch p.N {
		case 4:
			return Horner4(x, -2793.4565508841197, -1987.3632221106518, 385.7992853617571, 277.2362778379572, -7.840776922133643) /
            Horner4(x, 280.6156995997829, 941.9414019982657, 190.64429338894644, -63.93540494358966, 1)
		case 5:
			y := math.Log(-x)
			return -math.Exp(
				Horner3(y, 0.16415668298255184, -3.334873920301941, 2.4831415860003747, 4.173424474574879) /
        		Horner3(y, 0.031239411487374164, -1.2961659693400076, 4.517178492772906, 1),
			)
		case 6:
			y := math.Log(-x)
			return -math.Exp(
				Horner4(y, 0.000026987243254533254, -0.007692106448267341, 0.28793461719300206, -1.5267058884647018, -0.5370669268991288) /
				Horner4(y, 3.6006502104930343e-6, -0.0015552463555591487, 0.08801194682489769, -0.8973922357575583, 1),
			)
		case 7:
			return -1 -math.Sqrt(
				Horner4(x, 988.0070769375508, 1619.8111957356814, 989.2017745708083, 266.9332506485452, 26.875022558546036) /
        		Horner4(x, -205.50469464210596, -270.0440832897079, -109.554245632316, -11.275355431307334, 1),
			)		
		}
	}
	return 0
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

func (b *Branch) AsymptoticExpansion(x float64) float64 {
	sgn := float64(2 * b.Branch + 1)
	logsx := math.Log(sgn * x)
	logslogsx := math.Log(sgn * logsx)

	return AsymptoticExpansionImpl(logsx,logslogsx,b.Order)
}

func AsymptoticExpansionImpl(a,b float64, order int) float64 {
	h := Horner{"AsymptoticPolynomialA", order}

	return a + h.Eval2(1/a,b)
}