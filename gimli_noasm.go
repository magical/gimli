// +build !386,!amd64

package gimli

const useAsm = false

func permute(s *[48]byte) { permuteGeneric(s) }
