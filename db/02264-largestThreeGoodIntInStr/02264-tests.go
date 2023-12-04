package leettest

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	// ---TEST---

	// Test case 1
	num1 := "6777133339"
	t.Run("Test case 1", func(t *testing.T) {
		got := LargestGoodInteger(num1)
		want := "777"
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// Test case 2
	num2 := "2300019"
	t.Run("Test case 2", func(t *testing.T) {
		got := LargestGoodInteger(num2)
		want := "000"
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// Test case 3 - No good integer
	num3 := "42352338"
	t.Run("Test case 3", func(t *testing.T) {
		got := LargestGoodInteger(num3)
		want := ""
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// Test case 4 - All same digits
	num4 := "111"
	t.Run("Test case 4", func(t *testing.T) {
		got := LargestGoodInteger(num4)
		want := "111"
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// Test case 5 - Multiple good integers
	num5 := "222111333"
	t.Run("Test case 5", func(t *testing.T) {
		got := LargestGoodInteger(num5)
		want := "333"
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// Test case 6 - Long string with multiple good integers
	num6 := "888999777666555444333222111000"
	t.Run("Test case 6", func(t *testing.T) {
		got := LargestGoodInteger(num6)
		want := "999"
		if got != want {
			t.Errorf("Value should be '%s', Value is: '%s'", want, got)
		}
	})

	// ----END----
}
