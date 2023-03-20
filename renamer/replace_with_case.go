package renamer

import "strings"

func ReplaceWithCase(s, old, new string) string {
	var result strings.Builder
	rest := s

	for len(rest) > 0 {
		index := strings.Index(strings.ToLower(rest), strings.ToLower(old))
		if index == -1 {
			result.WriteString(rest)
			break
		}

		result.WriteString(rest[:index])
		result.WriteString(ApplyCase(new, rest[index:index+len(old)]))
		rest = rest[index+len(old):]
	}

	return result.String()
}
