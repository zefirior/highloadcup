package store

import (
	s "structures"
)

type Store struct {
	accounts []s.Account
	country Container
	city Container
	fname Container
	sname Container
	interests Container
}

func (store *Store) GetNumAccount() int { return len(store.accounts) }

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
	store.accounts = append(store.accounts, a)
}

func getInterests(store *Store, interests []string) (res []int) {
	for _, v := range interests {
		res = append(res, store.interests.Insert(v))
	}
	return
}
