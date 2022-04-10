package iter1

import (
	"math"
	"math/cmplx"
)

type FourierThing struct {
	num   complex128
	freq  float64
	amp   float64
	phase float64
}

func dft(signal []float64) []FourierThing {
	X := make([]FourierThing, len(signal))
	N := len(signal)
	var result complex128
	for k := range X {
		X[k].freq = float64(k)
		for n := 0; n < N; n++ {
			result += complex(signal[n], 0) * cmplx.Exp((-1i*2*complex(math.Pi, 0)*complex(float64(n*k), 0))/complex(float64(N), 0))
		}
		X[k].num = result
		X[k].amp = cmplx.Abs(result)
		X[k].phase = cmplx.Phase(result)
	}
	return X
}
