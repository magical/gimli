package gimli

const useAsm = true

//go:noescape
func permute(s *[48]byte)
