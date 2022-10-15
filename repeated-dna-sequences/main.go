package repeateddnasequences

import "errors"

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	nucleotides := map[rune]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	w := uint32(0)
	for _, r := range s[:10] {
		w <<= 2
		w |= nucleotides[r]
	}
	const mask = 1024*1024 - 1
	words := make(map[uint32]bool)
	viewed := map[uint32]bool{w: true}
	for _, r := range s[10:] {
		w <<= 2
		w = w & mask
		w |= nucleotides[r]
		if _, found := viewed[w]; found {
			words[w] = true
		} else {
			viewed[w] = true
		}
	}

	res := make([]string, 0, len(words))
	for w, _ := range words {
		res = append(res, decodeWord(w))
	}
	return res
}

func decodeWord(w uint32) string {
	var b [10]byte
	for i := 0; i < len(b); i++ {
		encodedNucleotide := w & 0b11
		w >>= 2
		var nucleotide byte
		switch encodedNucleotide {
		case 0:
			nucleotide = 'A'
		case 1:
			nucleotide = 'C'
		case 2:
			nucleotide = 'G'
		case 3:
			nucleotide = 'T'
		default:
			panic(errors.New("unexpected"))
		}
		b[len(b)-1-i] = nucleotide
	}
	return string(b[:])
}
