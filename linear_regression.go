package main

import  (
	"fmt"
	// "time"
	// "math"
	// "math/rand"
)

func mean(array[] float64) (mean float64){
	
	for i := 0; i < len(array); i++ {
		mean += array[i]
	}
	fmt.Println(len(array))
	mean /= float64(len(array))
	return
}

func linear_regression(X, Y []float64) {
	mean_x := mean(X)
	mean_y := mean(Y)
	m_num := 0.0
	m_denom := 0.0
	c := 0.0

	for i := 0; i < len(X); i++ {
		m_num += ((X[i] - mean_x) * (Y[i] - mean_y))
		m_denom += (X[i] - mean_x) * (X[i] - mean_x)
	}

	m:= m_num/m_denom
	c = mean_y - (m * mean_x)

	fmt.Println("Linear model: y =", m, "* x +", c)
}

func main() {
	X := [] float64 {1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	Y := [] float64 {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	linear_regression(X, Y)
}