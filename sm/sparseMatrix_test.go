package sm

import (
	"reflect"
	"testing"
)

func TestSparseMatrixIndex_Equal(t *testing.T) {
	type args struct {
		other SparseMatrixIndex
	}
	tests := []struct {
		name string
		si   SparseMatrixIndex
		args args
		want bool
	}{
		{name: "Equal with 1 dimension",
			si:   SparseMatrixIndex{1},
			args: args{other: SparseMatrixIndex{1}},
			want: true},
		{name: "Fail with 1 dimension",
			si:   SparseMatrixIndex{2},
			args: args{other: SparseMatrixIndex{1}},
			want: false},
		{name: "Equal with 2 dimensions",
			si:   SparseMatrixIndex{1, 2},
			args: args{other: SparseMatrixIndex{1, 2}},
			want: true},
		{name: "Fail with 2 dimensions",
			si:   SparseMatrixIndex{1, 1},
			args: args{other: SparseMatrixIndex{1, 2}},
			want: false},
		{name: "Equal with 3 dimensions",
			si:   SparseMatrixIndex{1, 2, 3},
			args: args{other: SparseMatrixIndex{1, 2, 3}},
			want: true},
		{name: "Fail with 3 dimensions",
			si:   SparseMatrixIndex{1, 1, 3},
			args: args{other: SparseMatrixIndex{1, 2, 3}},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.si.Equal(tt.args.other); got != tt.want {
				t.Errorf("SparseMatrixIndex.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSparseMatrix_Get(t *testing.T) {
	type fields struct {
		Indices []SparseMatrixIndex
		Values  []float64
	}
	type args struct {
		smi SparseMatrixIndex
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{name: "Find index (1,2,3)",
			fields:  fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}}, Values: []float64{0.123}},
			args:    args{smi: SparseMatrixIndex{1, 2, 3}},
			want:    0.123,
			wantErr: false},
		{name: "Fail find index (1,2,3) (dimension)",
			fields:  fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2}}, Values: []float64{0.123}},
			args:    args{smi: SparseMatrixIndex{1, 2, 3}},
			want:    -1.0,
			wantErr: true},
		{name: "Find index (1,2,3) returning 0.0",
			fields:  fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 4}}, Values: []float64{0.124}},
			args:    args{smi: SparseMatrixIndex{1, 2, 3}},
			want:    0.0,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SparseMatrix{
				Indices: tt.fields.Indices,
				Values:  tt.fields.Values,
			}
			got, err := s.Get(tt.args.smi)
			if (err != nil) != tt.wantErr {
				t.Errorf("SparseMatrix.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SparseMatrix.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSparseMatrix_Set(t *testing.T) {
	type fields struct {
		Indices []SparseMatrixIndex
		Values  []float64
	}
	type args struct {
		smi   SparseMatrixIndex
		value float64
	}
	type want struct {
		sm SparseMatrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{name: "Set empty spot",
			fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{1.0, 2.0}},
			args: args{smi: SparseMatrixIndex{3, 4, 5}, value: 3.0},
			want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}, SparseMatrixIndex{3, 4, 5}},
				Values: []float64{1.0, 2.0, 3.0}}}},
		{name: "Set existing spot",
			fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{1.0, 2.0}},
			args: args{smi: SparseMatrixIndex{1, 2, 3}, value: 3.0},
			want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{3.0, 2.0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SparseMatrix{
				Indices: tt.fields.Indices,
				Values:  tt.fields.Values,
			}
			s.Set(tt.args.smi, tt.args.value)
			if s.Equal(tt.want.sm) == false {
				t.Errorf("SparseMatrix.Set() resulted in %v want %v", s, tt.want.sm)
			}
		})
	}
}

func TestSparseMatrix_Add(t *testing.T) {
	type fields struct {
		Indices []SparseMatrixIndex
		Values  []float64
	}
	type args struct {
		smi   SparseMatrixIndex
		value float64
	}
	type want struct {
		sm SparseMatrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		// {name: "Add empty spot",
		// 	fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
		// 		Values: []float64{1.0, 2.0}},
		// 	args: args{smi: SparseMatrixIndex{3, 4, 5}, value: 3.0},
		// 	want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}, SparseMatrixIndex{3, 4, 5}},
		// 		Values: []float64{1.0, 2.0, 3.0}}}},
		{name: "Add existing spot",
			fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{1.0, 2.0}},
			args: args{smi: SparseMatrixIndex{1, 2, 3}, value: 3.0},
			want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{4.0, 2.0}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SparseMatrix{
				Indices: tt.fields.Indices,
				Values:  tt.fields.Values,
			}
			s.Add(tt.args.smi, tt.args.value)
			if s.Equal(tt.want.sm) == false {
				t.Errorf("SparseMatrix.Add() resulted in %v want %v", s, tt.want.sm)
			}
		})
	}
}
func TestSparseMatrix_Mul(t *testing.T) {
	type fields struct {
		Indices []SparseMatrixIndex
		Values  []float64
	}
	type args struct {
		smi   SparseMatrixIndex
		value float64
	}
	type want struct {
		sm SparseMatrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{name: "Mul empty spot",
			fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{2.0, 3.0}},
			args: args{smi: SparseMatrixIndex{3, 4, 5}, value: 3.0},
			want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{2.0, 3.0}}}},
		{name: "Mul existing spot",
			fields: fields{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{2.0, 1.0}},
			args: args{smi: SparseMatrixIndex{1, 2, 3}, value: 3.0},
			want: want{sm: SparseMatrix{Indices: []SparseMatrixIndex{SparseMatrixIndex{1, 2, 3}, SparseMatrixIndex{2, 3, 4}},
				Values: []float64{6.0, 1.0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SparseMatrix{
				Indices: tt.fields.Indices,
				Values:  tt.fields.Values,
			}
			s.Mul(tt.args.smi, tt.args.value)
			if s.Equal(tt.want.sm) == false {
				t.Errorf("SparseMatrix.Add() resulted in %v want %v", s, tt.want.sm)
			}
		})
	}
}

func TestSparseMatrix_Scale(t *testing.T) {
	type fields struct {
		Indices []SparseMatrixIndex
		Values  []float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SparseMatrix{
				Indices: tt.fields.Indices,
				Values:  tt.fields.Values,
			}
			s.Scale(tt.args.value)
		})
	}
}

func TestCreateSparseMatrix(t *testing.T) {
	tests := []struct {
		name string
		want SparseMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSparseMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSparseMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
