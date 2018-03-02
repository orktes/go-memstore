package examples

import "testing"

func TestSimpleStructMemStore(t *testing.T) {
	store := NewSimpleStructMemStore()

	insert := SimpleStruct{
		Foo: 1,
		Bar: 2,
		Biz: "foobarbiz",
		Val: 0.56,
	}
	store.Insert(insert)

	res, ok := store.Get(SimpleStructMemStoreQuery{Foo: 1, Bar: 2})
	if !ok {
		t.Error("No result found for query")
	}

	if res != insert {
		t.Error("Different value returned than inserted")
	}

	if _, ok := store.Get(SimpleStructMemStoreQuery{Foo: 1, Bar: 1}); ok {
		t.Error("Should not return record which has not been inserted")
	}
}
