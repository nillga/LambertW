package lambertw

/*
*	## CREDITS GO TO DARKO VEBERIC, WHO IMPLEMENTED THIS METHOD FIRST, IN C++ ##
*
*	 TODOS:
*	 ------
* 
*	- Simple Startup + exportation
*	- Implement Tests
*	- Benchmark this mess
*	- Split in multiple files
*	- refactoring
*	- add / test mezö integrals and / or fukushima methodz
*
*/

import (
	"fmt"
	"math"
)

type W struct {
	branch  int
	x float64
}

type branch struct {
	branch int
	order int
}

type horner struct {
	tag interface{}
	order int
}

type polynomial struct {
	tag interface{}
	order int
}

type iterator struct {
	iterationStep iterationStep
	depth int
}

type pade struct {
	branch int
	N int
}

type logRecursionImpl struct {
	sgn float64
	order int
	branch int
}

type HORNERFUNC func(float64) float64
type iterationStep func(float64,float64) float64

var branchPoints = map[int]float64 {
	0:-1, 1:1, 2:-0.333333333333333333e0, 3:0.152777777777777777e0, 4: -0.79629629629629630e-1,
	5: 0.44502314814814814e-1, 6:-0.25984714873603760e-1, 7:0.15635632532333920e-1, 8:-0.96168920242994320e-2,
	9:  0.60145432529561180e-2, 10: -0.38112980348919993e-2, 11:  0.24408779911439826e-2, 12: -0.15769303446867841e-2,
	13: 0.10262633205076071e-2, 14: -0.67206163115613620e-3, 15:  0.44247306181462090e-3, 16: -0.29267722472962746e-3,
	17:  0.19438727605453930e-3, 18: -0.12957426685274883e-3, 19:  0.86650358052081260e-4,
}
var asymptoticBs = [][]float64{
	{0,-1},{0,1},{0,-1,0.5},{0,1,-3.0/2.0,1.0/3.0},{0,-1,3,-11.0/6.0,0.25},{0,1,-5,35.0/6.0,-25.0/12.0,0.2},
}

func hornerLib(t interface{},order int, x float64) float64 {
	switch param := t.(type) {
	case string:
		switch t {
		case "branchPoint":
			return branchPoints[order]
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
				h := horner{order,order}
				return h.eval(x)
			}
		}	
	case int:
		return asymptoticBs[param][order]
	}
	return 0
}

func (w *W) setup(branch int) {
	if branch < 0 {
		w.branch = -1
		return
	}
	if branch > 0 {
		w.branch = 1
		fmt.Printf("Input %d is invalid, valid branches are 0 and -1", branch)
		return
	}
}

func (w *W) router() (res float64) {
	if w.branch == 0 {
		res = w.router0()
	}
	if w.branch == -1 {
		res = w.router1()
	}
	return res
}

func (w *W) router0() float64{
	if w.x < 1.38 {
		if w.x < -0.311 {
			if w.x < -0.367679 {
				b := branch{order: 8, branch: 0}
				return b.branchPointExpansion(w.x)
			}
			i := iterator{halleyStep,1}
			b := branch{order: 10, branch: 0}
			return i.recurse(w.x,b.branchPointExpansion(w.x))
		}
		i := iterator{halleyStep,1}
		p := pade{branch: 0, N: 1}
		return i.recurse(w.x,p.approximation(w.x))
	}
	if w.x < 236 { 
		i := iterator{halleyStep,1}
		p := pade{branch: 0, N: 2}
		return i.recurse(w.x,p.approximation(w.x)) 
	}
	i := iterator{halleyStep,1}
	b := branch{order: 6-1, branch: 0}
	return i.recurse(w.x, b.asymptoticExpansion(w.x))
}

func (w *W) router1() float64 {
	if w.x < -0.0509 {
		if w.x < -0.366079 {
			if w.x < -0.367579 {
				b := branch{order: 8, branch: -1}
				return b.branchPointExpansion(w.x)
			}
			i := iterator{halleyStep, 1}
			b := branch{order: 4, branch: -1}
			return i.recurse(w.x, b.branchPointExpansion(w.x))
		}
		if w.x < -0.289379 {
			i := iterator{halleyStep,1}
			p := pade{branch: -1, N: 7}
			return i.recurse(w.x, p.approximation(w.x))
		}
		i := iterator{halleyStep,1}
		p := pade{branch: -1, N: 4}
		return i.recurse(w.x, p.approximation(w.x))
	}
	if w.x < -0.000131826 {
		i:= iterator{halleyStep,1}
		p := pade{branch: -1, N: 5}
		return i.recurse(w.x, p.approximation(w.x))
	}
	if w.x < -6.30957e-31 {
		i:= iterator{halleyStep,1}
		p := pade{branch: -1, N: 6}
		return i.recurse(w.x, p.approximation(w.x))
	}
	i:= iterator{halleyStep,1}
	b := branch{branch: -1, order: 3}
	return i.recurse(w.x, b.logRecursion(w.x))
}

func (b *branch) branchPointExpansion(x float64) float64 {
	sgn := float64(2 * b.branch + 1)
	h := horner{"branchPoint", b.order}
	return h.eval(sgn * math.Sqrt(2.0 * (math.E * x + 1)))
}

func (h *horner) eval(x float64) float64 {
	p := polynomial{h.tag, h.order}
	if h.order == 0 {
		p.coeff()
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

func (p *polynomial) coeff() float64 {
	return hornerLib(p.tag, p.order,0)
}

func (p *polynomial) coeff2(x float64) float64 {
	return hornerLib(p.tag, p.order,x)
}

func (i *iterator) recurse(x,w float64) float64 {
	if i.depth == 1 {
		return i.iterationStep(x,w)
	}
	if i.depth == 0 {
		return w
	}
	i.depth--
	return i.recurse(x, i.iterationStep(x,w))
}

func halleyStep (x,w float64) float64 {
	ew := math.Exp(w)
	wew := w * ew
	wewx := wew-x
	w1 := w + 1
	return w - wewx / (ew * w1 - (w + 2) * wewx/(2*w1))
}

func (p *pade) approximation(x float64) float64 {
	if p.branch == 0 {
		if p.N == 1 {
			return x * horner4(x, 0.07066247420543414, 2.4326814530577687, 6.39672835731526, 4.663365025836821, 0.99999908757381) / 
			horner4(x, 1.2906660139511692, 7.164571775410987, 10.559985088953114, 5.66336307375819, 1)
		}
		if p.N == 2 {
			y := math.Log(0.5*x)-2
			return 2 + y * horner3(y, 0.00006979269679670452, 0.017110368846615806, 0.19338607770900237, 0.6666648896499793) / 
			horner2(y, 0.0188060684652668, 0.23451269827133317, 1)
		}
	}
	if p.branch == -1 {
		switch p.N {
		case 4:
			return horner4(x, -2793.4565508841197, -1987.3632221106518, 385.7992853617571, 277.2362778379572, -7.840776922133643) /
            horner4(x, 280.6156995997829, 941.9414019982657, 190.64429338894644, -63.93540494358966, 1)
		case 5:
			y := math.Log(-x)
			return -math.Exp(
				horner3(y, 0.16415668298255184, -3.334873920301941, 2.4831415860003747, 4.173424474574879) /
        		horner3(y, 0.031239411487374164, -1.2961659693400076, 4.517178492772906, 1),
			)
		case 6:
			y := math.Log(-x)
			return -math.Exp(
				horner4(y, 0.000026987243254533254, -0.007692106448267341, 0.28793461719300206, -1.5267058884647018, -0.5370669268991288) /
				horner4(y, 3.6006502104930343e-6, -0.0015552463555591487, 0.08801194682489769, -0.8973922357575583, 1),
			)
		case 7:
			return -1 -math.Sqrt(
				horner4(x, 988.0070769375508, 1619.8111957356814, 989.2017745708083, 266.9332506485452, 26.875022558546036) /
        		horner4(x, -205.50469464210596, -270.0440832897079, -109.554245632316, -11.275355431307334, 1),
			)		
		}
	}
	return 0
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

func (b *branch) asymptoticExpansion(x float64) float64 {
	sgn := float64(2 * b.branch + 1)
	logsx := math.Log(sgn * x)
	logslogsx := math.Log(sgn * logsx)

	return asymptoticExpansionImpl(logsx,logslogsx,b.order)
}

func asymptoticExpansionImpl(a,b float64, order int) float64 {
	h := horner{"AsymptoticPolynomialA", order}

	return a + h.eval2(1/a,b)
}

func (b *branch) logRecursion(x float64) float64 {
	sgn := float64(2 * b.branch + 1)

	l := logRecursionImpl{sgn: sgn, branch: b.branch, order: b.order}
	return l.Step(math.Log(sgn * x))
}

func (l *logRecursionImpl) Step(logsx float64) float64 {
	if l.order == 0 {
		return logsx
	}
	logRecursionImpl := logRecursionImpl{sgn: l.sgn, branch: l.branch, order: l.order - 1}
	return logsx - math.Log(l.sgn * logRecursionImpl.Step(logsx))
}
