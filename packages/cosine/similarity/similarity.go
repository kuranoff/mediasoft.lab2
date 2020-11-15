package similarity

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// PowVec ...
func PowVec(v []float64) []float64 {
	var pv64 = []float64{}
	for _, i := range v {
		j := math.Pow(i, 2)
		pv64 = append(pv64, j)
	}
	return pv64
}

// Calc ...
func Calc(w1 []float64, w2 []float64) float64 {

	u := mat.NewVecDense(100, w1)
	v := mat.NewVecDense(100, w2)

	nominator := mat.Dot(u, v)

	ux := mat.NewVecDense(100, PowVec(w1))
	vx := mat.NewVecDense(100, PowVec(w2))

	uxsum := mat.Sum(ux)
	uxnorm := math.Sqrt(uxsum)

	vxsum := mat.Sum(vx)
	vxnorm := math.Sqrt(vxsum)

	denominatorx := uxnorm * vxnorm

	cosinesimilarityx := nominator / denominatorx

	return cosinesimilarityx

}
