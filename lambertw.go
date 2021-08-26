package lambertw

/*
*	## CREDITS GO TO DARKO VEBERIC, WHO IMPLEMENTED THIS METHOD FIRST, IN C++ ##
*
*	 TODOS:
*	 ------
*
*	- Implement Tests -- Coverage: 52.2%
*	- Benchmark this mess
*	- Split in multiple files -- Routing file // methods
*	- refactoring
*	- add / test mezö integrals and / or fukushima methodz
*
*/

import (
	"fmt"
	"math"
)

type w struct {
	branch  int
	x float64
}
type branch struct {
	branch int
	order int
	sgn float64
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

type iterationStep func(float64,float64) float64

var (
	branchPoints = map[int]float64 {
		0:-1, 1:1, 2:-0.333333333333333333e0, 3:0.152777777777777777e0, 4: -0.79629629629629630e-1,
		5: 0.44502314814814814e-1, 6:-0.25984714873603760e-1, 7:0.15635632532333920e-1, 8:-0.96168920242994320e-2,
		9:  0.60145432529561180e-2, 10: -0.38112980348919993e-2, 11:  0.24408779911439826e-2, 12: -0.15769303446867841e-2,
		13: 0.10262633205076071e-2, 14: -0.67206163115613620e-3, 15:  0.44247306181462090e-3, 16: -0.29267722472962746e-3,
		17:  0.19438727605453930e-3, 18: -0.12957426685274883e-3, 19:  0.86650358052081260e-4,
	}
	asymptoticBs = [][]float64{
		{0,-1},{0,1},{0,-1,0.5},{0,1,-3.0/2.0,1.0/3.0},{0,-1,3,-11.0/6.0,0.25},{0,1,-5,35.0/6.0,-25.0/12.0,0.2},
	}
)

func W(branch int, x float64) float64 {
	wOfX := new(w)
	wOfX.setup(branch)
	return wOfX.router()
}
func (w *w) setup(branch int) {
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

func (b *branch) branchPointExpansion(x float64) float64 {
	h := horner{"branchPoint", b.order}
	return h.eval(b.sgn * math.Sqrt(2.0 * (math.E * x + 1)))
}
func (b *branch) asymptoticExpansion(x float64) float64 {
	logsx := math.Log(b.sgn * x)
	logslogsx := math.Log(b.sgn * logsx)

	return asymptoticExpansionImpl(logsx,logslogsx,b.order)
}
func (b *branch) logRecursion(x float64) float64 {
	l := logRecursionImpl{sgn: b.sgn, branch: b.branch, order: b.order}
	return l.step(math.Log(b.sgn * x))
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

func asymptoticExpansionImpl(a,b float64, order int) float64 {
	h := horner{"AsymptoticPolynomialA", order}

	return a + h.eval2(1/a,b)
}
func (l *logRecursionImpl) step(logsx float64) float64 {
	if l.order == 0 {
		return logsx
	}
	logRecursionImpl := logRecursionImpl{sgn: l.sgn, branch: l.branch, order: l.order - 1}
	return logsx - math.Log(l.sgn * logRecursionImpl.step(logsx))
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
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
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