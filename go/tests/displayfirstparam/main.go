package main

import (
	"strings"

	"github.com/01-edu/public/go/lib"
)

func main() {
	table := append(lib.MultRandWords(), " ")
	table = append(table, "1")
	table = append(table, "1 2")
	table = append(table, "1 2 3")

	for _, s := range table {
		lib.ChallengeMain("displayfirstparam", strings.Fields(s)...)
	}
}
