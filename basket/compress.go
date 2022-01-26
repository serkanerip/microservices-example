package basket

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

type Compressor interface {
	Compress(bytes []byte) []byte
	Decompress(bytes []byte) []byte
}

type ZlibCompressor struct {
}

func (c *ZlibCompressor) Compress(in []byte) []byte {
	var out bytes.Buffer
	w := zlib.NewWriter(&out)
	w.Write(in)
	w.Close()

	return out.Bytes()
}

func (c *ZlibCompressor) Decompress(in []byte) ([]byte, error) {
	b := bytes.NewBuffer(in)
	r, err := zlib.NewReader(b)
	if err != nil {
		return []byte{}, err
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}
	r.Close()
	return out, nil
}
