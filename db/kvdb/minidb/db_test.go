package minidb

import "testing"

func TestEntry(t *testing.T) {
	e := NewEntry([]byte("a"), []byte("v"), 11)
	buf := e.Encode()
	ne, _ := DecodeHeader(buf)
	if e.KeySize != ne.KeySize {
		t.Errorf("e=%d,ne=%d", e.KeySize, ne.Key)
	}
}
