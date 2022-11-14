package week1

import "testing"

func TestPlusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name     string
		args     args
		wantPos  string
		wantNeg  string
		wantZero string
	}{
		{
			name:     "example",
			args:     args{arr: []int32{1, 1, 0, -1, -1}},
			wantPos:  "0.400000",
			wantNeg:  "0.400000",
			wantZero: "0.200000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPos, gotNeg, gotZero := PlusMinus(tt.args.arr)
			if gotPos != tt.wantPos {
				t.Errorf("PlusMinus() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
			if gotNeg != tt.wantNeg {
				t.Errorf("PlusMinus() gotNeg = %v, want %v", gotNeg, tt.wantNeg)
			}
			if gotZero != tt.wantZero {
				t.Errorf("PlusMinus() gotZero = %v, want %v", gotZero, tt.wantZero)
			}
		})
	}
}
