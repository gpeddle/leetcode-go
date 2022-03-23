package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example-1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "example-2",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "example-3",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "zero",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "max-case",
			args: args{
				x: 230,
			},
			want: 32,
		},
		{
			name: "min-case",
			args: args{
				x: -231,
			},
			want: -132,
		},
		{
			name: "judge-1",
			args: args{
				x: 900000,
			},
			want: 9,
		},
		{
			name: "judge-2",
			args: args{
				x: 1463847412,
			},
			want: 2147483641,
		},
		{
			name: "judge-3",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
