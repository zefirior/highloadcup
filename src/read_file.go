package main

import (
	st "store"
	"structures"
	"utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const dir string = "./testdata/data"

func readAccounts(store *st.Store, fname string) (err error) {
	tj := &structures.Json{}
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, tj); err != nil {
		return err
	}

	for _, v := range tj.Accounts {
		store.ParseJAccount(v)
	}
	return
}

func main() {
	utils.PrintMemUsage()

	var store = &st.Store{}

	files, err := ioutil.ReadDir(dir)
	check(err)
	for i := 0; i < 20; i++ {
		for _, file := range files {
			name := file.Name()
			matched, err := regexp.Match(`accounts_[0-9]*\.json`, []byte(name))
			check(err)
			if matched {
				fmt.Println(name)
				check(readAccounts(store, dir+"/"+name))
			}
			fmt.Println(matched)
		}
	}

	fmt.Println("CNT Accounts: ", store.GetNumAccount())
	//acs.Accounts[0].Print()

	utils.PrintMemUsage()
	runtime.GC()
	utils.PrintMemUsage()

}
