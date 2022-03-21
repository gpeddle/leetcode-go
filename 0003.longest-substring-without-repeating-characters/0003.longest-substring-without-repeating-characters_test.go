package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "abcabcbb", //abc
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				s: "bbbbb", //b
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				s: "pwwkew", //wke
			},
			want: 3,
		},
		{
			name: "no-repeat case",
			args: args{
				s: "abcdefg", //abcdefg
			},
			want: 7,
		},
		{
			name: "judge 1",
			args: args{"aab"}, //ab
			want: 2,
		},
		{
			name: "judge 2",
			args: args{"dvdf"}, //vdf
			want: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
