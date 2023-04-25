package sm3

import (
	"github.com/tjfoc/gmsm/sm3"
)

// SM3 is the sm3 hashing method
type SM3 struct{}

// New creates a new SM3 hashing method
func New() *SM3 {
	return &SM3{}
}

// Hash generates a SM3 hash from a byte array
func (h *SM3) Hash(data []byte) []byte {
	hash := sm3.New()
	hash.Write(data)
	return hash.Sum(nil)
}
