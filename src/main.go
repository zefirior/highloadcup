package main

import (
	st "./store"
	"./structures"
	"./utils"
	"fmt"
)

const dir string = "./testdata/data"

func main() {
	var store = new(st.Store)

	utils.Check(ReadDir(dir, store))

	var (
		n int
		acc structures.Account
	)

	for _, acc = range store.Accounts{
		if acc.Id == 0 {break}
		if acc.Sex == 0 { n++ }
	}
	fmt.Println("Num account: ", n)
}
