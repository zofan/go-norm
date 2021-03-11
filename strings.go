package norm

import (
	"html"
	"regexp"
	"strings"
	"unicode"
)

var (
	replaceSymbols = map[rune]rune{
		'ь': 0,
		'ъ': 0,
		'ё': 'е',
		'й': 'и',
		'щ': 'ш',
	}

	tagRe    = regexp.MustCompile(`<[^<>]+>`)
	entityRe = regexp.MustCompile(`&#?\w+;`)
	spacesRe = regexp.MustCompile(`[\p{Zs}\s]+`)
	scriptRe = regexp.MustCompile(`(?s)<script[^>]*></script>`)
	styleRe  = regexp.MustCompile(`(?s)<style[^>]*></style>`)
)

func NonStrict(v string) string {
	var n []rune

	for _, c := range v {
		if r, ok := replaceSymbols[c]; ok {
			c = r
		}

		if c > 0 && (unicode.IsLetter(c) || unicode.IsDigit(c)) {
			c = unicode.ToLower(c)
			n = append(n, c)
		}
	}

	return string(n)
}

func ClearHtml(v string) string {
	v = RemoveCode(v)
	v = StripTags(v)
	v = ReplaceEntities(v)
	v = RemoveEntities(v)
	v = CollapseSpaces(v)
	v = strings.TrimSpace(v)
	return v
}

func StripTags(v string) string {
	return tagRe.ReplaceAllString(v, ` `)
}

func RemoveEntities(v string) string {
	return entityRe.ReplaceAllString(v, ` `)
}

func ReplaceEntities(v string) string {
	v = html.UnescapeString(v)

	return CollapseSpaces(v)
}

func RemoveCode(v string) string {
	v = styleRe.ReplaceAllString(v, ` `)
	v = scriptRe.ReplaceAllString(v, ` `)
	return v
}

func CollapseSpaces(v string) string {
	return spacesRe.ReplaceAllString(v, ` `)
}
