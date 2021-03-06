package dh_test

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"

	"encoding/csv"

	"github.com/kzahedi/goent/dh"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestDiscretiseVector(t *testing.T) {
	t.Log("Testing DiscretiseVector")
	p := make([]float64, 11, 11)

	p[0] = 0.0
	p[1] = 0.1
	p[2] = 0.2
	p[3] = 0.3
	p[4] = 0.4
	p[5] = 0.5
	p[6] = 0.6
	p[7] = 0.7
	p[8] = 0.8
	p[9] = 0.9
	p[10] = 1.0

	d := dh.DiscretiseVector(p, 10, 0.0, 1.0)

	if d[0] != 0 {
		t.Errorf("0.0  must be mapped to 0 and not %f", p[0])
	}
	if d[1] != 1 {
		t.Errorf("0.1 must be mapped to 0 and not %f", p[1])
	}
	if d[2] != 2 {
		t.Errorf("0.2 must be mapped to 0 and not %f", p[2])
	}
	if d[3] != 3 {
		t.Errorf("0.3 must be mapped to 0 and not %f", p[3])
	}
	if d[4] != 4 {
		t.Errorf("0.4 must be mapped to 0 and not %f", p[4])
	}
	if d[5] != 5 {
		t.Errorf("0.5 must be mapped to 0 and not %f", p[5])
	}
	if d[6] != 6 {
		t.Errorf("0.6 must be mapped to 0 and not %f", p[6])
	}
	if d[7] != 7 {
		t.Errorf("0.7 must be mapped to 0 and not %f", p[7])
	}
	if d[8] != 8 {
		t.Errorf("0.8 must be mapped to 0 and not %f", p[8])
	}
	if d[9] != 9 {
		t.Errorf("0.9 must be mapped to 0 and not %f", p[9])
	}
	if d[10] != 9 {
		t.Errorf("1.0 must be mapped to 0 and not %f", p[10])
	}

}

func TestDiscretise(t *testing.T) {
	t.Log("Testing Discretise")

	p := [][]float64{
		{0.0, 0.1, 0.2},
		{0.1, 0.2, 0.3},
		{0.2, 0.3, 0.4},
		{0.3, 0.4, 0.5},
		{0.4, 0.5, 0.6},
		{0.5, 0.6, 0.7},
		{0.6, 0.7, 0.8},
		{0.7, 0.8, 0.9},
		{0.8, 0.9, 1.0}}

	d := dh.Discretise(p,
		[]int{10, 10, 10},
		[]float64{0.0, 0.0, 0.0},
		[]float64{1.0, 1.0, 1.0})

	for i := 0; i < 9; i++ {
		if d[i][0] != i {
			t.Errorf("%f must be mapped to %d and not %d", p[i][0], i, d[i][0])
		}

		if d[i][1] != i+1 {
			t.Errorf("%f must be mapped to %d and not %d", p[i][1], i+1, d[i][1])
		}

		if d[i][2] != min(9, i+2) {
			t.Errorf("%f must be mapped to %d and not %d", p[i][2], i+2, d[i][2])
		}
	}
}

func TestMakeUnivariate(t *testing.T) {
	t.Log("Testing MakeUnivariate")

	p := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5, 6},
		{5, 6, 7},
		{6, 7, 8},
		{7, 8, 9}}

	d := dh.MakeUnivariate(p, []int{10, 10, 10})

	if d[0] != 1+10*2+100*3 {
		t.Errorf("%d,%d,%d must be mapped to %d and not %d",
			p[0][0], p[0][1], p[0][2],
			(1 + 10*2 + 100*3),
			d[0])
	}

	if d[1] != 2+10*3+100*4 {
		t.Errorf("%d,%d,%d must be mapped to %d and not %d",
			p[1][0], p[1][1], p[1][2],
			(2 + 10*3 + 100*4),
			d[1])
	}

}

func TestRelabel(t *testing.T) {
	t.Log("Testing Relabel")

	p := []int{10, 1, 4, 13, 871, 283, 123, 987, 2415, 88, 57, 10, 283, 987}
	q := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 5, 7}

	d := dh.Relabel(p)

	for i := range p {
		if d[i] != q[i] {
			t.Errorf("%d must be mapped to %d and not %d", p[i], q[i], d[i])
		}
	}
}

func TestExtractColumns(t *testing.T) {
	t.Log("Testing Relabel")

	d := [][]float64{
		{1.0, 10.0, 100.0, 1000.0},
		{2.0, 20.0, 200.0, 2000.0},
		{3.0, 30.0, 300.0, 3000.0},
		{4.0, 40.0, 400.0, 4000.0},
		{5.0, 50.0, 500.0, 5000.0},
		{6.0, 60.0, 600.0, 6000.0},
		{7.0, 70.0, 700.0, 7000.0},
		{8.0, 80.0, 800.0, 8000.0},
		{9.0, 90.0, 900.0, 9000.0},
	}

	c1 := dh.ExtractColumns(d, []int{1})

	for i := 0; i < 9; i++ {
		if int(c1[i][0]) != 10*(i+1) {
			t.Errorf("Values should be %d but it is %d", 10*(i+1), int(c1[i][0]))
		}
	}

	c2 := dh.ExtractColumns(d, []int{0, 3})

	for i := 0; i < 9; i++ {
		if int(c2[i][0]) != i+1 {
			t.Errorf("Values should be %d but it is %d", (i + 1), int(c2[i][0]))
		}
		if int(c2[i][1]) != (i+1)*1000 {
			t.Errorf("Values should be %d but it is %d", (1000 * (i + 1)), int(c2[i][1]))
		}
	}
}

func TestReadData(t *testing.T) {
	data := make([][]float64, 100, 100)
	rand.Seed(42)
	for i := 0; i < 100; i++ {
		data[i] = make([]float64, 100, 100)
		for j := 0; j < 100; j++ {
			data[i][j] = rand.Float64()
		}
	}

	strdata := make([][]string, 102, 102)
	for i := 0; i < 102; i++ {
		strdata[i] = make([]string, 100, 100)
	}
	strdata[0][0] = "# header"
	strdata[0][1] = "second column header"
	strdata[0][2] = "second column header"
	strdata[1][0] = "# some"
	strdata[1][1] = "comment"
	strdata[1][2] = "line"

	for i := 2; i < 102; i++ {
		for j := 0; j < 100; j++ {
			strdata[i][j] = fmt.Sprintf("%.8f", data[i-2][j])
		}
	}

	f, err := os.Create("/tmp/test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, value := range strdata {
		writer.Write(value)
	}

	writer.Flush()

	rdata := dh.ReadData("/tmp/test.csv")
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if math.Abs(data[i][j]-rdata[i][j]) > 0.00001 {
				t.Errorf("Values don't match %f != %f (%d,%d)", data[i][j], rdata[i][j], i, j)
			}
		}
	}
}

func TestGetMinMax(t *testing.T) {
	data := make([][]float64, 10, 10)
	index := 0.0
	for i := 0; i < 10; i++ {
		data[i] = make([]float64, 5, 5)
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			data[i][j] = index
			index++
		}
	}

	min, max := dh.GetMinMax(data)

	if min[0] != 0.0 {
		t.Errorf("min[0] should be 0.0 but it is %f", min[0])
	}
	if min[1] != 10.0 {
		t.Errorf("min[1] should be 10.0 but it is %f", min[1])
	}
	if min[2] != 20.0 {
		t.Errorf("min[2] should be 20.0 but it is %f", min[2])
	}
	if min[3] != 30.0 {
		t.Errorf("min[3] should be 30.0 but it is %f", min[3])
	}
	if min[4] != 40.0 {
		t.Errorf("min[4] should be 40.0 but it is %f", min[4])
	}

	if max[0] != 9.0 {
		t.Errorf("max[0] should be 9.0 but it is %f", max[0])
	}
	if max[1] != 19.0 {
		t.Errorf("max[1] should be 19.0 but it is %f", max[1])
	}
	if max[2] != 29.0 {
		t.Errorf("max[2] should be 29.0 but it is %f", max[2])
	}
	if max[3] != 39.0 {
		t.Errorf("max[3] should be 39.0 but it is %f", max[3])
	}
	if max[4] != 49.0 {
		t.Errorf("max[4] should be 49.0 but it is %f", max[4])
	}
}
