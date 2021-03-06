package discrete

import (
	"math"
)

func PX(pxyz [][][]float64) []float64 {
	r := make([]float64, 2, 2)
	r[0] = pxyz[0][0][0] + pxyz[0][0][1] + pxyz[0][1][0] + pxyz[0][1][1]
	r[1] = pxyz[1][0][0] + pxyz[1][0][1] + pxyz[1][1][0] + pxyz[1][1][1]
	return r
}

func PY(pxyz [][][]float64) []float64 {
	r := make([]float64, 2, 2)
	r[0] = pxyz[0][0][0] + pxyz[0][0][1] + pxyz[1][0][0] + pxyz[1][0][1]
	r[1] = pxyz[0][1][0] + pxyz[0][1][1] + pxyz[1][1][0] + pxyz[1][1][1]
	return r
}

func PZ(pxyz [][][]float64) []float64 {
	r := make([]float64, 2, 2)
	r[0] = pxyz[0][0][0] + pxyz[0][1][0] + pxyz[1][0][0] + pxyz[1][1][0]
	r[1] = pxyz[0][0][1] + pxyz[0][1][1] + pxyz[1][0][1] + pxyz[1][1][1]
	return r
}

func PYZ(pxyz [][][]float64) [][]float64 {
	r := make([][]float64, 2, 2)
	r[0] = make([]float64, 2, 2)
	r[1] = make([]float64, 2, 2)
	r[0][0] = pxyz[0][0][0] + pxyz[1][0][0]
	r[0][1] = pxyz[0][0][1] + pxyz[1][0][1]
	r[1][0] = pxyz[0][1][0] + pxyz[1][1][0]
	r[1][1] = pxyz[0][1][1] + pxyz[1][1][1]
	return r
}

func PXZ(pxyz [][][]float64) [][]float64 {
	r := make([][]float64, 2, 2)
	r[0] = make([]float64, 2, 2)
	r[1] = make([]float64, 2, 2)
	r[0][0] = pxyz[0][0][0] + pxyz[0][1][0]
	r[0][1] = pxyz[0][0][1] + pxyz[0][1][1]
	r[1][0] = pxyz[1][0][0] + pxyz[1][1][0]
	r[1][1] = pxyz[1][0][1] + pxyz[1][1][1]
	return r
}

func PXY(pxyz [][][]float64) [][]float64 {
	r := make([][]float64, 2, 2)
	r[0] = make([]float64, 2, 2)
	r[1] = make([]float64, 2, 2)
	r[0][0] = pxyz[0][0][0] + pxyz[0][0][1]
	r[0][1] = pxyz[0][1][0] + pxyz[0][1][1]
	r[1][0] = pxyz[1][0][0] + pxyz[1][0][1]
	r[1][1] = pxyz[1][1][0] + pxyz[1][1][1]
	return r
}

func H3(pxyz [][][]float64) float64 {
	r := 0.0

	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				r -= pxyz[x][y][z] * math.Log2(pxyz[x][y][z])
			}
		}
	}
	return r
}

func H2(pxy [][]float64) float64 {
	r := 0.0

	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			r -= pxy[x][y] * math.Log2(pxy[x][y])
		}
	}
	return r
}

func H1(px []float64) float64 {
	r := 0.0

	for x := 0; x < 2; x++ {
		r -= px[x] * math.Log2(px[x])
	}
	return r
}

func Pt(pxyz [][][]float64, a, b float64) [][][]float64 {
	A := make([][][]float64, 2, 2)
	A[0] = make([][]float64, 2, 2)
	A[1] = make([][]float64, 2, 2)
	A[0][0] = make([]float64, 2, 2)
	A[0][1] = make([]float64, 2, 2)
	A[1][0] = make([]float64, 2, 2)
	A[1][1] = make([]float64, 2, 2)

	B := make([][][]float64, 2, 2)
	B[0] = make([][]float64, 2, 2)
	B[1] = make([][]float64, 2, 2)
	B[0][0] = make([]float64, 2, 2)
	B[0][1] = make([]float64, 2, 2)
	B[1][0] = make([]float64, 2, 2)
	B[1][1] = make([]float64, 2, 2)

	A[0][0][0] = 1.0
	A[0][0][1] = -1.0
	A[0][1][0] = -1.0
	A[0][1][1] = 1.0

	B[1][0][0] = 1.0
	B[1][0][1] = -1.0
	B[1][1][0] = -1.0
	B[1][1][1] = 1.0

	r := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		r[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			r[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				r[i][j][k] = pxyz[i][j][k] + a*A[i][j][k] + b*B[i][j][k]
			}
		}
	}
	return r
}

func MiXvYgZ(pxyz [][][]float64) float64 {
	return H2(PXZ(pxyz)) + H2(PYZ(pxyz)) - H3(pxyz) - H1(PZ(pxyz))
}

func MiXvZgY(pxyz [][][]float64) float64 {
	return H2(PXY(pxyz)) + H2(PYZ(pxyz)) - H3(pxyz) - H1(PY(pxyz))
}

func MiXvY(pxyz [][][]float64) float64 {
	return H1(PX(pxyz)) + H1(PY(pxyz)) - H2(PXY(pxyz))
}

func MiXvYZ(pxyz [][][]float64) float64 {
	return H1(PX(pxyz)) + H2(PYZ(pxyz)) - H3(pxyz)
}

func CoI(pxyz [][][]float64) float64 {
	return MiXvY(pxyz) - MiXvYgZ(pxyz)
}

func MinMax(pxyz [][][]float64, resolution int) (float64, float64, float64, float64, float64, float64) {
	amin := math.Max(-pxyz[0][0][0], -pxyz[0][1][1])
	amax := math.Min(pxyz[0][0][1], pxyz[0][1][0])
	adelta := (amax - amin) / float64(resolution)
	bmin := math.Max(-pxyz[1][0][0], -pxyz[1][1][1])
	bmax := math.Min(pxyz[1][0][1], pxyz[1][1][0])
	bdelta := (bmax - bmin) / float64(resolution)
	return amin, amax, adelta, bmin, bmax, bdelta
}

// InformationDecomposition return the UI(X;Y\Z), UI(X;Z\Y), CI(X;Y,Z), and SI(X;Y,Z)
// according to
// N. Bertschinger, J. Rauh, E. Olbrich, J. Jost, and N. Ay, Quantifying unique information, CoRR, 2013
func InformationDecomposition(pxyz [][][]float64, resolution int) (float64, float64, float64) {

	amin, amax, adelta, bmin, bmax, bdelta := MinMax(pxyz, resolution)

	uniqueXY := MiXvYgZ(Pt(pxyz, 0.0, 0.0))
	uniqueXZ := MiXvZgY(Pt(pxyz, 0.0, 0.0))
	maxCoI := CoI(Pt(pxyz, 0.0, 0.0))

	r := 0.0
	for a := amin; a <= amax; a += adelta {
		for b := bmin; b <= bmax; b += bdelta {
			q := Pt(pxyz, a, b)

			r = CoI(q)
			if r > maxCoI {
				maxCoI = r
			}

			r = MiXvZgY(q)
			if r < uniqueXY {
				uniqueXY = r
			}

			r = MiXvYgZ(q)
			if r < uniqueXZ {
				uniqueXZ = r
			}
		}
	}

	coI := maxCoI - CoI(pxyz)

	return coI, uniqueXY, uniqueXZ
}
