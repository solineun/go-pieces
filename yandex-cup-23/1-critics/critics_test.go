package main

import "testing"

func TestPoints(t *testing.T) {
	tt := []struct{
		name string
		n, m int
		fpts []int
		want int
	}{
		{
			"test 1",
			10, 5,
			[]int{10, 0, 1, 0, 3},
			117,
		},
		{
			"test 2",
			5, 5,
			[]int{0, 0, 0, 0, 0},
			0,
		},
		{
			"test 3",
			1, 3,
			[]int{1, 1, 1},
			5,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := Points(test.n, test.m, test.fpts)
			if got != test.want {
				t.Errorf("WANT: %d, GOT: %d", test.want, got)
			}
		})
	}
}