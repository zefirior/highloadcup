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
