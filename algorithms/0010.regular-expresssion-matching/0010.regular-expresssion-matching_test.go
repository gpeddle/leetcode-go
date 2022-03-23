package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example-1-a",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "example-2-a-dot",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "example-3-dotstar",
			args: args{
				s: "aa",
				p: ".*",
			},
			want: true,
		},
		{
			name: "added-1-a-dot",
			args: args{
				s: "bb",
				p: "A.",
			},
			want: false,
		},
		{
			name: "added-2-dotstar-a-dotstar",
			args: args{
				s: "bbAbb",
				p: ".*aA.*",
			},
			want: true,
		},
		{
			name: "added-3-dot-a-dot",
			args: args{
				s: "bbAbb",
				p: ".A.",
			},
			want: false,
		},
		{
			name: "added-4-dotstar-a-dot",
			args: args{
				s: "bbAbb",
				p: ".*A.",
			},
			want: false,
		},
		{
			name: "added-5-dot-a-dotstar",
			args: args{
				s: "bbAbb",
				p: ".A.*",
			},
			want: false,
		},
		{
			name: "added-6-long-pattern",
			args: args{
				s: "bbAbb",
				p: ".A.*",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
