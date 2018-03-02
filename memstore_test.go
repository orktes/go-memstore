package memstore

import (
	"testing"
)

func TestMemstore(t *testing.T) {
	s := New()

	s.Insert("foo", int(1))
	s.Insert("bar", int(2))

	foo, _ := s.Get("foo")
	bar, _ := s.Get("bar")

	if foo.(int) != 1 {
		t.Error("Wrong value recevied")
	}

	if bar.(int) != 2 {
		t.Error("Wrong value recevied")
	}

	return
}
