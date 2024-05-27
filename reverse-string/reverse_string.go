package reverse

func Reverse(input string) string {
	reverse := ""
	for _, r := range input {
		reverse = string(r)+reverse
	}
	return reverse
}
