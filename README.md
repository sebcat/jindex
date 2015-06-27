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

JSON decoder keeping track of the number of read bytes

#### func  NewDecoder

```go
func NewDecoder(r io.Reader) *Decoder
```
Create a new JSON decoder, keeping track of the number of read bytes

#### func (*Decoder) Decode

```go
func (d *Decoder) Decode(v interface{}) error
```
Read the next JSON-encoded value and store it in the value pointed to by v

#### func (*Decoder) NRead

```go
func (d *Decoder) NRead() int64
```
Return the number of bytes read by this decoder
