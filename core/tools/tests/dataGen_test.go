package tools_test

import (
	"api_sensor/core/tools"
	"math/rand"
	"testing"
)

func TestRandomiser(t *testing.T) {
	// Set a seed for the random number generator to ensure consistent results
	rand.Seed(12345)

	// Test that the function produces different results for different inputs
	t.Run("DifferentInputs", func(t *testing.T) {
		data1 := float32(1.0)
		data2 := float32(2.0)
		result1 := tools.Randomiser(data1)
		result2 := tools.Randomiser(data2)
		if result1 == result2 {
			t.Errorf("randomiser(%f) = %f; randomiser(%f) = %f; want different results", data1, result1, data2, result2)
		}
	})

	// Test that the function produces different results for the same input when called multiple times
	t.Run("DifferentOutputs", func(t *testing.T) {
		data := float32(1.0)
		result1 := tools.Randomiser(data)
		result2 := tools.Randomiser(data)
		if result1 == result2 {
			t.Errorf("randomiser(%f) = %f; randomiser(%f) = %f; want different results", data, result1, data, result2)
		}
	})

	// Test that the function produces values within a reasonable range around the input value
	t.Run("WithinRange", func(t *testing.T) {
		data := float32(1.0)
		for i := 0; i < 1000; i++ {
			result := tools.Randomiser(data)
			if result < data-0.5*data || result > data+0.5*data {
				t.Errorf("randomiser(%f) = %f; want within %f of %f", data, result, 0.5*data, data)
			}
		}
	})
}
