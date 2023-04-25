package sm3

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// _byteArray is a helper to turn a string in to a byte array
func _byteArray(input string) []byte {
	x, _ := hex.DecodeString(input)
	return x
}

func TestHash(t *testing.T) {
	var tests = []struct {
		data   []byte
		output []byte
	}{
		{
			data:   _byteArray("e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637"),
			output: _byteArray("92c7a270abba6545cff680c3452f1573b3b672d66f663b4c1d1d3ce7c35b5170"),
		},
	}

	hash := New()
	for i, test := range tests {
		output := hash.Hash(test.data)
		assert.Equal(t, test.output, output, fmt.Sprintf("failed at test %d", i))
	}
}
