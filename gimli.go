package gimli

func rot(v uint32, n uint) uint32 {
	return v<<n | v>>(32-n)
}

func permute(s *[12]uint32) {
	for r := 24; r > 0; r-- {
		for j := 0; j < 4; j++ {
			x := rot(s[0+j], 24)
			y := rot(s[4+j], 9)
			z := s[8+j]
			s[8+j] = x ^ (z << 1) ^ ((y & z) << 2)
			s[4+j] = y ^ x ^ ((x | z) << 1)
			s[0+j] = z ^ y ^ ((x & y) << 3)
		}
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
