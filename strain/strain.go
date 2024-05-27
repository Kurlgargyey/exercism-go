package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](s []T, cond func(T) bool) []T {
	selection := make([]T,0)
	for _, t := range s {
		if cond(t) {
			selection = append(selection, t)
		}
	}
	return selection
}

func Discard[T any](s []T, cond func(T) bool) []T {
	remainder := make([]T,0)
	for _, t := range s {
		if !cond(t) {
			remainder = append(remainder, t)
		}
	}
	return remainder
}