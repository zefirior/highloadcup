package main

import (
	st "./store"
	"./utils"
	"fmt"
)

const dir string = "./testdata/data"

func main() {
	var store = new(st.Store)

	utils.Check(ReadDir(dir, store))

	var n int
	for _, acc := range store.Accounts{
		if acc.Id == 0 {break}
		if acc.Sex == 0 { n++ }
	}
	fmt.Println("Some stats: ", n)
}
