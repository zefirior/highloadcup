package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main(){
	f, err := os.Open("./testdata/data/accounts_1.json")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 100)
	n, err := f.Read(b)
	n = utf8.RuneCount(b)

	var (
		r rune
		rs []rune
		i int
		)
	for ; i < n; i++ {
		r, _ = utf8.DecodeRune(b[i:])
		rs = append(rs, r)
	}

	if err != nil {
		panic(err)
	}
	fmt.Print("READ: ", n, " bytes - ", string(rs))
	f.Close()
}
