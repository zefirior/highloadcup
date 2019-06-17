package store

import "fmt"

type Container struct {
	values []string
}

func (c *Container) Size() int {
	return len(c.values)
}

func (c *Container) Insert(value string) int {
	for i, v := range c.values {
		if v == value { return i }
	}
	c.values = append(c.values, value)
	return len(c.values) - 1
}

func (c *Container) Get(i int) string {
	return c.values[i]
}

func (c *Container) Length() int {
	return len(c.values)
}

func (c *Container) PrintValue() {
	fmt.Println(c.values)
}
