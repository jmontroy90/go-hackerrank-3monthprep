package week1

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	type args struct {
		scores []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "example",
			args: args{scores: []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}},
			want: []int32{2, 4},
		},
		{
			name: "with-zero",
			args: args{scores: []int32{0, 9, 3, 10, 2, 20}},
			want: []int32{3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakingRecords(tt.args.scores); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreakingRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
