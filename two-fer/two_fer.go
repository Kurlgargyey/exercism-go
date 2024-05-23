package twofer

import "fmt"

func ShareWith(name string) string {
	switch name {
	case "":
		return "One for you, one for me."
	default:
		return fmt.Sprintf("One for %s, one for me.", name)
	}
}
