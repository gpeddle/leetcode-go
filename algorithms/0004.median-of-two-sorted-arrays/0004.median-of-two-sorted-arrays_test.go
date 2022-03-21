package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0000,
		},
		{
			name: "example 2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5000,
		},
		{
			name: "long and short array",
			args: args{
				nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				nums2: []int{3, 4},
			},
			want: 4.0000,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
