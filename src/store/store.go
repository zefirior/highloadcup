package store

import (
	s "../structures"
	"fmt"
)

const NumAccounts = 1400000

type Store struct {
	Accounts  [NumAccounts]s.Account
	AccOffset int
	country   Container
	city      Container
	fname     Container
	sname     Container
	interests Container
}

func (store *Store) GetNumAccount() int { return store.AccOffset }

func (store *Store) GetNumCountry() int 	{ return store.country.Length() }
func (store *Store) GetNumCity() int 		{ return store.city.Length() }
func (store *Store) GetNumFname() int 		{ return store.fname.Length() }
func (store *Store) GetNumSname() int 		{ return store.sname.Length() }
func (store *Store) GetNumInterests() int 	{ return store.interests.Length() }

func (store *Store) PrintStat() {
	//fmt.Println("Country: ")
	//store.country.PrintValue()

	fmt.Println("CNT Account: ", 	store.GetNumAccount())
	fmt.Println("CNT Country: ", 	store.GetNumCountry())
	fmt.Println("CNT City: ", 		store.GetNumCity())
	fmt.Println("CNT Fname: ", 	store.GetNumFname())
	fmt.Println("CNT Sname: ", 	store.GetNumSname())
	fmt.Println("CNT Interests: ", store.GetNumInterests())
}

func (store *Store) ParseJAccount(account s.JAccount) {
	a := s.Account{
		Id: account.Id,
		Birth: account.Birth,
		Joined: account.Joined,
		Email: account.Email,
		Phone: account.Phone,
		Likes: account.Likes,
		Premium: account.Premium,
		Country: store.country.Insert(account.Country),
		City: store.city.Insert(account.City),
		Fname: store.fname.Insert(account.Fname),
		Sname: store.sname.Insert(account.Sname),

		Sex: s.SexIndex(account.Sex),
		Status: s.StatusIndex(account.Status),
		Interests: getInterests(store, account.Interests),
	}
	store.Accounts[store.AccOffset] = a
	store.AccOffset++
}

func getInterests(store *Store, interests []string) (res []int) {
	for _, v := range interests {
		res = append(res, store.interests.Insert(v))
	}
	return
}

type Query struct {
	Sex string
	StatusEq string
	StatusNeq string
}

func (store *Store) FilterNum(q Query) (n int) {
	for i := 0; i < store.AccOffset; i++ {
		match := true
		acc := store.Accounts[i]

		if q.Sex != "" {
			match = match && acc.Sex == s.SexIndex(q.Sex)
		}

		if q.StatusEq != "" {
			match = match && acc.Status == s.StatusIndex(q.StatusEq)
		} else if q.StatusNeq != "" {
			match = match && acc.Status != s.StatusIndex(q.StatusNeq)
		}

		if match {
			n++
		}
	}
	return
}