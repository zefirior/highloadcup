package main

import (
	"fmt"
	_ "highload/server"
	st "highload/store"
	"highload/structures"
	"highload/utils"
	"net/http"
)

const dir string = "./testdata/data"

func main() {
	var store = new(st.Store)

	utils.Check(ReadDir(dir, store))

	var (
		n   int
		acc *structures.Account
	)

	for _, acc = range store.Accounts {
		if acc == nil {
			break
		}
		if acc.Sex == 0 {
			n++
		}
	}
	fmt.Println("Num account: ", n)
	acc = store.Accounts[3]
	acc.Print()

	println("Server start")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
