package mathtool

import "testing"

// TestRandomIntBetween test if RandomIntBetween returns correct value
func TestRandomIntBetween(t *testing.T) {
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(5, 10)
		if result < 5 || result > 10 {
			t.Fatal("result should be between 5 and 10, got ", result)
		}
	}
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(10, 5)
		if result < 5 || result > 10 {
			t.Fatal("result should be between 5 and 10, got ", result)
		}
	}
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(-5, -10)
		if result < -10 || result > -5 {
			t.Fatal("result should be between -5 and -10, got ", result)
		}
	}
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(-10, -5)
		if result < -10 || result > -5 {
			t.Fatal("result should be between -5 and -10, got ", result)
		}
	}
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(-5, 10)
		if result < -5 || result > 10 {
			t.Fatal("result should be between -5 and 10, got ", result)
		}
	}
	for i := 0; i < 10; i++ {
		result := RandomIntBetween(10, -5)
		if result < -5 || result > 10 {
			t.Fatal("result should be between -5 and 10, got ", result)
		}
	}
	result := RandomIntBetween(5, 5)
	if result != 5 {
		t.Fatal("result should be equals to 5, got ", result)
	}
}
