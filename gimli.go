package gimli

func permute(s *[48]uint8) {
	sx0 := uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 | uint32(s[3])<<24
	sx1 := uint32(s[4]) | uint32(s[5])<<8 | uint32(s[6])<<16 | uint32(s[7])<<24
	sx2 := uint32(s[8]) | uint32(s[9])<<8 | uint32(s[10])<<16 | uint32(s[11])<<24
	sx3 := uint32(s[12]) | uint32(s[13])<<8 | uint32(s[14])<<16 | uint32(s[15])<<24

	sy0 := uint32(s[16]) | uint32(s[17])<<8 | uint32(s[18])<<16 | uint32(s[19])<<24
	sy1 := uint32(s[20]) | uint32(s[21])<<8 | uint32(s[22])<<16 | uint32(s[23])<<24
	sy2 := uint32(s[24]) | uint32(s[25])<<8 | uint32(s[26])<<16 | uint32(s[27])<<24
	sy3 := uint32(s[28]) | uint32(s[29])<<8 | uint32(s[30])<<16 | uint32(s[31])<<24

	sz0 := uint32(s[32]) | uint32(s[33])<<8 | uint32(s[34])<<16 | uint32(s[35])<<24
	sz1 := uint32(s[36]) | uint32(s[37])<<8 | uint32(s[38])<<16 | uint32(s[39])<<24
	sz2 := uint32(s[40]) | uint32(s[41])<<8 | uint32(s[42])<<16 | uint32(s[43])<<24
	sz3 := uint32(s[44]) | uint32(s[45])<<8 | uint32(s[46])<<16 | uint32(s[47])<<24

	for r := 24; r > 0; r -= 4 {
		// round 4
		x0, x1, x2, x3 := sx0, sx1, sx2, sx3
		y0, y1, y2, y3 := sy0, sy1, sy2, sy3
		z0, z1, z2, z3 := sz0, sz1, sz2, sz3

		x0 = x0<<24 | x0>>8
		x1 = x1<<24 | x1>>8
		x2 = x2<<24 | x2>>8
		x3 = x3<<24 | x3>>8
		y0 = y0<<9 | y0>>23
		y1 = y1<<9 | y1>>23
		y2 = y2<<9 | y2>>23
		y3 = y3<<9 | y3>>23

		sz0 = x0 ^ (z0 << 1) ^ ((y0 & z0) << 2)
		sz1 = x1 ^ (z1 << 1) ^ ((y1 & z1) << 2)
		sz2 = x2 ^ (z2 << 1) ^ ((y2 & z2) << 2)
		sz3 = x3 ^ (z3 << 1) ^ ((y3 & z3) << 2)

		sy0 = y0 ^ x0 ^ ((x0 | z0) << 1)
		sy1 = y1 ^ x1 ^ ((x1 | z1) << 1)
		sy2 = y2 ^ x2 ^ ((x2 | z2) << 1)
		sy3 = y3 ^ x3 ^ ((x3 | z3) << 1)

		sx0 = z0 ^ y0 ^ ((x0 & y0) << 3)
		sx1 = z1 ^ y1 ^ ((x1 & y1) << 3)
		sx2 = z2 ^ y2 ^ ((x2 & y2) << 3)
		sx3 = z3 ^ y3 ^ ((x3 & y3) << 3)

		// small swap
		sx0, sx1, sx2, sx3 = sx1, sx0, sx3, sx2

		// round constant
		sx0 ^= 0x9e377900 ^ uint32(r)

		// round 3
		x0, x1, x2, x3 = sx0, sx1, sx2, sx3
		y0, y1, y2, y3 = sy0, sy1, sy2, sy3
		z0, z1, z2, z3 = sz0, sz1, sz2, sz3

		x0 = x0<<24 | x0>>8
		x1 = x1<<24 | x1>>8
		x2 = x2<<24 | x2>>8
		x3 = x3<<24 | x3>>8
		y0 = y0<<9 | y0>>23
		y1 = y1<<9 | y1>>23
		y2 = y2<<9 | y2>>23
		y3 = y3<<9 | y3>>23

		sz0 = x0 ^ (z0 << 1) ^ ((y0 & z0) << 2)
		sz1 = x1 ^ (z1 << 1) ^ ((y1 & z1) << 2)
		sz2 = x2 ^ (z2 << 1) ^ ((y2 & z2) << 2)
		sz3 = x3 ^ (z3 << 1) ^ ((y3 & z3) << 2)

		sy0 = y0 ^ x0 ^ ((x0 | z0) << 1)
		sy1 = y1 ^ x1 ^ ((x1 | z1) << 1)
		sy2 = y2 ^ x2 ^ ((x2 | z2) << 1)
		sy3 = y3 ^ x3 ^ ((x3 | z3) << 1)

		sx0 = z0 ^ y0 ^ ((x0 & y0) << 3)
		sx1 = z1 ^ y1 ^ ((x1 & y1) << 3)
		sx2 = z2 ^ y2 ^ ((x2 & y2) << 3)
		sx3 = z3 ^ y3 ^ ((x3 & y3) << 3)

		// round 2
		x0, x1, x2, x3 = sx0, sx1, sx2, sx3
		y0, y1, y2, y3 = sy0, sy1, sy2, sy3
		z0, z1, z2, z3 = sz0, sz1, sz2, sz3

		x0 = x0<<24 | x0>>8
		x1 = x1<<24 | x1>>8
		x2 = x2<<24 | x2>>8
		x3 = x3<<24 | x3>>8
		y0 = y0<<9 | y0>>23
		y1 = y1<<9 | y1>>23
		y2 = y2<<9 | y2>>23
		y3 = y3<<9 | y3>>23

		sz0 = x0 ^ (z0 << 1) ^ ((y0 & z0) << 2)
		sz1 = x1 ^ (z1 << 1) ^ ((y1 & z1) << 2)
		sz2 = x2 ^ (z2 << 1) ^ ((y2 & z2) << 2)
		sz3 = x3 ^ (z3 << 1) ^ ((y3 & z3) << 2)

		sy0 = y0 ^ x0 ^ ((x0 | z0) << 1)
		sy1 = y1 ^ x1 ^ ((x1 | z1) << 1)
		sy2 = y2 ^ x2 ^ ((x2 | z2) << 1)
		sy3 = y3 ^ x3 ^ ((x3 | z3) << 1)

		sx0 = z0 ^ y0 ^ ((x0 & y0) << 3)
		sx1 = z1 ^ y1 ^ ((x1 & y1) << 3)
		sx2 = z2 ^ y2 ^ ((x2 & y2) << 3)
		sx3 = z3 ^ y3 ^ ((x3 & y3) << 3)

		// big swap
		sx0, sx1, sx2, sx3 = sx2, sx3, sx0, sx1

		// round 1
		x0, x1, x2, x3 = sx0, sx1, sx2, sx3
		y0, y1, y2, y3 = sy0, sy1, sy2, sy3
		z0, z1, z2, z3 = sz0, sz1, sz2, sz3

		x0 = x0<<24 | x0>>8
		x1 = x1<<24 | x1>>8
		x2 = x2<<24 | x2>>8
		x3 = x3<<24 | x3>>8
		y0 = y0<<9 | y0>>23
		y1 = y1<<9 | y1>>23
		y2 = y2<<9 | y2>>23
		y3 = y3<<9 | y3>>23

		sz0 = x0 ^ (z0 << 1) ^ ((y0 & z0) << 2)
		sz1 = x1 ^ (z1 << 1) ^ ((y1 & z1) << 2)
		sz2 = x2 ^ (z2 << 1) ^ ((y2 & z2) << 2)
		sz3 = x3 ^ (z3 << 1) ^ ((y3 & z3) << 2)

		sy0 = y0 ^ x0 ^ ((x0 | z0) << 1)
		sy1 = y1 ^ x1 ^ ((x1 | z1) << 1)
		sy2 = y2 ^ x2 ^ ((x2 | z2) << 1)
		sy3 = y3 ^ x3 ^ ((x3 | z3) << 1)

		sx0 = z0 ^ y0 ^ ((x0 & y0) << 3)
		sx1 = z1 ^ y1 ^ ((x1 & y1) << 3)
		sx2 = z2 ^ y2 ^ ((x2 & y2) << 3)
		sx3 = z3 ^ y3 ^ ((x3 & y3) << 3)
	}

	s[0], s[1], s[2], s[3] = byte(sx0), byte(sx0>>8), byte(sx0>>16), byte(sx0>>24)
	s[4], s[5], s[6], s[7] = byte(sx1), byte(sx1>>8), byte(sx1>>16), byte(sx1>>24)
	s[8], s[9], s[10], s[11] = byte(sx2), byte(sx2>>8), byte(sx2>>16), byte(sx2>>24)
	s[12], s[13], s[14], s[15] = byte(sx3), byte(sx3>>8), byte(sx3>>16), byte(sx3>>24)

	s[16], s[17], s[18], s[19] = byte(sy0), byte(sy0>>8), byte(sy0>>16), byte(sy0>>24)
	s[20], s[21], s[22], s[23] = byte(sy1), byte(sy1>>8), byte(sy1>>16), byte(sy1>>24)
	s[24], s[25], s[26], s[27] = byte(sy2), byte(sy2>>8), byte(sy2>>16), byte(sy2>>24)
	s[28], s[29], s[30], s[31] = byte(sy3), byte(sy3>>8), byte(sy3>>16), byte(sy3>>24)

	s[32], s[33], s[34], s[35] = byte(sz0), byte(sz0>>8), byte(sz0>>16), byte(sz0>>24)
	s[36], s[37], s[38], s[39] = byte(sz1), byte(sz1>>8), byte(sz1>>16), byte(sz1>>24)
	s[40], s[41], s[42], s[43] = byte(sz2), byte(sz2>>8), byte(sz2>>16), byte(sz2>>24)
	s[44], s[45], s[46], s[47] = byte(sz3), byte(sz3>>8), byte(sz3>>16), byte(sz3>>24)
}

const rate = 16

func hash(in []byte, out []byte) {
	var state [48]uint8
	var blockSize int

	// absorb
	for len(in) > 0 {
		blockSize = len(in)
		if blockSize > rate {
			blockSize = rate
		}

		for i := 0; i < blockSize; i++ {
			state[i] ^= in[i]
		}
		in = in[blockSize:]

		if blockSize == rate {
			permute(&state)
			blockSize = 0
		}
	}

	// pad
	state[blockSize] ^= 0x1f
	state[rate-1] ^= 0x80

	permute(&state)

	for len(out) > 0 {
		blockSize = len(out)
		if blockSize > rate {
			blockSize = rate
		}
		for i := 0; i < blockSize; i++ {
			out[i] = state[i]
		}
		out = out[blockSize:]
		if len(out) > 0 {
			permute(&state)
		}
	}

}
