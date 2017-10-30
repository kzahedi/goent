package discrete_test

import (
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestCreate2D(t *testing.T) {
	a := discrete.Create2D(20, 30)

	if len(a) != 20 {
		t.Errorf("1st dimension should be 20")
	}
	if len(a[0]) != 30 {
		t.Errorf("2nd dimension should be 30")
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			if a[i][j] != 0.0 {
				t.Errorf("a[%d][%d] should be 0.0 but is %f", i, j, a[i][j])
			}
		}
	}
}

func TestCreate3D(t *testing.T) {
	a := discrete.Create3D(20, 30, 40)

	if len(a) != 20 {
		t.Errorf("1st dimension should be 20")
	}
	if len(a[0]) != 30 {
		t.Errorf("2nd dimension should be 30")
	}
	if len(a[0][0]) != 40 {
		t.Errorf("3rd dimension should be 40")
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			for k := 0; k < 40; k++ {
				if a[i][j][k] != 0.0 {
					t.Errorf("a[%d][%d][%d] should be 0.0 but is %f", i, j, k, a[i][j][k])
				}
			}
		}
	}
}

func TestCreate2DInt(t *testing.T) {
	a := discrete.Create2DInt(20, 30)

	if len(a) != 20 {
		t.Errorf("1st dimension should be 20")
	}
	if len(a[0]) != 30 {
		t.Errorf("2nd dimension should be 30")
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			if a[i][j] != 0 {
				t.Errorf("a[%d][%d] should be 0.0 but is %f", i, j, a[i][j])
			}
		}
	}
}

func TestCreate3DInt(t *testing.T) {
	a := discrete.Create3DInt(20, 30, 40)

	if len(a) != 20 {
		t.Errorf("1st dimension should be 20")
	}
	if len(a[0]) != 30 {
		t.Errorf("2nd dimension should be 30")
	}
	if len(a[0][0]) != 40 {
		t.Errorf("3rd dimension should be 40")
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			for k := 0; k < 40; k++ {
				if a[i][j][k] != 0 {
					t.Errorf("a[%d][%d][%d] should be 0.0 but is %f", i, j, k, a[i][j][k])
				}
			}
		}
	}
}
