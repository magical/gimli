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

func benchAEAD(b *testing.B, size int) {
	a := NewAEAD([]byte("key"))
	var nonce [15]byte
	plaintext := make([]byte, size)
	dst := make([]byte, 0, size+32)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Seal(dst, nonce[:], plaintext, nil)
	}
}

func BenchmarkAEAD_8(b *testing.B)    { benchAEAD(b, 8) }
func BenchmarkAEAD_1024(b *testing.B) { benchAEAD(b, 1024) }
func BenchmarkAEAD_8192(b *testing.B) { benchAEAD(b, 8192) }
