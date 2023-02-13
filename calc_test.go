package calc

import "testing"

func TestCompute(t *testing.T) {
	s, err := Compute("1+1")
	if err != nil {
		t.Fatal(err)
	}
	if s != "2" {
		t.Errorf("Compute(1+1) = %s, want 2", s)
	}
}
