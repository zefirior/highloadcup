package main

import (
	st "./store"
	"./structures"
	"./utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
	"time"
)

const (
	NumRepeatAccs    = 1
	NumFilterRequest = 1
)

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

func ReadDir(dir string, store *st.Store) error {
	utils.PrintMemUsage()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i := 0; i < NumRepeatAccs; i++ {
		for _, file := range files {
			name := file.Name()
			matched, err := regexp.Match(`accounts_[0-9]*\.json`, []byte(name))

			if err != nil {
				return err
			}

			if matched {
				fmt.Println(i, name)
				err = readAccounts(store, dir+"/"+name)

				if err != nil {
					return err
				}
			}
		}
	}

	store.PrintStat()

	t := time.Now()

	utils.PrintMemUsage()
	var n int
	for i := 0; i < NumFilterRequest; i++ {
		n = store.FilterNum(st.Query{Sex: "f", StatusEq: "заняты"})
	}

	fmt.Println("Filter result:", n, ". Time: ", time.Now().Sub(t))

	utils.PrintMemUsage()
	runtime.GC()
	utils.PrintMemUsage()
	//debug.FreeOSMemory()
	//utils.PrintMemUsage()

	return nil
}
