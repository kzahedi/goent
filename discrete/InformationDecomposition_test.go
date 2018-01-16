package discrete_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func randomP() [][][]float64 {
	s := 0.0
	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				p := rand.Float64()
				d[i][j][k] = p
				s += p
			}
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				d[i][j][k] /= s
			}
		}
	}

	return d
}

func TestPX(t *testing.T) {
	s := []float64{664.0, 444.0}

	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				d[i][j][k] = float64((i + 1) + (j+1)*10 + (k+1)*100)
			}
		}
	}

	r := discrete.PX(d)

	for i := 0; i < 2; i++ {
		if math.Abs(r[i]-s[i]) > 0.0001 {
			t.Errorf(fmt.Sprintf("r[%d] should be %f but is %f", i, s[i], r[i]))
		}
	}
}

func TestPY(t *testing.T) {
	s := []float64{646.0, 686.0}

	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				d[i][j][k] = float64((i + 1) + (j+1)*10 + (k+1)*100)
			}
		}
	}

	r := discrete.PY(d)

	for i := 0; i < 2; i++ {
		if math.Abs(r[i]-s[i]) > 0.0001 {
			t.Errorf(fmt.Sprintf("r[%d] should be %f but is %f", i, s[i], r[i]))
		}
	}
}

func TestPZ(t *testing.T) {
	s := []float64{466.0, 866.0}

	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				d[i][j][k] = float64((i + 1) + (j+1)*10 + (k+1)*100)
			}
		}
	}

	r := discrete.PZ(d)

	for i := 0; i < 2; i++ {
		if math.Abs(r[i]-s[i]) > 0.0001 {
			t.Errorf(fmt.Sprintf("r[%d] should be %f but is %f", i, s[i], r[i]))
		}
	}
}

func TestPYZ(t *testing.T) {
	s := [][]float64{{223.0, 423.0}, {243.0, 443.0}}

	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				d[i][j][k] = float64((i + 1) + (j+1)*10 + (k+1)*100)
			}
		}
	}

	r := discrete.PYZ(d)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if math.Abs(r[i][j]-s[i][j]) > 0.0001 {
				t.Errorf(fmt.Sprintf("r[%d][%d] should be %f but is %f", i, j, s[i][j], r[i][j]))
			}
		}
	}
}

func TestPXZ(t *testing.T) {
	s := [][]float64{{223.0, 423.0}, {234.0, 434.0}}

	d := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			d[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				d[i][j][k] = float64((i + 1) + (j+1)*10 + (k+1)*100)
			}
		}
	}

	r := discrete.PXZ(d)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if math.Abs(r[i][j]-s[i][j]) > 0.0001 {
				t.Errorf(fmt.Sprintf("r[%d][%d] should be %f but is %f", i, j, s[i][j], r[i][j]))
			}
		}
	}
}

func TestH3(t *testing.T) {
	p1 := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		p1[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			p1[i][j] = make([]float64, 2, 2)
		}
	}
	p1[0][0][0] = 1.0

	p2 := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		p2[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			p2[i][j] = make([]float64, 2, 2)
			for k := 0; k < 2; k++ {
				p2[i][j][k] = 1.0 / 8.0
			}
		}
	}

	r1 := discrete.H3(p1)
	r2 := discrete.H3(p2)

	if math.Abs(r1) > 0.0001 {
		t.Errorf(fmt.Sprintf("H3 of uniform should be %f, but it is %f", 0.0, r1))
	}
	if math.Abs(r2-math.Log2(8)) > 0.0001 {
		t.Errorf(fmt.Sprintf("H3 of uniform should be %f, but it is %f", math.Log2(8), r2))
	}
}

func TestH2(t *testing.T) {
	p1 := make([][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		p1[i] = make([]float64, 2, 2)
	}
	p1[0][0] = 1.0

	p2 := make([][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		p2[i] = make([]float64, 2, 2)
		for j := 0; j < 2; j++ {
			p2[i][j] = 1.0 / 4.0
		}
	}

	r1 := discrete.H2(p1)
	r2 := discrete.H2(p2)

	if math.Abs(r1) > 0.0001 {
		t.Errorf(fmt.Sprintf("H2 of uniform should be %f, but it is %f", 0.0, r1))
	}
	if math.Abs(r2-math.Log2(4)) > 0.0001 {
		t.Errorf(fmt.Sprintf("H2 of uniform should be %f, but it is %f", math.Log2(4), r2))
	}
}

func TestH1(t *testing.T) {
	p1 := make([]float64, 2, 2)
	p1[0] = 1.0

	p2 := make([]float64, 2, 2)
	p2[0] = 1.0 / 2.0
	p2[1] = 1.0 / 2.0

	r1 := discrete.H1(p1)
	r2 := discrete.H1(p2)

	if math.Abs(r1) > 0.0001 {
		t.Errorf(fmt.Sprintf("H1 of uniform should be %f, but it is %f", 0.0, r1))
	}
	if math.Abs(r2-math.Log2(2)) > 0.0001 {
		t.Errorf(fmt.Sprintf("H1 of uniform should be %f, but it is %f", math.Log2(2), r2))
	}
}

func TestMiXvYgZ(t *testing.T) {
	p := randomP()
	s := discrete.H2(discrete.PXZ(p)) + discrete.H2(discrete.PYZ(p)) - discrete.H3(p) - discrete.H1(discrete.PZ(p))
	r := discrete.MiXvYgZ(p)
	if math.Abs(r-s) > 0.0001 {
		t.Errorf(fmt.Sprintf("MiXvYgZ should be %f, but it is %f", s, r))
	}
}

func TestMiXvZgY(t *testing.T) {
	p := randomP()
	s := discrete.H2(discrete.PXY(p)) + discrete.H2(discrete.PYZ(p)) - discrete.H3(p) - discrete.H1(discrete.PY(p))
	r := discrete.MiXvZgY(p)
	if math.Abs(r-s) > 0.0001 {
		t.Errorf(fmt.Sprintf("MiXvZgY should be %f, but it is %f", s, r))
	}
}

func TestMiXvY(t *testing.T) {
	p := randomP()
	s := discrete.H1(discrete.PX(p)) + discrete.H1(discrete.PY(p)) - discrete.H2(discrete.PXY(p))
	r := discrete.MiXvY(p)
	if math.Abs(r-s) > 0.0001 {
		t.Errorf(fmt.Sprintf("MiXvY should be %f, but it is %f", s, r))
	}
}

func TestCoI(t *testing.T) {
	p := randomP()
	s := discrete.MiXvY(p) - discrete.MiXvYgZ(p)
	r := discrete.CoI(p)
	if math.Abs(r-s) > 0.0001 {
		t.Errorf(fmt.Sprintf("CoI should be %f, but it is %f", s, r))
	}
}

func TestPt(t *testing.T) {
	p := make([][][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		p[i] = make([][]float64, 2, 2)
		for j := 0; j < 2; j++ {
			p[i][j] = make([]float64, 2, 2)
		}
	}

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

	for i := 0; i < 100; i++ {
		a := 2.0*rand.Float64() - 1
		b := 2.0*rand.Float64() - 1
		q := discrete.Pt(p, a, b)
		for x := 0; x < 2; x++ {
			for y := 0; y < 2; y++ {
				for z := 0; z < 2; z++ {
					v := q[x][y][z]
					w := a*A[x][y][z] + b*B[x][y][z]
					if math.Abs(v-w) > 0.0001 {
						t.Errorf(fmt.Sprintf("Pt should be %f, but it is %f", w, v))
					}
				}
			}
		}
	}
}
