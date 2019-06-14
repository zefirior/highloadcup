package main

import (
	s "../structures"
	"../utils"
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

func readAccounts(j *s.JsonAcc, fname string) (err error) {
	tj := &s.JsonAcc{}
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, tj); err != nil {
		return err
	}

	j.Accounts = append(j.Accounts, tj.Accounts...)
	return
}

func main() {
	utils.PrintMemUsage()

	var acs = &s.JsonAcc{}
	acs.Accounts = make([]s.Account, 0)

	files, err := ioutil.ReadDir(dir)
	check(err)
	for _, file := range files {
		name := file.Name()
		matched, err := regexp.Match(`accounts_[0-9]*\.json`, []byte(name))
		check(err)
		if matched {
			fmt.Println(name)
			check(readAccounts(acs, dir+"/"+name))
		}
		fmt.Println(matched)
	}

	fmt.Println("CNT Accounts: ", len(acs.Accounts))
	acs.Accounts[0].Print()

	utils.PrintMemUsage()
	runtime.GC()
	utils.PrintMemUsage()

}
