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
		x, y, z := sx0, sy0, sz0
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz0 = x ^ (z << 1) ^ ((y & z) << 2)
		sy0 = y ^ x ^ ((x | z) << 1)
		sx0 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx1, sy1, sz1
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz1 = x ^ (z << 1) ^ ((y & z) << 2)
		sy1 = y ^ x ^ ((x | z) << 1)
		sx1 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx2, sy2, sz2
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz2 = x ^ (z << 1) ^ ((y & z) << 2)
		sy2 = y ^ x ^ ((x | z) << 1)
		sx2 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx3, sy3, sz3
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz3 = x ^ (z << 1) ^ ((y & z) << 2)
		sy3 = y ^ x ^ ((x | z) << 1)
		sx3 = z ^ y ^ ((x & y) << 3)

		// small swap
		sx0, sx1, sx2, sx3 = sx1, sx0, sx3, sx2

		// round constant
		sx0 ^= 0x9e377900 ^ uint32(r)

		// round 3
		x, y, z = sx0, sy0, sz0
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz0 = x ^ (z << 1) ^ ((y & z) << 2)
		sy0 = y ^ x ^ ((x | z) << 1)
		sx0 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx1, sy1, sz1
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz1 = x ^ (z << 1) ^ ((y & z) << 2)
		sy1 = y ^ x ^ ((x | z) << 1)
		sx1 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx2, sy2, sz2
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz2 = x ^ (z << 1) ^ ((y & z) << 2)
		sy2 = y ^ x ^ ((x | z) << 1)
		sx2 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx3, sy3, sz3
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz3 = x ^ (z << 1) ^ ((y & z) << 2)
		sy3 = y ^ x ^ ((x | z) << 1)
		sx3 = z ^ y ^ ((x & y) << 3)

		// round 2
		x, y, z = sx0, sy0, sz0
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz0 = x ^ (z << 1) ^ ((y & z) << 2)
		sy0 = y ^ x ^ ((x | z) << 1)
		sx0 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx1, sy1, sz1
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz1 = x ^ (z << 1) ^ ((y & z) << 2)
		sy1 = y ^ x ^ ((x | z) << 1)
		sx1 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx2, sy2, sz2
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz2 = x ^ (z << 1) ^ ((y & z) << 2)
		sy2 = y ^ x ^ ((x | z) << 1)
		sx2 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx3, sy3, sz3
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz3 = x ^ (z << 1) ^ ((y & z) << 2)
		sy3 = y ^ x ^ ((x | z) << 1)
		sx3 = z ^ y ^ ((x & y) << 3)

		// big swap
		sx0, sx1, sx2, sx3 = sx2, sx3, sx0, sx1

		// round 1
		x, y, z = sx0, sy0, sz0
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz0 = x ^ (z << 1) ^ ((y & z) << 2)
		sy0 = y ^ x ^ ((x | z) << 1)
		sx0 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx1, sy1, sz1
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz1 = x ^ (z << 1) ^ ((y & z) << 2)
		sy1 = y ^ x ^ ((x | z) << 1)
		sx1 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx2, sy2, sz2
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz2 = x ^ (z << 1) ^ ((y & z) << 2)
		sy2 = y ^ x ^ ((x | z) << 1)
		sx2 = z ^ y ^ ((x & y) << 3)

		x, y, z = sx3, sy3, sz3
		x = x<<24 | x>>8
		y = y<<9 | y>>23
		sz3 = x ^ (z << 1) ^ ((y & z) << 2)
		sy3 = y ^ x ^ ((x | z) << 1)
		sx3 = z ^ y ^ ((x & y) << 3)
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
