// +build !amd64

package gimli

const useAsm = true

func permute(s *[48]byte) { permuteGeneric(s) }
