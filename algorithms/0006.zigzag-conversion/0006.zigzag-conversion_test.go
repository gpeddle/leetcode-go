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
			name: "",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "example 2",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "added 2",
			args: args{
				s:       "NOWISTHETIMEFORALLGOODMENTOCOMETOTHEAIDOFTHEIRCOUNTRY",
				numRows: 5,
			},
			want: "NTLNOFUOEIALETTTOTONWHMRGMOEHDHCTITEOODCMEIERRSFOOAIY",
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
