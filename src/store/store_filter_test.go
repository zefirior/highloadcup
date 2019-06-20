package store

import (
	s "../structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_FilterNum_FindFnameNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Fname = "foo"

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{FnameNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{FnameNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindSnameNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Sname = "foo"

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{SnameNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{SnameNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindPhoneNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Phone = "foo"

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{PhoneNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{PhoneNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindCountryNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Country = "foo"

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{CountryNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{CountryNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindCityNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.City = "foo"

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{CityNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{CityNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindPremiumNull(t *testing.T) {
	var q Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Premium = &s.Premium{}

	store.ParseJAccount(getDefaultAcc())
	store.ParseJAccount(ja)
	store.ParseJAccount(ja)

	q = Query{PremiumNull: "0"}
	assert.Equal(t, store.FilterNum(q), 2)

	q = Query{PremiumNull: "1"}
	assert.Equal(t, store.FilterNum(q), 1)
}
