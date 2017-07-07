package gimli

type Hash struct {
	buf    [48]byte
	w      int // Write position
	r      int // Read position
	padded bool
}

func (h *Hash) Reset() {
	*h = Hash{}
}

const rate = 16

func (h *Hash) Write(b []byte) (written int, err error) {
	written = len(b)

	// TODO: panic if Read already called?

	// absorb
	for len(b) > 0 {
		n := len(b)
		if n > rate-h.w {
			n = rate - h.w
		}

		for i := 0; i < n; i++ {
			h.buf[h.w+i] ^= b[i]
		}

		b = b[n:]
		h.w += n
		if h.w >= rate {
			permute(&h.buf)
			h.w = 0
		}
	}

	return
}

func (h *Hash) Read(b []byte) (read int, err error) {
	read = len(b)

	// pad
	if !h.padded {
		h.buf[h.w] ^= 0x1f
		h.buf[rate-1] ^= 0x80
		h.padded = true
		permute(&h.buf)
	}

	for len(b) > 0 {
		if h.r >= rate {
			permute(&h.buf)
			h.r = 0
		}
		n := copy(b, h.buf[h.r:rate])
		b = b[n:]
		h.r += n
	}

	return
}

func hash(in []byte, out []byte) {
	var h Hash
	h.Write(in)
	h.Read(out)
}
