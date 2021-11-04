package tempconv

import (
	"math"
	"testing"
)

// TestKToC tests KToC
func TestKToC(t *testing.T) {
	// Test the conversion from Kelvin to Celsius
	k := Kelvin(0)
	c := KToC(k)
	if c != -273.15 {
		t.Errorf("KToC(273.15) failed: %v\n", c)
	}
}

// TestKToF tests KToF
func TestKToF(t *testing.T) {
	// Test the conversion from Kelvin to Celsius
	k := Kelvin(0)
	f := KToF(k)
	if math.Round(float64(f)*100)/100 != -459.67 {
		t.Errorf("KToF(523.67) failed: %v\n", f)
	}
}
