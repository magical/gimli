package gimli

import (
	"bytes"
	"crypto/cipher"
	"testing"
)

var _ cipher.AEAD = &AEAD{}

func TestAEAD(t *testing.T) {
	a := NewAEAD([]byte("key"))
	var nonce [15]byte
	var plaintext = []byte("plaintext")
	ciphertext := a.Seal(nil, nonce[:], plaintext, nil)
	out, err := a.Open(nil, nonce[:], ciphertext, nil)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(plaintext, out) {
		t.Errorf("got %q, want %q", out, plaintext)
	}
}

func BenchmarkAEAD(b *testing.B) {
	a := NewAEAD([]byte("key"))
	var nonce [15]byte
	plaintext := make([]byte, 8192)
	dst := make([]byte, 0, 8192+32)
	b.SetBytes(int64(len(plaintext)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Seal(dst, nonce[:], plaintext, nil)
	}
}
