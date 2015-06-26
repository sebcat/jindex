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
	d := NewDecoder(r, 0)
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

	d := NewDecoder(f, *off)
	if err := d.Decode(v); err != nil {
		t.Fatal(err)
	}
	newOff := d.Offset()
	t.Logf("read %v, %v\n", *off, newOff)
	*off = newOff
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

	t.Log("got foo")
	readValueAt(t, path, &off, &i)
	if i != 123 {
		t.Fatalf("expected 123, got %v\n", i)
	}

	t.Log("got 123")
	readValueAt(t, path, &off, &nv)
	if nv.Name != "value" {
		t.Fatalf("expected \"value\", got \"%v\"\n", nv.Name)
	}

	t.Log("got struct")
	readValueAt(t, path, &off, &str)
	if str != "bar" {
		t.Fatalf("expected \"bar\", got \"%v\"\n", str)
	}

	t.Log("got bar")
}

func TestReadNoWhitespace(t *testing.T) {
	readTestSequence(t, "testdata/t0.json")
}

func TestReadWithWhitespace(t *testing.T) {
	readTestSequence(t, "testdata/t1.json")
}
