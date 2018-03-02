# go-memstore

In memory storage for storing structs based on a compound index

# Install

```sh
go get -u github.com/orktes/go-memstore/...

```

# Usage

First create a struct that you want to store and add necessary field tags to tell what field are used to build the index

```go
//go:generate memstore-cli

// SimpleStruct simple struct that a store will be generated
// memstore:generate
type SimpleStruct struct {
	Foo int `memstore:"index"` // This will be part of the index
	Bar int `memstore:"index"`  // This will be part of the index
	Biz string
	Val float64
}


func main() {
	store := NewSimpleStructMemStore()

	insert := SimpleStruct{
		Foo: 1,
		Bar: 2,
		Biz: "foobarbiz",
		Val: 0.56,
	}
	store.Insert(insert)

	res, _ := store.Get(SimpleStructMemStoreQuery{Foo: 1, Bar: 2})
	fmt.Printf("%+v\n", res)
}


```

Then use go generate to generate memstore for the struct

```sh
go generate
```