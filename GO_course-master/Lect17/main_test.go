package main

import "testing"

type TestTable struct {
	a, b     int
	expected int
}

var tt = []TestTable{
	{2, 3, 5},
	{3, -3, 0},
	{4, 5, 9},
	{1, 10, 11},
	{0, 0, 0},
}

func TestAdd(t *testing.T) {
	for _, test := range tt {
		if out := Add(test.a, test.b); out != test.expected {
			t.Error("Test Failed!")
		}
	}
}

func TestMult(t *testing.T) {
	if Mult(2, 3) != 6 {
		t.Errorf("Test Failed: %d x %d != 6\n", 2, 3)
	}
}
