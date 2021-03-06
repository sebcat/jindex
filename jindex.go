// Indexed reading of JSON streams
package jindex

import (
	"bytes"
	"encoding/json"
	"io"
)

type reader struct {
	r io.Reader
	n int64
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.n += int64(n)
	return
}

// JSON decoder keeping track of the number of read bytes
type Decoder struct {
	dec *json.Decoder
	r   reader
}

// Create a new JSON decoder, keeping track of the number of read bytes
func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{r: reader{r: r}}
	dec.dec = json.NewDecoder(&dec.r)
	return dec
}

// Return the number of bytes read by this decoder
func (d *Decoder) NRead() int64 {
	r := d.dec.Buffered().(*bytes.Reader)
	return d.r.n - int64(r.Len())
}

// Read the next JSON-encoded value and store it in the value pointed to by v
func (d *Decoder) Decode(v interface{}) error {
	return d.dec.Decode(v)
}
