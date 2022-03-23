package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example-1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "example-2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "added-madamimadam",
			args: args{
				s: "ahgdMADAMIMADAMss",
			},
			want: "MADAMIMADAM",
		},
		{
			name: "added-tacocat",
			args: args{
				s: "ftacocatxx",
			},
			want: "tacocat",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
