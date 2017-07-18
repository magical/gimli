package gimli

//go:noescape
func permuteAsm(s *[48]byte)

func permute(s *[48]byte) { permuteAsm(s) }

//var permute = permuteAsm
