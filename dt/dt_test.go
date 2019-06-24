package main

import (
	"testing"
)

func TestDtProcessHappy(t *testing.T) {
	base := "20190524"
	cases := []struct {
		in   int
		want string
	}{
		{0, "2019/05/24 00:00:00 (KST) // 1558623600"},
		{1, "2019/05/25 00:00:00 (KST) // 1558710000"},
		{5, "2019/05/29 00:00:00 (KST) // 1559055600"},
		{-1, "2019/05/23 00:00:00 (KST) // 1558537200"},
	}
	for _, c := range cases {
		got := Process(base, c.in)
		if c.want != got {
			t.Errorf("err.. [%q] vs [%q]", c.want, got)
		}
	}
}
