package basket

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmihailenco/msgpack/v5"
)

func TestZlibCompressor(t *testing.T) {
	sc := ShoppingCart{
		Username: "serkan",
		Items: []ShoppingCartItem{
			{
				Quantity:    1,
				Color:       "Red",
				Price:       9.99,
				ProductID:   "1",
				ProductName: "Apple",
			},
		},
	}
	compressor := ZlibCompressor{}
	bytes, err := msgpack.Marshal(sc)
	assert.Nil(t, err)
	cBytes := compressor.Compress(bytes)

	dBytes, err := compressor.Decompress(cBytes)
	assert.Nil(t, err)
	var actualSC ShoppingCart
	if err := msgpack.Unmarshal(dBytes, &actualSC); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sc, actualSC)
}
