# jindex
--
    import "github.com/sebcat/jindex"

Indexed reading of JSON streams

## Usage

#### type Decoder

```go
type Decoder struct {
}
```

JSON decoder keeping track of the read offset

#### func  NewDecoder

```go
func NewDecoder(r io.Reader) *Decoder
```
Create a new indexing JSON decoder

#### func (*Decoder) Decode

```go
func (d *Decoder) Decode(v interface{}) error
```
Read the next JSON-encoded value and store it in the value pointed to by v

#### func (*Decoder) Offset

```go
func (d *Decoder) Offset() int64
```
Return the offset past the end of the last read JSON value in the read stream
