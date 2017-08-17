package goent_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/kzahedi/goent"
)

func TestMIasEntropies(t *testing.T) {
	t.Log("Testing Mutual Information as Entropy minus Conditional Entropy")
	rand.Seed(time.Now().UnixNano())
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	px := []float64{1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0}

	mi1 := goent.MutualInformation2(p1)
	ch1 := goent.ConditionalEntropy2(p1)
	h1 := goent.Entropy2(px)
	diff1 := mi1 - (h1 - ch1)

	if math.Abs(diff1) > 0.0001 {
		t.Errorf(" I(X;Y) = H(X) - H(X|Y) but the difference is %f, MI: %f, cH: %f, H:%f", math.Abs(diff1), mi1, ch1, h1)
	}

	p2 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	mi2 := goent.MutualInformation2(p2)  // I(X;Y) = H(X) - H(X|Y)
	ch2 := goent.ConditionalEntropy2(p2) // H(X|Y)
	h2 := goent.Entropy2(px)             // H(X)
	diff2 := mi2 - (h2 - ch2)

	if math.Abs(diff2) > 0.0001 {
		t.Errorf(" I(X;Y) = H(X) - H(X|Y) but the difference is %f, MI: %f, cH: %f, H:%f", math.Abs(diff2), mi2, ch2, h2)
	}
}

func TestCMIasMI(t *testing.T) {
	t.Log("Testing Conditional Mutual Information as Mutual Informations")
	pxyz := make([][][]float64, 5, 5)

	sum := 0.0
	for x := 0; x < 5; x++ {
		pxyz[x] = make([][]float64, 5, 5)
		for y := 0; y < 5; y++ {
			pxyz[x][y] = make([]float64, 5, 5)
		}
	}
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				v := rand.Float64()
				pxyz[x][y][z] += v
				sum += v
			}
		}
	}
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				pxyz[x][y][z] /= sum
			}
		}
	}

	pxz := make([][]float64, 5, 5)
	for x := 0; x < 5; x++ {
		pxz[x] = make([]float64, 5, 5)
	}
	sum = 0.0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				pxz[x][z] += pxyz[x][y][z]
				sum += pxyz[x][y][z]
			}
		}
	}

	if math.Abs(sum-1.0) > 0.0001 {
		t.Errorf("\\sum_{x,y,z} p(x,y,z) should be 1.0 but it is %f", sum)
	}

	px_yz := make([][]float64, 5, 5)
	for x := 0; x < 5; x++ {
		px_yz[x] = make([]float64, 25, 25)
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				px_yz[x][y*5+z] = pxyz[x][y][z]
			}
		}
	}

	cmi := goent.ConditionalMutualInformation2(pxyz)
	multi := goent.MutualInformation2(px_yz)
	mi := goent.MutualInformation2(pxz)
	diff := cmi - (multi - mi)

	if math.Abs(diff) > 0.0001 {
		t.Errorf("I(X;Y|Z) = I(X;Y,Z) - I(X;Z) but the difference is %f, I(X;Y|Z): %f, I(X;Y,Z): %f, I(X;Z):%f", math.Abs(diff), cmi, multi, mi)
	}

}
