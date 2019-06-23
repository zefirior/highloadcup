package store

import (
	"fmt"
	"highload/query"
	s "highload/structures"
	"strings"
)

const NumAccounts = 40000

type Store struct {
	Accounts  [NumAccounts]*s.Account
	AccOffset int
	country   Container
	city      Container
	fname     Container
	sname     Container
	interests Container
	domains   Container
}

func (store *Store) GetNumAccount() int { return store.AccOffset }

func (store *Store) GetNumCountry() int   { return store.country.Length() }
func (store *Store) GetNumCity() int      { return store.city.Length() }
func (store *Store) GetNumDomain() int    { return store.domains.Length() }
func (store *Store) GetNumFname() int     { return store.fname.Length() }
func (store *Store) GetNumSname() int     { return store.sname.Length() }
func (store *Store) GetNumInterests() int { return store.interests.Length() }

func (store *Store) PrintStat() {
	//fmt.Println("Country: ")
	//store.country.PrintValue()

	fmt.Println("CNT Account: ", store.GetNumAccount())
	fmt.Println("CNT Country: ", store.GetNumCountry())
	fmt.Println("CNT City: ", store.GetNumCity())
	fmt.Println("CNT Domain: ", store.GetNumCity())
	fmt.Println("CNT Fname: ", store.GetNumFname())
	fmt.Println("CNT Sname: ", store.GetNumSname())
	fmt.Println("CNT Interests: ", store.GetNumInterests())
}

func (store *Store) ParseJAccount(account s.JAccount) {
	domain := strings.Split(account.Email, "@")[1]

	a := &s.Account{
		Id:      account.Id,
		Birth:   account.Birth,
		Joined:  account.Joined,
		Email:   account.Email,
		Phone:   account.Phone,
		Likes:   account.Likes,
		Premium: account.Premium,
		Country: int16(store.country.Insert(account.Country)),
		City:    int16(store.city.Insert(account.City)),
		Fname:   int16(store.fname.Insert(account.Fname)),
		Sname:   int16(store.sname.Insert(account.Sname)),
		Domain:  int16(store.domains.Insert(domain)),

		Sex:       s.SexIndex(account.Sex),
		Status:    s.StatusIndex(account.Status),
		Interests: getInterests(store, account.Interests),
	}
	store.Accounts[store.AccOffset] = a
	store.AccOffset++
}

func getInterests(store *Store, interests []string) (res []int16) {
	for _, v := range interests {
		res = append(res, int16(store.interests.Insert(v)))
	}
	return
}

type fnameSetType map[string]struct{}

func parseFnameSet(set string) fnameSetType {
	fnameSet := fnameSetType{}
	for _, fname := range strings.Split(set, ",") {
		fnameSet[fname] = struct{}{}
	}
	return fnameSet
}

func (store *Store) FilterNum(q *query.Query) (n int) {
	fnameSet := parseFnameSet(q.FnameAny)

	for i := 0; i < store.AccOffset; i++ {
		match := true
		acc := store.Accounts[i]

		match = store.PropFilterSex(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterStatus(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterEmail(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterFname(match, q, acc, fnameSet)
		if !match {
			continue
		}

		match = store.PropFilterSname(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterPhone(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterCountry(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterCity(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterBirth(match, q, acc)
		if !match {
			continue
		}

		match = store.PropFilterPremium(match, q, acc)
		if !match {
			continue
		}

		if match {
			n++
		}
	}
	return
}
