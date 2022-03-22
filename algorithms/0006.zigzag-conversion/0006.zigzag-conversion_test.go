package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example-1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "example-2-single-char",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "2-rows-odd",
			args: args{
				s:       "ABC",
				numRows: 2,
			},
			want: "ACB",
		},
		{
			name: "2-rows-even",
			args: args{
				s:       "ABCD",
				numRows: 2,
			},
			want: "ACBD",
		},
		{
			name: "3-rows-odd",
			args: args{
				s:       "ABCDEFGHI",
				numRows: 3,
			},
			want: "AEIBDFHCG",
		},
		{
			name: "3-rows-even",
			args: args{
				s:       "ABCDEFGHIJ",
				numRows: 3,
			},
			want: "AEIBDFHJCG",
		},
		{
			name: "4-rows-odd",
			args: args{
				s:       "ABCDEFGHIJK",
				numRows: 4,
			},
			want: "AGBFHCEIKDJ",
		},
		{
			name: "4-rows-even",
			args: args{
				s:       "ABCDEFGHIJKLMNOP",
				numRows: 4,
			},
			want: "AGMBFHLNCEIKODJP",
		},
		{
			name: "typing-test",
			args: args{
				s:       "NOWISTHETIMEFORALLGOODMENTOCOMETOTHEAIDOFTHEIRCOUNTRY",
				numRows: 5,
			},
			want: "NTLNOFUOEIALETTTOTONWHMRGMOEHDHCTITEOODCMEIERRSFOOAIY",
		},
		{
			name: "1-row-long",
			args: args{
				s:       "ENGLANDEXPECTSEVERYMANTODOHISDUTY",
				numRows: 1,
			},
			want: "ENGLANDEXPECTSEVERYMANTODOHISDUTY",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
