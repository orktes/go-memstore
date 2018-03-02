package examples

import (
	"math/rand"
	"testing"
)

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

func BenchmarkSimpleStructMemStore(b *testing.B) {
	store := NewSimpleStructMemStore()
	rand.Seed(42)

	for i := 0; i < b.N; i++ {
		insert := SimpleStruct{
			Foo: rand.Intn(b.N),
			Bar: rand.Intn(b.N),
			Biz: "foobarbiz",
			Val: 0.56,
		}
		store.Insert(insert)

		_, ok := store.Get(SimpleStructMemStoreQuery{Foo: insert.Foo, Bar: insert.Bar})
		if !ok {
			b.Error("No result found for query")
		}
	}
}
