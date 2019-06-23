package store

import (
	"github.com/stretchr/testify/assert"
	"highload/query"
	s "highload/structures"
	"testing"
)

func getDefaultAcc() s.JAccount {
	ja := s.JAccount{Sex: "f", Status: "свободны", Email: "foo@bar.com"}
	return ja
}

// SEX

func TestStore_Filter_Sex(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Sex = "f"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{Sex: "m"}
	assert.Equal(t, store.PropFilterSex(true, q, acc), false)

	q = &query.Query{Sex: "f"}
	assert.Equal(t, store.PropFilterSex(true, q, acc), true)
}

// STATUS

func TestStore_Filter_StatusEq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Status = "свободны"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{StatusEq: "заняты"}
	assert.Equal(t, store.PropFilterStatus(true, q, acc), false)

	q = &query.Query{StatusEq: "свободны"}
	assert.Equal(t, store.PropFilterStatus(true, q, acc), true)
}

func TestStore_Filter_StatusNeq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Status = "свободны"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{StatusNeq: "свободны"}
	assert.Equal(t, store.PropFilterStatus(true, q, acc), false)
	q = &query.Query{StatusNeq: "заняты"}
	assert.Equal(t, store.PropFilterStatus(true, q, acc), true)
}

// Email

func TestStore_Filter_EmailLt(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Email = "foo@bar.com"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{EmailLt: "zoo"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), false)
	q = &query.Query{EmailLt: "doo"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), true)
}

func TestStore_Filter_EmailGt(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Email = "foo@bar.com"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{EmailGt: "doo"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), false)
	q = &query.Query{EmailGt: "zoo"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), true)
}

func TestStore_Filter_EmailDom(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Email = "foo@bar.com"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{EmailDom: "foo"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), false)

	q = &query.Query{EmailDom: "bar.com"}
	assert.Equal(t, store.PropFilterEmail(true, q, acc), true)
}

// FNAME

func TestStore_Filter_FnameEq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Fname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	fset := fnameSetType{"": {}}

	q = &query.Query{FnameEq: "boo"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, fset), false)

	q = &query.Query{FnameEq: "foo"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, fset), true)
}

func TestStore_Filter_FnameAny(t *testing.T) {
	var (
		q    *query.Query
		fset fnameSetType
	)

	store := Store{}
	ja := getDefaultAcc()
	ja.Fname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{FnameAny: "1,2,3"}
	fset = parseFnameSet(q.FnameAny)
	assert.Equal(t, store.PropFilterFname(true, q, acc, fset), false)

	q = &query.Query{FnameAny: "1,2,foo,4"}
	fset = parseFnameSet(q.FnameAny)
	assert.Equal(t, store.PropFilterFname(true, q, acc, fset), true)
}

func TestStore_Filter_FnameNull_WithFname(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Fname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{FnameNull: "0"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, parseFnameSet("")), true)

	q = &query.Query{FnameNull: "1"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, parseFnameSet("")), false)
}

func TestStore_Filter_FnameNull_WithoutFname(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{FnameNull: "0"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, parseFnameSet("")), false)

	q = &query.Query{FnameNull: "1"}
	assert.Equal(t, store.PropFilterFname(true, q, acc, parseFnameSet("")), true)
}

// SNAME

func TestStore_Filter_SnameEq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Sname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{SnameEq: "boo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)

	q = &query.Query{SnameEq: "foo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), true)
}

func TestStore_Filter_SnameStarts(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Sname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{SnameStarts: "bo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)
	q = &query.Query{SnameStarts: "fop"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)
	q = &query.Query{SnameStarts: "fooo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)

	q = &query.Query{SnameStarts: "fo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), true)
	q = &query.Query{SnameStarts: "foo"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), true)
}

func TestStore_Filter_SnameNull_WithSname(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Sname = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{SnameNull: "0"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), true)

	q = &query.Query{SnameNull: "1"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)
}

func TestStore_Filter_SnameNull_WithoutSname(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{SnameNull: "0"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), false)

	q = &query.Query{SnameNull: "1"}
	assert.Equal(t, store.PropFilterSname(true, q, acc), true)
}

// PHONE

func TestStore_Filter_PhoneNull_WithPhone(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Phone = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{PhoneNull: "0"}
	assert.Equal(t, store.PropFilterPhone(true, q, acc), true)

	q = &query.Query{PhoneNull: "1"}
	assert.Equal(t, store.PropFilterPhone(true, q, acc), false)
}

func TestStore_Filter_PhoneNull_WithoutPhone(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{PhoneNull: "0"}
	assert.Equal(t, store.PropFilterPhone(true, q, acc), false)

	q = &query.Query{PhoneNull: "1"}
	assert.Equal(t, store.PropFilterPhone(true, q, acc), true)
}

// COUNTRY

func TestStore_Filter_CountryEq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Country = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CountryEq: "boo"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), false)

	q = &query.Query{CountryEq: "foo"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), true)
}

func TestStore_Filter_CountryNull_With(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Country = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CountryNull: "0"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), true)

	q = &query.Query{CountryNull: "1"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), false)
}

func TestStore_Filter_CountryNull_Without(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CountryNull: "0"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), false)

	q = &query.Query{CountryNull: "1"}
	assert.Equal(t, store.PropFilterCountry(true, q, acc), true)
}

// City

func TestStore_Filter_CityEq(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.City = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CityEq: "boo"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), false)

	q = &query.Query{CityEq: "foo"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), true)
}

func TestStore_Filter_CityNull_With(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.City = "foo"

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CityNull: "0"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), true)

	q = &query.Query{CityNull: "1"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), false)
}

func TestStore_Filter_CityNull_Without(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{CityNull: "0"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), false)

	q = &query.Query{CityNull: "1"}
	assert.Equal(t, store.PropFilterCity(true, q, acc), true)
}

// Birth

func TestStore_Filter_BirthLt(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Birth = 123456

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{BirthLt: "123"}
	assert.Equal(t, store.PropFilterBirth(true, q, acc), false)
	q = &query.Query{BirthLt: "1234567"}
	assert.Equal(t, store.PropFilterBirth(true, q, acc), true)
}

func TestStore_Filter_BirthGt(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Birth = 123456

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{BirthGt: "1234567"}
	assert.Equal(t, store.PropFilterBirth(true, q, acc), false)
	q = &query.Query{BirthGt: "123"}
	assert.Equal(t, store.PropFilterBirth(true, q, acc), true)
}

// Premium

func TestStore_Filter_PremiumNull_With(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()
	ja.Premium = &s.Premium{}

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{PremiumNull: "0"}
	assert.Equal(t, store.PropFilterPremium(true, q, acc), true)

	q = &query.Query{PremiumNull: "1"}
	assert.Equal(t, store.PropFilterPremium(true, q, acc), false)
}

func TestStore_Filter_PremiumNull_Without(t *testing.T) {
	var q *query.Query

	store := Store{}
	ja := getDefaultAcc()

	store.ParseJAccount(ja)
	acc := store.Accounts[0]

	q = &query.Query{PremiumNull: "0"}
	assert.Equal(t, store.PropFilterPremium(true, q, acc), false)

	q = &query.Query{PremiumNull: "1"}
	assert.Equal(t, store.PropFilterPremium(true, q, acc), true)
}
