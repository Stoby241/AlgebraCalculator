package V3

import (
	"fmt"
	"math"
)

type testCase struct {
	a [][]float64
	b []float64
}

var tc = testCase{
	a: [][]float64{
		{1.00, 0.00, 0.00, 0.00, 0.00, 0.00},
		{1.00, 0.63, 0.39, 0.25, 0.16, 0.10},
		{1.00, 1.26, 1.58, 1.98, 2.49, 3.13},
		{1.00, 1.88, 3.55, 6.70, 12.62, 23.80},
		{1.00, 2.51, 6.32, 15.88, 39.90, 100.28},
		{1.00, 3.14, 9.87, 31.01, 97.41, 306.02},
	},
	b: []float64{-0.01, 0.61, 0.91, 0.99, 0.60, 0.02},
}

var tc2 = testCase{
	a: [][]float64{
		{6, 1, 2},
		{5, 3, 3},
		{3, 2, 1},
	},
	b: []float64{1, 4, 14},
}

// result from above test case turns out to be correct to this tolerance.
const tolerance = 1e-14

func testGausianElimination() {
	x, err := gaussPartial(tc2.a, tc2.b)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(x)
}

func gaussPartial(a0 [][]float64, b0 []float64) ([]float64, error) {
	m := len(b0)
	a := make([][]float64, m)
	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	for k := range a {
		iMax := 0
		max := -1.
		for i := k; i < m; i++ {
			row := a[i]
			s := -1.
			for j := k; j < m; j++ {
				x := math.Abs(row[j])
				if x > s {
					s = x
				}
			}
			if abs := math.Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		if a[iMax][k] == 0 {
			return nil, newError(errorTypSolving, errorCriticalLevelPartial, "Gaussian Singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
	}
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}