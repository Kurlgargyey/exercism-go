package strand

func ToRNA(dna string) string {
	rna_map := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	rna_string := ""
	for _, r := range dna {
		rna_string += string(rna_map[r])
	}
	return rna_string
}
