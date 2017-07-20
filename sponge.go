package gimli

type Sponge struct {
	buf    [48]byte
	w      int // Write position
	r      int // Read position
	padded bool
}

func (s *Sponge) Reset() {
	*s = Sponge{}
}

const rate = 16

func (s *Sponge) Write(b []byte) (written int, err error) {
	written = len(b)

	// TODO: panic if Read already called?

	// absorb
	for len(b) > 0 {
		n := len(b)
		if n > rate-s.w {
			n = rate - s.w
		}

		for i := 0; i < n; i++ {
			s.buf[s.w+i] ^= b[i]
		}

		b = b[n:]
		s.w += n
		if s.w >= rate {
			permute(&s.buf)
			s.w = 0
		}
	}

	return
}

func (s *Sponge) Read(b []byte) (read int, err error) {
	read = len(b)

	// pad
	if !s.padded {
		s.buf[s.w] ^= 0x1f
		s.buf[rate-1] ^= 0x80
		s.padded = true
		permute(&s.buf)
	}

	for len(b) > 0 {
		if s.r >= rate {
			permute(&s.buf)
			s.r = 0
		}
		n := copy(b, s.buf[s.r:rate])
		b = b[n:]
		s.r += n
	}

	return
}

func hash(in []byte, out []byte) {
	var s Sponge
	s.Write(in)
	s.Read(out)
}
