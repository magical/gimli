package gimli

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
)

var _ io.Writer = &Sponge{}
var _ io.Reader = &Sponge{}

var tests = []struct {
	input string
	hash  string
}{
	{"", "b0634b2c0b082aedc5c0a2fe4ee3adcfc989ec05de6f00addb04b3aaac271f67"},
	{"There's plenty for the both of us, may the best Dwarf win.", "4afb3ff784c7ad6943d49cf5da79facfa7c4434e1ce44f5dd4b28f91a84d22c8"},
	{"If anyone was to ask for my opinion, which I note they're not, I'd say we were taking the long way around.", "ba82a16a7b224c15bed8e8bdc88903a4006bc7beda78297d96029203ef08e07c"},
	{"Speak words we can all understand!", "8dd4d132059b72f8e8493f9afb86c6d86263e7439fc64cbb361fcbccf8b01267"},
	{"It's true you don't see many Dwarf-women. And in fact, they are so alike in voice and appearance, that they are often mistaken for Dwarf-men.  And this in turn has given rise to the belief that there are no Dwarf-women, and that Dwarves just spring out of holes in the ground! Which is, of course, ridiculous.", "ebe9bfc05ce15c73336fc3c5b52b01f75cf619bb37f13bfc7f567f9d5603191a"},
}

func TestHash(t *testing.T) {
	for _, tt := range tests {
		input := []byte(tt.input)
		want, err := hex.DecodeString(tt.hash)
		if err != nil {
			t.Error(err)
			continue
		}
		got := make([]byte, 32)
		hash(input, got)
		if !bytes.Equal(got, want) {
			t.Errorf("got %x, want %x", got, want)
			t.Errorf("input %q", tt.input)
			continue
		}

		// make sure odd-sized writes work
		var s Sponge
		for i := 0; i < len(input); i++ {
			s.Reset()
			s.Write(input[:i])
			s.Write(input[i:])
			s.Read(got)
			if !bytes.Equal(got, want) {
				t.Errorf("%d/%d write: got %x, want %x", i, len(input)-i, got, want)
				t.Errorf("input %q", tt.input)
			}
		}

		// odd-sized reads
		for i := 0; i < len(got); i++ {
			s.Reset()
			s.Write(input)
			s.Read(got[:i])
			s.Read(got[i:])
			if !bytes.Equal(got, want) {
				t.Errorf("%d/%d read: got %x, want %x", i, len(got)-i, got, want)
				t.Errorf("input %q", tt.input)
			}
		}
	}
}

func benchmark(b *testing.B, size int) {
	var in = make([]byte, size)
	var out = make([]byte, 32)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash(in, out)
	}
}

func BenchmarkHash8(b *testing.B)    { benchmark(b, 8) }
func BenchmarkHash4096(b *testing.B) { benchmark(b, 4096) }
