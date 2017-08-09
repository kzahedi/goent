package goent_test

import (
	"testing"

	"github.com/kzahedi/goent"
)

func TestDiscretiseVector(t *testing.T) {
	t.Log("Testing DiscrestiseVector")
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

	d := goent.DiscrestiseVector(p, 10, 0.0, 1.0)

	if d[0] != 0 {
		t.Errorf("0.0  must be mapped to 0 and not ", p[0])
	}
	if d[1] != 1 {
		t.Errorf("0.1 must be mapped to 0 and not ", p[1])
	}
	if d[2] != 2 {
		t.Errorf("0.2 must be mapped to 0 and not ", p[2])
	}
	if d[3] != 3 {
		t.Errorf("0.3 must be mapped to 0 and not ", p[3])
	}
	if d[4] != 4 {
		t.Errorf("0.4 must be mapped to 0 and not ", p[4])
	}
	if d[5] != 5 {
		t.Errorf("0.5 must be mapped to 0 and not ", p[5])
	}
	if d[6] != 6 {
		t.Errorf("0.6 must be mapped to 0 and not ", p[6])
	}
	if d[7] != 7 {
		t.Errorf("0.7 must be mapped to 0 and not ", p[7])
	}
	if d[8] != 8 {
		t.Errorf("0.8 must be mapped to 0 and not ", p[8])
	}
	if d[9] != 9 {
		t.Errorf("0.9 must be mapped to 0 and not ", p[9])
	}
	if d[10] != 9 {
		t.Errorf("1.0 must be mapped to 0 and not ", p[10])
	}

}
