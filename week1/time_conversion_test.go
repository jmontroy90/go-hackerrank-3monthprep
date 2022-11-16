package week1

import "testing"

func TestTimeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "12-AM",
			args: args{s: "12:01:00AM"},
			want: "00:01:00",
		},
		{
			name: "12-PM",
			args: args{s: "12:01:00PM"},
			want: "12:01:00",
		},
		{
			name: "normal-AM",
			args: args{s: "07:01:00AM"},
			want: "07:01:00",
		},
		{
			name: "normal-PM",
			args: args{s: "07:01:00PM"},
			want: "19:01:00",
		},
		{
			name: "failed-case", // why
			args: args{s: "02:34:50PM"},
			want: "14:34:50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeConversion(tt.args.s); got != tt.want {
				t.Errorf("TimeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
