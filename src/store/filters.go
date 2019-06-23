package store

import (
	q "highload/query"
	s "highload/structures"
	"strconv"
	"strings"
)

func (store *Store) PropFilterSex(match bool, q *q.Query, acc *s.Account) bool {
	if q.Sex != "" {
		match = match && acc.Sex == s.SexIndex(q.Sex)
	}
	return match
}

func (store *Store) PropFilterStatus(match bool, q *q.Query, acc *s.Account) bool {
	if q.StatusEq != "" {
		match = match && acc.Status == s.StatusIndex(q.StatusEq)
	} else if q.StatusNeq != "" {
		match = match && acc.Status != s.StatusIndex(q.StatusNeq)
	}
	return match
}

func (store *Store) PropFilterEmail(match bool, q *q.Query, acc *s.Account) bool {
	if q.EmailGt != "" {
		match = match && acc.Email < q.EmailGt
	} else if q.EmailLt != "" {
		match = match && acc.Email > q.EmailLt
	} else if q.EmailDom != "" {
		match = match && store.domains.Get(int(acc.Domain)) == q.EmailDom
	}
	return match
}

func (store *Store) PropFilterFname(match bool, q *q.Query, acc *s.Account, fnameSet fnameSetType) bool {
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
	return match
}

func (store *Store) PropFilterSname(match bool, q *q.Query, acc *s.Account) bool {
	if q.SnameEq != "" {
		match = match && store.sname.Get(int(acc.Sname)) == q.SnameEq
	} else if q.SnameStarts != "" {
		match = match && strings.HasPrefix(store.sname.Get(int(acc.Sname)), q.SnameStarts)
	} else if q.SnameNull == "1" {
		match = match && acc.Sname == int16(store.sname.Insert(""))
	} else if q.SnameNull == "0" {
		match = match && acc.Sname != int16(store.sname.Insert(""))
	}
	return match
}

func (store *Store) PropFilterPhone(match bool, q *q.Query, acc *s.Account) bool {
	if q.PhoneNull == "1" {
		match = match && acc.Phone == ""
	} else if q.PhoneNull == "0" {
		match = match && acc.Phone != ""
	}
	return match
}

func (store *Store) PropFilterCountry(match bool, q *q.Query, acc *s.Account) bool {
	if q.CountryEq != "" {
		match = match && store.country.Get(int(acc.Country)) == q.CountryEq
	} else if q.CountryNull == "1" {
		match = match && acc.Country == int16(store.country.Insert(""))
	} else if q.CountryNull == "0" {
		match = match && acc.Country != int16(store.country.Insert(""))
	}
	return match
}

func (store *Store) PropFilterCity(match bool, q *q.Query, acc *s.Account) bool {
	if q.CityEq != "" {
		match = match && store.city.Get(int(acc.City)) == q.CityEq
	} else if q.CityNull == "1" {
		match = match && acc.City == int16(store.city.Insert(""))
	} else if q.CityNull == "0" {
		match = match && acc.City != int16(store.city.Insert(""))
	}
	return match
}

func (store *Store) PropFilterBirth(match bool, q *q.Query, acc *s.Account) bool {
	if q.BirthLt != "" {
		birth, err := strconv.Atoi(q.BirthLt)
		if err != nil {
			return false
		}
		match = match && acc.Birth < birth
	} else if q.BirthGt != "" {
		birth, err := strconv.Atoi(q.BirthGt)
		if err != nil {
			return false
		}

		match = match && acc.Birth > birth
	}
	return match
}

func (store *Store) PropFilterPremium(match bool, q *q.Query, acc *s.Account) bool {
	if q.PremiumNull == "1" {
		match = match && acc.Premium == nil
	} else if q.PremiumNull == "0" {
		match = match && acc.Premium != nil
	}
	return match
}
