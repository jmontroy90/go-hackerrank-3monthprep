package week1

import "testing"

func TestMinMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name    string
		args    args
		wantMin int64
		wantMax int64
	}{
		{
			name:    "example",
			args:    args{arr: []int32{1, 3, 5, 7, 9}},
			wantMin: 16,
			wantMax: 24,
		},
		{
			name:    "example-reversed",
			args:    args{arr: []int32{9, 7, 5, 3, 1}},
			wantMin: 16,
			wantMax: 24,
		},
		{
			name:    "unsorted",
			args:    args{arr: []int32{9, 1, 3, 7, 5}},
			wantMin: 16,
			wantMax: 24,
		},
		{
			name:    "dupes",
			args:    args{arr: []int32{1, 1, 2, 5, 2}},
			wantMin: 6,
			wantMax: 10,
		},
		{
			name:    "overflow",
			args:    args{arr: []int32{256741038, 623958417, 467905213, 714532089, 938071625}},
			wantMin: 2063136757,
			wantMax: 2744467344,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := MinMaxSum(tt.args.arr)
			if gotMin != tt.wantMin {
				t.Errorf("MinMaxSum() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("MinMaxSum() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
