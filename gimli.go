package gimli

func rot(v uint32, n uint) uint32 {
	return v<<n | v>>(32-n)
}

func permute(s *[12]uint32) {
	for r := 24; r > 0; r-- {
		x0, x1, x2, x3 := s[0], s[1], s[2], s[3]
		y0, y1, y2, y3 := s[4], s[5], s[6], s[7]
		z0, z1, z2, z3 := s[8], s[9], s[10], s[11]

		x0 = rot(x0, 24)
		x1 = rot(x1, 24)
		x2 = rot(x2, 24)
		x3 = rot(x3, 24)
		y0 = rot(y0, 9)
		y1 = rot(y1, 9)
		y2 = rot(y2, 9)
		y3 = rot(y3, 9)

		s[8+0] = x0 ^ (z0 << 1) ^ ((y0 & z0) << 2)
		s[8+1] = x1 ^ (z1 << 1) ^ ((y1 & z1) << 2)
		s[8+2] = x2 ^ (z2 << 1) ^ ((y2 & z2) << 2)
		s[8+3] = x3 ^ (z3 << 1) ^ ((y3 & z3) << 2)

		s[4+0] = y0 ^ x0 ^ ((x0 | z0) << 1)
		s[4+1] = y1 ^ x1 ^ ((x1 | z1) << 1)
		s[4+2] = y2 ^ x2 ^ ((x2 | z2) << 1)
		s[4+3] = y3 ^ x3 ^ ((x3 | z3) << 1)

		s[0] = z0 ^ y0 ^ ((x0 & y0) << 3)
		s[1] = z1 ^ y1 ^ ((x1 & y1) << 3)
		s[2] = z2 ^ y2 ^ ((x2 & y2) << 3)
		s[3] = z3 ^ y3 ^ ((x3 & y3) << 3)

		if r%4 == 0 {
			s[0], s[1], s[2], s[3] = s[1], s[0], s[3], s[2]
		} else if r%4 == 2 {
			s[0], s[1], s[2], s[3] = s[2], s[3], s[0], s[1]
		}
		if r%4 == 0 {
			s[0] ^= 0x9e377900 ^ uint32(r)
		}
	}
}

const rate = 16

func hash(in []byte, out []byte) {
	var state [12]uint32
	var blockSize int

	// absorb
	for len(in) > 0 {
		blockSize = len(in)
		if blockSize > rate {
			blockSize = rate
		}

		for i := 0; i < blockSize; i++ {
			state[i/4] ^= uint32(in[i]) << (8 * (uint(i) % 4))
		}
		in = in[blockSize:]

		if blockSize == rate {
			permute(&state)
			blockSize = 0
		}
	}

	// pad
	state[blockSize/4] ^= 0x1f << (8 * (uint(blockSize) % 4))
	state[(rate-1)/4] ^= 0x80 << (8 * ((rate - 1) % 4))

	permute(&state)

	for len(out) > 0 {
		blockSize = len(out)
		if blockSize > rate {
			blockSize = rate
		}
		for i := 0; i < blockSize; i++ {
			out[i] = byte(state[i/4] >> (8 * (uint(i) % 4)))
		}
		out = out[blockSize:]
		if len(out) > 0 {
			permute(&state)
		}
	}

}
