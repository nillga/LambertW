package lambertw

import "fmt"

type W struct {
	Branch  int
	results [2]float64
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