package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(-|\*|=|~)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i:password).*"`)
	count:= 0
	for _, line := range lines {
		count += len(re.FindAll([]byte(line), -1))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[\d]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[[:space:]]+([A-z]*[\d]*)`)
	output := make([]string,len(lines))
	for i, line := range lines {
		if re.MatchString(line) {
			name := re.FindStringSubmatch(line)[1]
			output[i] = fmt.Sprintf("[USR] %s ", name) + line
		} else {
			output[i] = line
		}
	}
	return output
}
