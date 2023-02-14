package calc

import (
	"fmt"
	"strconv"
	"testing"
	"testing/quick"
)

// func TestCompute(t *testing.T) {
// 	s, err := Compute("1+1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s != "2" {
// 		t.Errorf("Compute(1+1) = %s, want 2", s)
// 	}
// }

// func TestCompute(t *testing.T) {
// 	computeTests := []struct {
// 		in string
// 		out string
// 	}{
// 		{"1+1", "2"},
// 		{"1.0/2.0", "0.5"},
// 	}

// 	// {in, out}

// 	for _, test := range computeTests {
// 		s, err := Compute(test.in)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if s != test.out {
// 			t.Errorf("Compute(%s) = %s. wamt %s", test.in, s, test.out)
// 		}
// 	}
// }

func TestCompu(t *testing.T) {
	add := func(a, b int16) bool {
		s, err := Compute(fmt.Sprintf("%d+%d", a, b))
		if err != nil {
			t.Fatal(err)
		}
		expected := strconv.Itoa((int(a)+int(b)))
		if s != expected {
			t.Logf("Compute(%d+%d) = %s, want %s", a, b, s, expected)
			return false
		}
		return true
	}
	if err := quick.Check(add, nil); err != nil {
		t.Fatal(err)
	}
}