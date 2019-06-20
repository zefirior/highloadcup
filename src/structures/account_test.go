package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountEqual(t *testing.T) {
	a1 := Account{Fname: 123}
	a2 := a1
	assert.True(t, a1.Equal(a2), "Accounts should be equal")
}

func TestSexIndex(t *testing.T) {
	assert.Equal(t, SexIndex("f"), int8(0))
	assert.Equal(t, SexIndex("m"), int8(1))
}

func TestSexValue(t *testing.T) {
	assert.Equal(t, SexValue(0), "f")
	assert.Equal(t, SexValue(1), "m")
}

func TestStatusIndex(t *testing.T) {
	assert.Equal(t, StatusIndex("свободны"), int8(0))
	assert.Equal(t, StatusIndex("заняты"), int8(1))
	assert.Equal(t, StatusIndex("всё сложно"), int8(2))
}

func TestStatusValue(t *testing.T) {
	assert.Equal(t, StatusValue(0), "свободны")
	assert.Equal(t, StatusValue(1), "заняты")
	assert.Equal(t, StatusValue(2), "всё сложно")
}
