package sonic

import (
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func TestEncodeInto(t *testing.T) {
	var p Person
	p.Name = "sonic"
	p.Age = 20
	p.Address = "china"
	buf := make([]byte, 0, 100)
	err := EncodeInto(&buf, p, DefaultEncoderOpts)
	if err != nil {
		t.Fatal(err)
	}
	var p1 Person
	err = DecodeString(string(buf), &p1, DefaultDecoderOpts)
	if err != nil {
		t.Fatal(err)
	}

	if p1 != p {
		t.Fatal("not equal")
	}
}
