package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string] int)
	for k, v := range in {
		for _, s := range v {
			out[strings.ToLower(s)] = k
		}
	}
	return out
}
