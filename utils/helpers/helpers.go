package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

func GenerateSlug(key string) string {
	key = strings.ReplaceAll(key, " ", "-")
	key = strings.ReplaceAll(key, "#", "")
	key = strings.ReplaceAll(key, "@", "")
	key = strings.ReplaceAll(key, `%`, "")
	key = strings.ReplaceAll(key, `?`, "")
	key = strings.ReplaceAll(key, `(`, "")
	key = strings.ReplaceAll(key, `)`, "")
	key = strings.ReplaceAll(key, `[`, "")
	key = strings.ReplaceAll(key, `]`, "")
	key = strings.ReplaceAll(key, `{`, "")
	key = strings.ReplaceAll(key, `}`, "")
	key = strings.ReplaceAll(key, `!`, "")
	key = strings.ReplaceAll(key, `=`, "")
	//key = regexp.MustCompile(`[^a-zA-Z0-9?:-]+`).ReplaceAllString(key, "")
	key = strings.ToLower(key)
	key = ReplaceNonUTF8(key)
	return key
}

func ReplaceNonUTF8(filename string) string  {
	re := regexp.MustCompile("[^a-zA-Z0-9./:]")
	return fmt.Sprintf(re.ReplaceAllString(filename, ""))
}

