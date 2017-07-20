package gimli

var useAsm = supportSSE2()

func supportSSE2() bool

//go:noescape
func permuteSSE2(s *[48]byte)

func permute(s *[48]byte) {
	if useAsm {
		permuteSSE2(s)
	} else {
		permuteGeneric(s)
	}
}
