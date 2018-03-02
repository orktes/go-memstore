# go-memstore

# Install

```sh
go get -u github.com/Applifier/go-memstore/...

```

# Usage

First create a struct that you want to store and add necessary field tags to tell what field are used to build the index

```go
//go:generate memstore-cli

// SimpleStruct simple struct that a store will be generated
// memstore:generate
type SimpleStruct struct {
	Foo int `memstore:"index"`
	Bar int `memstore:"index"`
	Biz string
	Val float64
}
```

Then use go generate to generate memstore for the struct

```sh
go generate
```