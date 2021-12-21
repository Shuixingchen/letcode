package main

import "testing"

//go test -run cal_test.go -v

func TestMul(t *testing.T) {
	cases := []struct {
		Name         string
		a, b, expect int64
	}{
		{"pos", 1, 2, 2},
		{"pos", 22, 2, 44},
	}
	for _, c := range cases {
		t.Run(t.Name(), func(t *testing.T) {
			if res := Mul(c.a, c.b); res != c.expect {
				t.Fatalf("%d*%d expect %d but %d", c.a, c.b, c.expect, res)
			}
		})
	}
}
