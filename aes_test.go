package gimli_test

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

func BenchmarkAESCBC(b *testing.B) {
	key := []byte("my 16-byte key. ")
	var iv [16]byte
	text := make([]byte, 8192)
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Error(err)
		return
	}
	mode := cipher.NewCBCEncrypter(block, iv[:])
	b.SetBytes(int64(len(text)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mode.CryptBlocks(text, text)
	}
}

func BenchmarkAESCTR(b *testing.B) {
	key := []byte("my 16-byte key. ")
	var iv [16]byte
	text := make([]byte, 8192)
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Error(err)
		return
	}
	mode := cipher.NewCTR(block, iv[:])
	b.SetBytes(int64(len(text)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mode.XORKeyStream(text, text)
	}
}

func BenchmarkAESGCM(b *testing.B) {
	key := []byte("my 16-byte key. ")
	var nonce [12]byte
	text := make([]byte, 8192)
	dst := make([]byte, 0, 8192+16)
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Error(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		b.Error(err)
		return
	}
	b.SetBytes(int64(len(text)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Seal(dst, nonce[:], text, nil)
	}
}
