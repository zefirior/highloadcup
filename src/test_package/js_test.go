package test_package

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	e := Example{
		Row: "foo",
		Nodes: []Node{
			{"left", "right"},
			{"1", "2"},
		},
	}
	js, _ := json.Marshal(e)
	if string(js) != `{"Row":"foo","Nodes":[{"Left":"left","Right":"right"},{"Left":"1","Right":"2"}]}` {
		t.Error("Got: ", string(js))
	}
}

func TestEqual(t *testing.T) {
	e1 := Example{"foo", []Node{
		{"1", "2"},
	}}
	e2 := e1
	if !e1.Equal(&e2) {
		t.Error("Structures 'Example' not equals")
		t.Error("Left: ", e1)
		t.Error("Right: ", e2)
	}
}

func TestNotEqual(t *testing.T) {
	e1 := Example{"foo", []Node{
		{"1", "2"},
	}}
	e2 := e1
	e2.Row = "boo"
	if e1.Equal(&e2) {
		t.Error("Different structures 'Example' equals")
		t.Error("Left: ", e1)
		t.Error("Right: ", e2)
	}
}

func TestNotEqualInNode(t *testing.T) {
	e1 := Example{"foo", []Node{
		{"1", "2"},
	}}
	e2 := e1
	e2.Nodes = []Node{{"boo", "2"}}
	if e1.Equal(&e2) {
		t.Error("Different structures 'Example' equals")
		t.Error("Left: ", e1)
		t.Error("Right: ", e2)
	}
}

func TestUnmarshal(t *testing.T) {
	got := &Example{}
	js := []byte(`{"asd":123,"Row":"some str","Foo":119,"Nodes":[{"Left":"1","Right":"2"}]}`)
	if err := json.Unmarshal(js, got); err != nil {
		t.Error(err)
	}
	expect := &Example{"some str", []Node{{"1", "2"}}}
	if !got.Equal(expect) {
		t.Error("Got: ", got)
		t.Error("Expect: ", expect)
	}
}
