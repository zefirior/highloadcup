package store

import (
	s "../structures"
	"fmt"
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

type Query struct {
	Sex string

	StatusEq  string
	StatusNeq string

	EmailLt  string
	EmailGt  string
	EmailDom string

	FnameEq   string
	FnameAny  string
	FnameNull string

	SnameEq     string
	SnameStarts string
	SnameNull   string

	PhoneNull string

	CountryNull string

	CityNull string

	PremiumNull string
}

func (store *Store) FilterNum(q Query) (n int) {
	fnameSet := map[string]struct{}{}
	if q.FnameAny != "" {
		for _, fname := range strings.Split(q.FnameAny, ",") {
			fnameSet[fname] = struct{}{}
		}
	}

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

		if q.EmailGt != "" {
			match = match && acc.Email < q.EmailGt
		} else if q.EmailLt != "" {
			match = match && acc.Email > q.EmailLt
		} else if q.EmailDom != "" {
			match = match && store.domains.Get(int(acc.Domain)) == q.EmailDom
		}

		if q.FnameEq != "" {
			match = match && store.fname.Get(int(acc.Fname)) == q.FnameEq
		} else if q.FnameAny != "" {
			_, ok := fnameSet[store.fname.Get(int(acc.Fname))]
			match = match && ok
		} else if q.FnameNull == "1" {
			match = match && acc.Fname == int16(store.fname.Insert(""))
		} else if q.FnameNull == "0" {
			match = match && acc.Fname != int16(store.fname.Insert(""))
		}

		if q.SnameEq != "" {
			match = match && store.sname.Get(int(acc.Sname)) == q.SnameEq
		} else if q.SnameStarts != "" {
			match = match && strings.HasPrefix(store.sname.Get(int(acc.Sname)), q.SnameStarts)
		} else if q.SnameNull == "1" {
			match = match && acc.Sname == int16(store.sname.Insert(""))
		} else if q.SnameNull == "0" {
			match = match && acc.Sname != int16(store.sname.Insert(""))
		}

		if q.PhoneNull == "1" {
			match = match && acc.Phone == ""
		} else if q.PhoneNull == "0" {
			match = match && acc.Phone != ""
		}

		if q.CountryNull == "1" {
			match = match && acc.Country == int16(store.country.Insert(""))
		} else if q.CountryNull == "0" {
			match = match && acc.Country != int16(store.country.Insert(""))
		}

		if q.CityNull == "1" {
			match = match && acc.City == int16(store.city.Insert(""))
		} else if q.CityNull == "0" {
			match = match && acc.City != int16(store.city.Insert(""))
		}

		if q.PremiumNull == "1" {
			match = match && acc.Premium == nil
		} else if q.PremiumNull == "0" {
			match = match && acc.Premium != nil
		}

		if match {
			n++
		}
	}
	return
}
