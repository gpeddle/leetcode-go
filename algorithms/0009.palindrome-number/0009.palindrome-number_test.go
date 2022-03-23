package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example-1",
			args: args{x: 121},
			want: true,
		},
		{
			name: "example-2",
			args: args{x: -121},
			want: false,
		},
		{
			name: "example-3",
			args: args{x: 10},
			want: false,
		},
		{
			name: "added-1",
			args: args{x: 1001},
			want: true,
		},
		{
			name: "added-2",
			args: args{x: 1000111},
			want: false,
		},
		{
			name: "added-3",
			args: args{x: -110011},
			want: false,
		},
		{
			name: "added-4",
			args: args{x: 1211188},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
