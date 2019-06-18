package store

import (
	s "../structures"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestAssertLib(t *testing.T) {
	var val1, val2 = "foo", "foo"
	assert.Equal(t, val1, val2, "Should equal")
}


func TestStore_ParseJAccount(t *testing.T) {
	store := Store{}
	ja := s.JAccount{Birth:1234, Fname:"foo", Sex:"f", Status:"свободны"}

	assert.Equal(t, store.GetNumAccount(), 0, "Store should be empty")
	store.ParseJAccount(ja)
	assert.Equal(t, store.GetNumAccount(), 1, "Account not stored")
}

func TestStore_FilterNum_Empty(t *testing.T) {
	var q Query

	store := Store{}

	q = Query{}
	assert.Equal(t, store.FilterNum(q), 0)

}

func TestStore_FilterNum_FindSex(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Birth:1234, Fname:"foo", Sex:"f", Status:"свободны"}

	store.ParseJAccount(ja)

	q = Query{Sex: "m"}
	assert.Equal(t, store.FilterNum(q), 0)
	q = Query{Sex: "f"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindStatusEq(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Birth:1234, Fname:"foo", Sex:"f", Status:"свободны"}

	store.ParseJAccount(ja)

	q = Query{StatusEq: "заняты"}
	assert.Equal(t, store.FilterNum(q), 0)
	q = Query{StatusEq: "свободны"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindStatusNeq(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Birth:1234, Fname:"foo", Sex:"f", Status:"свободны"}

	store.ParseJAccount(ja)

	q = Query{StatusNeq: "свободны"}
	assert.Equal(t, store.FilterNum(q), 0)
	q = Query{StatusNeq: "заняты"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindEmailLt(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Sex:"f", Status:"свободны", Email: "foo@goo.com"}

	store.ParseJAccount(ja)

	q = Query{EmailLt: "zoo"}
	assert.Equal(t, store.FilterNum(q), 0)
	q = Query{EmailLt: "doo"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindEmailGt(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Sex:"f", Status:"свободны", Email: "foo@goo.com"}

	store.ParseJAccount(ja)

	q = Query{EmailGt: "doo"}
	assert.Equal(t, store.FilterNum(q), 0)
	q = Query{EmailGt: "zoo"}
	assert.Equal(t, store.FilterNum(q), 1)
}

func TestStore_FilterNum_FindSexStatus(t *testing.T) {
	var q Query

	store := Store{}
	ja := s.JAccount{Birth:1234, Fname:"foo", Sex:"f", Status:"свободны"}

	store.ParseJAccount(ja)

	q = Query{Sex: "f", StatusNeq: "свободны"}
	assert.Equal(t, store.FilterNum(q), 0)

	q = Query{Sex: "f", StatusEq: "свободны"}
	assert.Equal(t, store.FilterNum(q), 1)

	q = Query{Sex: "f", StatusNeq: "заняты"}
	assert.Equal(t, store.FilterNum(q), 1)

	q = Query{Sex: "m", StatusNeq: "заняты"}
	assert.Equal(t, store.FilterNum(q), 0)

}
