package fizzbuzz

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// job isn't to test function but to specify a desired behaviour!
func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		arg  int
		want string
	}{
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("fizzbuzz(%d)", tt.arg), func(t *testing.T) {
			if got := fizzbuzz(tt.arg); got != tt.want {
				t.Errorf("fizzbuzz(%d) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("should generate numbers from 1 to specified limit", func(t *testing.T) {
		var buf bytes.Buffer

		run(2, &buf, func(n int) string { return fmt.Sprintf("%d", n) })

		exp := "1\n2\n"
		got := string(buf.Bytes())

		if exp != got {
			t.Errorf("expected %q, got: %q", exp, got)
		}
	})

	t.Run("should apply function to generated numbers", func(t *testing.T) {
		var buf bytes.Buffer

		run(5, &buf, fizzbuzz)

		exp := "1\n2\nFizz\n4\nBuzz\n"
		got := string(buf.Bytes())

		if exp != got {
			t.Errorf("expected %q, got: %q", exp, got)
		}
	})

	t.Run("test fizzbuzz on first 100 numbers", func(t *testing.T) {
		run(100, os.Stdout, fizzbuzz)
	})
}
