package calc

import "testing"

// func TestCompute(t *testing.T) {
// 	s, err := Compute("1+1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s != "2" {
// 		t.Errorf("Compute(1+1) = %s, want 2", s)
// 	}
// }


func TestCompute(t *testing.T) {
	computeTests := []struct {
		in string
		out string
	}{
		{"1+1", "2"},
		{"1.0/2.0", "0.5"},
	}
	
	for _, test := range computeTests {
		s, err := Compute(test.in)
		if err != nil {
			t.Fatal(err)
		}
		if s != test.out {
			t.Errorf("Compute(%s) = %s. wamt %s", test.in, s, test.out)
		}
	}
}