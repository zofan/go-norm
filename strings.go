package norm

import (
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
