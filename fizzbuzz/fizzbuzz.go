package fizzbuzz

import (
	"fmt"
	"io"
)

func fizzbuzz(n int) string {
	var result string

	if n%3 == 0 {
		result += "Fizz"
	}

	if n%5 == 0 {
		result += "Buzz"
	}

	if result == "" {
		result = fmt.Sprintf("%d", n)
	}

	return result
}

func run(limit int, writer io.Writer, fn func(int) string) {
	for i := 1; i <= limit; i++ {
		if _, err := fmt.Fprintf(writer, "%s\n", fn(i)); err != nil {
			panic(err)
		}
	}
}
