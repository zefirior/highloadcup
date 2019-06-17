package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_GetDefault(t *testing.T) {
	c := Container{}
	assert.Zero(t, c.Size(), "Container should be empty")

	i := c.Insert("qwe")
	assert.Equal(t, c.Size(), 1)

	assert.Equal(t, c.Get(i), "qwe", "Container return wrong data")
}
