package jindex

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestReadString(t *testing.T) {
	var str string
	r := strings.NewReader("\"str\"")
	d := NewDecoder(r)
	if off := d.Offset(); off != 0 {
		t.Fatal("expected 0, got", off)
	}

	if err := d.Decode(&str); err != nil {
		t.Fatal(err)
	}

	if str != "str" {
		t.Fatalf("expected \"str\", got \"%v\"\n", str)
	}

	if off := d.Offset(); off != int64(len("\"str\"")) {
		t.Fatalf("expected len %v, got %v\n", len("\"str\""), off)
	}

	if err := d.Decode(&str); err != io.EOF {
		t.Fatalf("expected EOF, got \"%v\"\n", err)
	}
}

type namedval struct {
	Name string `json:"name"`
}

func readValueAt(t *testing.T, path string, off *int64, v interface{}) {
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()
	n, err := f.Seek(*off, 0)
	if err != nil {
		t.Fatal(err)
	} else if n != *off {
		t.Fatal("n != off")
	}

	d := NewDecoder(f)
	if err := d.Decode(v); err != nil {
		t.Fatal(err)
	}

	*off = d.Offset()
}

func readTestSequence(t *testing.T, path string) {
	var (
		off int64
		str string
		i   int
		nv  namedval
	)

	readValueAt(t, path, &off, &str)
	if str != "foo" {
		t.Fatalf("expected \"foo\", got \"%v\"\n", str)
	}

	readValueAt(t, path, &off, &i)
	if i != 123 {
		t.Fatalf("expected 123, got %v\n", i)
	}

	readValueAt(t, path, &off, &nv)
	if nv.Name != "value" {
		t.Fatalf("expected \"value\", got \"%v\"\n", nv.Name)
	}

	readValueAt(t, path, &off, &str)
	if str != "bar" {
		t.Fatalf("expected \"bar\", got \"%v\"\n", str)
	}
}
