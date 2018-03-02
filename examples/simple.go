package examples

//go:generate memstore-cli

// SimpleStruct simple struct that a store will be generated
// memstore:generate
type SimpleStruct struct {
	Foo int `memstore:"index"`
	Bar int `memstore:"index"`
	Biz string
	Val float64
}
