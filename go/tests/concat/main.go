package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func concat(a, b string) string {
	return a + b
}

func main() {
	table := [][]string{}

	// 30 valid pair of ramdom strings to concatenate
	for i := 0; i < 30; i++ {
		value := []string{lib.RandASCII(), lib.RandASCII()}
		table = append(table, value)
	}
	table = append(table,
		[]string{"Hello!", " How are you?"},
	)
	for _, arg := range table {
		lib.Challenge("Concat", student.Concat, concat, arg[0], arg[1])
	}
}
