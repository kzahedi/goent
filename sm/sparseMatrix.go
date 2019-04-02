package sm

import (
	"fmt"
	"math"
)

// SparseMatrixIndex is the index of the tensor, e.g. index can be (0,1) or (0,1,23),
// referencing a cell in a 2d or 3d matrix
type SparseMatrixIndex []int

// SparseMatrix is a sparse matrix implementation
type SparseMatrix struct {
	Indices []SparseMatrixIndex
	Values  []float64
}

// Equal return true of the indices are equal
func (si SparseMatrixIndex) Equal(other SparseMatrixIndex) bool {
	if len(si) != len(other) {
		return false
	}
	for i := 0; i < len(si); i++ {
		if si[i] != other[i] {
			return false
		}
	}
	return true
}

// Get returns the value specified by the index, or -1.0 and an error otherwise
func (s SparseMatrix) Get(smi SparseMatrixIndex) (float64, error) {
	for i, v := range s.Indices {
		if v.Equal(smi) {
			return s.Values[i], nil
		}
		if len(s.Indices[i]) != len(smi) {
			return -1.0, fmt.Errorf("Sparse Matrix: Dimension mismatch %d vs %d", len(s.Indices[i]), len(smi))
		}
	}
	return 0.0, nil
}

// Set set the value specified by the index
func (s *SparseMatrix) Set(smi SparseMatrixIndex, value float64) {
	for i, v := range s.Indices {
		if v.Equal(smi) {
			s.Values[i] = value
			return
		}
	}
	s.Indices = append(s.Indices, smi)
	s.Values = append(s.Values, value)
}

// Add add the value to the value specified by the index
func (s *SparseMatrix) Add(smi SparseMatrixIndex, value float64) {
	for i, v := range s.Indices {
		if v.Equal(smi) {
			s.Values[i] += value
			return
		}
	}
	// will only be called if the value was not previously present
	s.Indices = append(s.Indices, smi)
	s.Values = append(s.Values, value)
}

// Mul multiplies the value to the value specified by the index
func (s SparseMatrix) Mul(smi SparseMatrixIndex, value float64) {
	for i, v := range s.Indices {
		if v.Equal(smi) {
			s.Values[i] *= value
			return
		}
	}
	// else 0 * value = 0, nothing will change
}

// Scale multiplies the value to all cells
func (s SparseMatrix) Scale(value float64) {
	for i := range s.Values {
		s.Values[i] *= value
	}
}

// Equal returns true if equal, false if not equal
func (s SparseMatrix) Equal(other SparseMatrix) bool {
	if len(s.Indices) != len(other.Indices) {
		return false
	}
	if len(s.Values) != len(other.Values) {
		return false
	}
	for i := range s.Indices {
		if len(s.Indices[i]) != len(other.Indices[i]) {
			return false
		}
		if math.Abs(s.Values[i]-other.Values[i]) > 0.000001 {
			return false
		}
		for j := range s.Indices[i] {
			if s.Indices[i][j] != other.Indices[i][j] {
				return false
			}
		}
	}
	return true
}

// CreateSparseMatrix returns an empty sparse matrix
func CreateSparseMatrix() SparseMatrix {
	return SparseMatrix{Indices: []SparseMatrixIndex{}, Values: []float64{}}
}
