package discrete_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

const precision = 0.0001

func random3D1() [][][]float64 {
	return [][][]float64{{{0.09, 0.1}, {0.12, 0.18}}, {{0.11, 0.09}, {0.14, 0.17}}}
}

func random3D2() [][][]float64 {
	return [][][]float64{{{0.00, 0.03}, {0.22, 0.06}}, {{0.15, 0.14}, {0.21, 0.19}}}
}

func random2D() [][]float64 {
	return [][]float64{{0.39, 0.03}, {0.51, 0.07}}
}

func random1D() []float64 {
	return []float64{0.80, 0.20}
}

func check3D(r, s [][][]float64, label string, t *testing.T) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if math.Abs(r[i][j][k]-s[i][j][k]) > precision {
					t.Errorf(fmt.Sprintf("%s should be %f, but it is %f", label, s[i][j][k], r[i][j][k]))
				}
			}
		}
	}
}

func check2D(r, s [][]float64, label string, t *testing.T) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if math.Abs(r[i][j]-s[i][j]) > precision {
				t.Errorf(fmt.Sprintf("%s should be %f, but it is %f", label, s[i][j], r[i][j]))
			}
		}
	}
}

func check1D(r, s []float64, label string, t *testing.T) {
	for i := 0; i < 2; i++ {
		if math.Abs(r[i]-s[i]) > precision {
			t.Errorf(fmt.Sprintf("%s should be %f, but it is %f", label, s[i], r[i]))
		}
	}
}

func TestPX1(t *testing.T) {
	p := random3D1()
	r := discrete.PX(p)
	s := []float64{0.49, 0.51}
	check1D(r, s, "PX 1", t)
}

func TestPX2(t *testing.T) {
	p := random3D2()
	r := discrete.PX(p)
	s := []float64{0.31, 0.69}
	check1D(r, s, "PX 2", t)
}

func TestPY1(t *testing.T) {
	p := random3D1()
	r := discrete.PY(p)
	s := []float64{0.39, 0.61}
	check1D(r, s, "PY 1", t)
}

func TestPY2(t *testing.T) {
	p := random3D2()
	r := discrete.PY(p)
	s := []float64{0.32, 0.68}
	check1D(r, s, "PY 2", t)
}

func TestPZ1(t *testing.T) {
	p := random3D1()
	r := discrete.PZ(p)
	s := []float64{0.46, 0.54}
	check1D(r, s, "PZ 1", t)
}

func TestPZ2(t *testing.T) {
	p := random3D2()
	r := discrete.PZ(p)
	s := []float64{0.58, 0.42}
	check1D(r, s, "PZ 2", t)
}

func TestPYZ(t *testing.T) {
	p := random3D1()
	r := discrete.PYZ(p)
	s := [][]float64{{0.2, 0.19}, {0.26, 0.35}}
	check2D(r, s, "PYZ", t)
}

func TestPXZ(t *testing.T) {
	p := random3D1()
	r := discrete.PXZ(p)
	s := [][]float64{{0.21, 0.28}, {0.25, 0.26}}
	check2D(r, s, "PXZ", t)
}

func TestPXY(t *testing.T) {
	p := random3D1()
	r := discrete.PXY(p)
	s := [][]float64{{0.19, 0.3}, {0.2, 0.31}}
	check2D(r, s, "PXY", t)
}

func TestH3(t *testing.T) {
	p := random3D1()
	r := discrete.H3(p)
	s := 2.95186
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("H3 should be %f, but it is %f", s, r))
	}
}

func TestH2(t *testing.T) {
	p := random2D()
	r := discrete.H2(p)
	s := 1.44555
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("H2 should be %f, but it is %f", s, r))
	}
}

func TestH1(t *testing.T) {
	p := random1D()
	r := discrete.H1(p)
	s := 0.721928

	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("H1 should be %f, but it is %f", s, r))
	}
}

func TestMiXvYgZ1(t *testing.T) {
	p := random3D1()
	s := 0.0000952709
	r := discrete.MiXvYgZ(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvYgZ should be %f, but it is %f", s, r))
	}
}

func TestMiXvYgZ2(t *testing.T) {
	p := random3D2()
	s := 0.0122117
	r := discrete.MiXvYgZ(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvYgZ should be %f, but it is %f", s, r))
	}
}

func TestPt(t *testing.T) {
	p := random3D1()

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

	r := discrete.Pt(p, -1.0, -1.0)
	s := [][][]float64{{{-0.91, 1.1}, {1.12, -0.82}}, {{-0.89, 1.09}, {1.14, -0.83}}}
	check3D(r, s, "Pt", t)

	r = discrete.Pt(p, 1.0, -1.0)
	s = [][][]float64{{{1.09, -0.9}, {-0.88, 1.18}}, {{-0.89, 1.09}, {1.14, -0.83}}}
	check3D(r, s, "Pt", t)

	r = discrete.Pt(p, -1.0, 1.0)
	s = [][][]float64{{{-0.91, 1.1}, {1.12, -0.82}}, {{1.11, -0.91}, {-0.86, 1.17}}}
	check3D(r, s, "Pt", t)

	r = discrete.Pt(p, 1.0, 1.0)
	s = [][][]float64{{{1.09, -0.9}, {-0.88, 1.18}}, {{1.11, -0.91}, {-0.86, 1.17}}}
	check3D(r, s, "Pt", t)
}

func TestMiXvZgY1(t *testing.T) {
	p := random3D1()
	s := 0.00283869
	r := discrete.MiXvZgY(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvZgY should be %f, but it is %f", s, r))
	}
}

func TestMiXvZgY2(t *testing.T) {
	p := random3D2()
	s := 0.0436263
	r := discrete.MiXvZgY(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvZgY should be %f, but it is %f", s, r))
	}
}

func TestMiXvYZ1(t *testing.T) {
	p := random3D1()
	s := 0.00283869
	r := discrete.MiXvYZ(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvYgZ should be %f, but it is %f", s, r))
	}
}

func TestMiXvYZ2(t *testing.T) {
	p := random3D2()
	s := 0.00283869
	r := discrete.MiXvYZ(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvYgZ should be %f, but it is %f", s, r))
	}
}

func TestMiXvY1(t *testing.T) {
	p := random3D1()
	s := 0.0000146819
	r := discrete.MiXvY(p)
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvY should be %f, but it is %f", s, r))
	}
}

func TestMiXvY2(t *testing.T) {
	p := random3D2()
	s := 0.0848927
	r := discrete.MiXvY(p)
	fmt.Println(discrete.H1(discrete.PX(p)), " ", discrete.H1(discrete.PY(p)), " ", discrete.H2(discrete.PXY(p)))
	if math.Abs(r-s) > precision {
		t.Errorf(fmt.Sprintf("MiXvY should be %f, but it is %f", s, r))
	}
}

// func TestID1(t *testing.T) {
// p := random31()
// a, b, c := discrete.InformationDecomposition(p, 100)
// sa := 0.0
// sb := 0.000913594
// sc := 0.00479433

// if math.Abs(a-sa) > precision {
// t.Errorf(fmt.Sprintf("Synergy should be %f, but it is %f", sa, a))
// }

// if math.Abs(b-sb) > precision {
// t.Errorf(fmt.Sprintf("UniqueXY should be %f, but it is %f", sb, b))
// }

// if math.Abs(c-sc) > precision {
// t.Errorf(fmt.Sprintf("UniqueXZ should be %f, but it is %f", sc, c))
// }
// }
