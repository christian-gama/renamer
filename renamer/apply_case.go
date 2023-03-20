package renamer

import (
	"strings"
	"unicode"
)

func ApplyCase(new, old string) string {
	var result strings.Builder

	for i, oldChar := range old {
		if i >= len(new) {
			break
		}
		c := rune(new[i])
		if unicode.IsUpper(oldChar) {
			c = unicode.ToUpper(c)
		} else {
			c = unicode.ToLower(c)
		}
		result.WriteRune(c)
	}

	if len(new) > len(old) {
		result.WriteString(new[len(old):])
	}

	return result.String()
}
