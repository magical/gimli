package gimli

import "errors"

// AEAD implements the crypto.AEAD interface
type AEAD struct {
	state [48]byte
}

const duplexRate = 15
const tagsize = 32

func NewAEAD(key []byte) *AEAD {
	var aead AEAD
	s := &aead.state

	// absorb key
	w := 0
	for i := range key {
		if w >= duplexRate {
			s[duplexRate] ^= 0x83
			permute(s)
			w = 0
		}
		s[w] ^= key[i]
		w++
	}
	s[w] ^= 0x02
	s[duplexRate] ^= 0x80
	permute(s)

	return &aead
}

func (a *AEAD) NonceSize() int { return duplexRate }
func (a *AEAD) Overhead() int  { return tagsize }

func (a *AEAD) Seal(dst, nonce, plaintext, header []byte) []byte {
	s := a.state // copy state

	// spongewrap

	// absorb nonce
	w := 0
	for i := range nonce {
		if w >= duplexRate {
			s[duplexRate] ^= 0x82
			permute(&s)
			w = 0
		}
		s[w] ^= nonce[i]
		w++
	}
	s[w] ^= 0x02
	s[duplexRate] ^= 0x80
	permute(&s)

	// absorb additional data (header)
	w = 0
	for i := range header {
		if w >= duplexRate {
			s[duplexRate] ^= 0x82
			permute(&s)
			w = 0
		}
		s[w] ^= header[i]
		w++
	}
	s[w] ^= 0x03
	s[duplexRate] ^= 0x80
	permute(&s)

	// duplex plaintext
	r := 0
	for i := range plaintext {
		if r >= duplexRate {
			s[r] ^= 0x83
			permute(&s)
			r = 0
		}
		s[r] ^= plaintext[i]
		dst = append(dst, s[r])
		r++
	}
	s[r] ^= 0x02
	s[duplexRate] ^= 0x80
	permute(&s)

	for i := 0; i < tagsize; i++ {
		if r >= duplexRate {
			s[r] ^= 0x82
			permute(&s)
			r = 0
		}
		dst = append(dst, s[r])
		r++
	}

	for i := range s {
		s[i] = 0
	}

	return dst
}

func (a *AEAD) Open(dst, nonce, ciphertext, header []byte) ([]byte, error) {
	s := a.state // copy state

	// absorb nonce
	w := 0
	for i := range nonce {
		if w >= duplexRate {
			s[duplexRate] ^= 0x82
			permute(&s)
			w = 0
		}
		s[w] ^= nonce[i]
		w++
	}
	s[w] ^= 0x02
	s[duplexRate] ^= 0x80
	permute(&s)

	// absorb additional data (header)
	w = 0
	for i := range header {
		if w >= duplexRate {
			s[duplexRate] ^= 0x82
			permute(&s)
			w = 0
		}
		s[w] ^= header[i]
		w++
	}
	s[w] ^= 0x03
	s[duplexRate] ^= 0x80
	permute(&s)

	// duplex ciphertext
	r := 0
	for i := range ciphertext[:len(ciphertext)-tagsize] {
		if r >= duplexRate {
			s[r] ^= 0x83
			permute(&s)
			r = 0
		}
		dst = append(dst, ciphertext[i]^s[r])
		s[r] = ciphertext[i]
		r++
	}
	s[r] ^= 0x02
	s[duplexRate] ^= 0x80
	permute(&s)

	var tag byte
	for i := 0; i < tagsize; i++ {
		if r >= duplexRate {
			s[r] ^= 0x82
			permute(&s)
			r = 0
		}
		tag |= s[r] ^ ciphertext[len(ciphertext)-tagsize+i]
		r++
	}

	for i := range s {
		s[i] = 0
	}

	if tag != 0 {
		return dst, errors.New("gimli: invalid ciphertext")
	}

	return dst, nil
}
