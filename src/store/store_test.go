package store

import (
	"github.com/stretchr/testify/assert"
	"highload/query"
	s "highload/structures"
	"testing"
)

func TestAssertLib(t *testing.T) {
	var val1, val2 = "foo", "foo"
	assert.Equal(t, val1, val2, "Should equal")
}

func TestStore_ParseJAccount(t *testing.T) {
	store := Store{}
	ja := s.JAccount{
		Id:        1,
		Birth:     1,
		Joined:    1,
		Country:   "foo",
		City:      "foo",
		Email:     "foo@boo",
		Fname:     "foo",
		Sname:     "foo",
		Phone:     "foo",
		Sex:       "f",
		Status:    "свободны",
		Interests: []string{"foo"},
		Likes:     []s.Likes{{}},
		Premium:   &s.Premium{},
	}

	assert.Equal(t, store.GetNumAccount(), 0, "Store should be empty")
	store.ParseJAccount(ja)
	assert.Equal(t, store.GetNumAccount(), 1, "Account not stored")
}

func TestStore_FilterNum_Empty(t *testing.T) {
	var q *query.Query

	store := Store{}

	q = &query.Query{}
	assert.Equal(t, store.FilterNum(q), 0)

}
