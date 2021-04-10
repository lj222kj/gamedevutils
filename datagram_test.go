package gogamedevutils

import (
	"testing"
)

func TestToBytes(t *testing.T) {
	d := Datagram{Id: 1}
	b, err := d.ToBytes()

	if err != nil {
		t.Errorf("ToBytes failed, expected %v to be nil", err)
	}

	if len(b) != 5 {
		t.Errorf("ToBytes failed, expected %v to be of length", b)
	}
}

func TestFromBytes(t *testing.T) {
	d := Datagram{Id: 1}
	b, _ := d.ToBytes()

	r, err := FromBytes(b)
	if err != nil {
		t.Errorf("FromBytes failed, expected decoded datagram, got %v ", err)
	}
	if r.Id != 1 {
		t.Errorf("FromBytes failed, expected datagram to contain id, got %v ", r)
	}
}

