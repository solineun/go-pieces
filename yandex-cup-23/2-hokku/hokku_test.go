package main

import "testing"

func TestPoints(t *testing.T) {
	tt := []struct{
		name string
		input string
		want string
	}{
		{
			"test 1",
			"yandexcup",
			"YandexCup",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := Format(test.input)
			if got != test.want {
				t.Errorf("WANT: %s, GOT: %s", test.want, got)
			}
		})
	}
}