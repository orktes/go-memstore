package memstore

import (
	"io"

	radix "github.com/armon/go-radix"
)

// Store instance of a memostore
type Store struct {
	r *radix.Tree
}

// New initializes a new instance of a store
func New() *Store {
	r := radix.New()
	return &Store{r: r}
}

// Insert inserts a key to the memstore
func (s *Store) Insert(key string, val interface{}) {
	s.r.Insert(key, val)
}

// Get returns a value from the memstore
func (s *Store) Get(key string) (val interface{}, found bool) {
	return s.r.Get(key)
}

// WriteTo write store contents to
func (s *Store) WriteTo(writer io.Writer) error {
	return nil
}
