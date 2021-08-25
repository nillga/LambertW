package lambertw

import "testing"

func TestHorner_recurse (t *testing.T) {
	t.Run("Order = 0 --> Recursion finish case", func(t *testing.T) {
		tests := []struct{
			name string
			tag interface{}
			want float64
		}{
			{"BranchPoint","branchPoint",0},{"AsymptoticPolynomialB",1,1},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				h := horner{test.tag,0}
				got := h.recurse(1,1)

				if got != test.want {
					t.Errorf("Got %f but wanted %f", got, test.want)
				}
			})
		}
	})
	t.Run("Order = 1 --> Recursion enabled", func(t *testing.T) {
		tests := []struct{
			name string
			tag interface{}
			want float64
		}{
			{"BranchPoint","branchPoint",1},{"AsymptoticPolynomialB",1,2},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				h := horner{test.tag,1}
				got := h.recurse(1,1)

				if got != test.want {
					t.Errorf("Got %f but wanted %f", got, test.want)
				}
			})
		}
	})
}
func TestHorner_recurse2 (t *testing.T) {
	tests := []struct{
		name string
		tag interface{}
		order int
		want float64
	}{
		{"Order = 0 --> Recursion finish case","AsymptoticPolynomialA",0,-1},
		{"Order = 1 --> Recursion enabled","AsymptoticPolynomialA",1,-1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := horner{test.tag,test.order}
			got := h.recurse2(0,0,1)

			if got != test.want {
				t.Errorf("Got %f but wanted %f", got, test.want)
			}
		})
	}
}

func TestHorner_eval (t *testing.T) {
	t.Run("Order = 0 --> Recursion finish case", func(t *testing.T) {
		tests := []struct{
			name string
			tag interface{}
			want float64
		}{
			{"BranchPoint","branchPoint",-1},{"AsymptoticPolynomialB",1,0},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				h := horner{test.tag,0}
				got := h.eval(0)

				if got != test.want {
					t.Errorf("Got %f but wanted %f", got, test.want)
				}
			})
		}
	})
	t.Run("Order = 1 --> Recursion enabled", func(t *testing.T) {
		tests := []struct{
			name string
			tag interface{}
			want float64
		}{
			{"BranchPoint","branchPoint",0},{"AsymptoticPolynomialB",1,1},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				h := horner{test.tag,1}
				got := h.eval(1)

				if got != test.want {
					t.Errorf("Got %f but wanted %f", got, test.want)
				}
			})
		}
	})
}
func TestHorner_eval2 (t *testing.T) {
	tests := []struct{
		name string
		tag interface{}
		order int
		want float64
	}{
		{"Order = 0 --> Recursion finish case","AsymptoticPolynomialA",0,-1},
		{"Order = 1 --> Recursion enabled","AsymptoticPolynomialA",1,-1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := horner{test.tag,test.order}
			got := h.eval2(0,1)

			if got != test.want {
				t.Errorf("Got %f but wanted %f", got, test.want)
			}
		})
	}
}