package model

import (
	"strings"
	"unicode"
)

// CaseType represents a text case format.
type CaseType string

const (
	CaseCamel          CaseType = "camel"
	CasePascal         CaseType = "pascal"
	CaseSnake          CaseType = "snake"
	CaseScreamingSnake CaseType = "screaming_snake"
	CaseKebab          CaseType = "kebab"
	CaseScreamingKebab CaseType = "screaming_kebab"
	CaseDot            CaseType = "dot"
	CaseTitle          CaseType = "title"
	CaseSentence       CaseType = "sentence"
	CaseFlat           CaseType = "flat"
	CaseTrain          CaseType = "train"
	CaseAlternating    CaseType = "alternating"
	CaseInverse        CaseType = "inverse"
)

// AllCases returns all supported case types.
func AllCases() []CaseType {
	return []CaseType{
		CaseCamel,
		CasePascal,
		CaseSnake,
		CaseScreamingSnake,
		CaseKebab,
		CaseScreamingKebab,
		CaseDot,
		CaseTitle,
		CaseSentence,
		CaseFlat,
		CaseTrain,
		CaseAlternating,
		CaseInverse,
	}
}

// CaseName returns a human-readable name for a case type.
func CaseName(c CaseType) string {
	switch c {
	case CaseCamel:
		return "camelCase"
	case CasePascal:
		return "PascalCase"
	case CaseSnake:
		return "snake_case"
	case CaseScreamingSnake:
		return "SCREAMING_SNAKE_CASE"
	case CaseKebab:
		return "kebab-case"
	case CaseScreamingKebab:
		return "SCREAMING-KEBAB-CASE"
	case CaseDot:
		return "dot.case"
	case CaseTitle:
		return "Title Case"
	case CaseSentence:
		return "Sentence case"
	case CaseFlat:
		return "flatcase"
	case CaseTrain:
		return "Train-Case"
	case CaseAlternating:
		return "aLtErNaTiNg CaSe"
	case CaseInverse:
		return "InVeRsE cAsE"
	default:
		return string(c)
	}
}

// SplitWords splits text into words regardless of the input case format.
// It handles camelCase, PascalCase, snake_case, kebab-case, dot.case,
// SCREAMING_SNAKE_CASE, and space-separated words.
func SplitWords(s string) []string {
	if s == "" {
		return nil
	}

	// First, replace common separators with spaces
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ".", " ")

	// Now split on spaces
	parts := strings.Fields(s)

	// For each part, split camelCase / PascalCase on uppercase boundaries
	var words []string
	for _, part := range parts {
		words = append(words, splitCamel(part)...)
	}

	// If no words were found (e.g. input was all separators), return the original
	if len(words) == 0 {
		return []string{s}
	}

	return words
}

// splitCamel splits a word on uppercase letter boundaries.
// e.g. "helloWorld" -> ["hello", "World"]
// e.g. "HTMLParser" -> ["HTML", "Parser"]
// e.g. "parseHTML" -> ["parse", "HTML"]
func splitCamel(s string) []string {
	if s == "" {
		return nil
	}

	var words []string
	var current strings.Builder

	runes := []rune(s)
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			// Check if previous char is lowercase — start of new word
			if unicode.IsLower(runes[i-1]) {
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			} else if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				// Current is uppercase, next is lowercase — start of new word
				// e.g. "HTMLParser" -> "HTML" | "Parser"
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}
		current.WriteRune(r)
	}

	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}

// capitalize returns the word with the first letter uppercased.
func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// Convert converts text to the specified case type.
func Convert(text string, to CaseType) string {
	// Alternating and inverse operate on the raw text, not word-split
	switch to {
	case CaseAlternating:
		return toAlternating(text)
	case CaseInverse:
		return toInverse(text)
	}

	words := SplitWords(text)
	if len(words) == 0 {
		return text
	}

	switch to {
	case CaseCamel:
		return toCamel(words)
	case CasePascal:
		return toPascal(words)
	case CaseSnake:
		return toSnake(words)
	case CaseScreamingSnake:
		return toScreamingSnake(words)
	case CaseKebab:
		return toKebab(words)
	case CaseScreamingKebab:
		return toScreamingKebab(words)
	case CaseDot:
		return toDot(words)
	case CaseTitle:
		return toTitle(words)
	case CaseSentence:
		return toSentence(words)
	case CaseFlat:
		return toFlat(words)
	case CaseTrain:
		return toTrain(words)
	default:
		return text
	}
}

func toCamel(words []string) string {
	var sb strings.Builder
	for i, w := range words {
		if i == 0 {
			sb.WriteString(strings.ToLower(w))
		} else {
			sb.WriteString(capitalize(strings.ToLower(w)))
		}
	}
	return sb.String()
}

func toPascal(words []string) string {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(capitalize(strings.ToLower(w)))
	}
	return sb.String()
}

func toSnake(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = strings.ToLower(w)
	}
	return strings.Join(parts, "_")
}

func toScreamingSnake(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = strings.ToUpper(w)
	}
	return strings.Join(parts, "_")
}

func toKebab(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = strings.ToLower(w)
	}
	return strings.Join(parts, "-")
}

func toScreamingKebab(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = strings.ToUpper(w)
	}
	return strings.Join(parts, "-")
}

func toDot(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = strings.ToLower(w)
	}
	return strings.Join(parts, ".")
}

func toTitle(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = capitalize(strings.ToLower(w))
	}
	return strings.Join(parts, " ")
}

func toSentence(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		if i == 0 {
			parts[i] = capitalize(strings.ToLower(w))
		} else {
			parts[i] = strings.ToLower(w)
		}
	}
	return strings.Join(parts, " ")
}

func toFlat(words []string) string {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(strings.ToLower(w))
	}
	return sb.String()
}

func toTrain(words []string) string {
	parts := make([]string, len(words))
	for i, w := range words {
		parts[i] = capitalize(strings.ToLower(w))
	}
	return strings.Join(parts, "-")
}

func toAlternating(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if i%2 == 0 {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(unicode.ToUpper(r))
		}
	}
	return sb.String()
}

func toInverse(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsUpper(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			sb.WriteRune(unicode.ToUpper(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// Detect attempts to identify the case type of the input text.
func Detect(text string) CaseType {
	if text == "" {
		return ""
	}

	hasSpace := strings.Contains(text, " ")
	hasUnderscore := strings.Contains(text, "_")
	hasHyphen := strings.Contains(text, "-")
	hasDot := strings.Contains(text, ".")
	hasUpper := false
	hasLower := false
	allUpper := true
	firstUpper := false

	for i, r := range text {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsLower(r) {
			hasLower = true
			allUpper = false
		}
		if i == 0 && unicode.IsUpper(r) {
			firstUpper = true
		}
	}

	// Check for alternating case
	isAlternating := true
	for i, r := range text {
		if i%2 == 0 && unicode.IsLetter(r) && unicode.IsUpper(r) {
			isAlternating = false
		}
		if i%2 == 1 && unicode.IsLetter(r) && unicode.IsLower(r) {
			isAlternating = false
		}
	}
	if isAlternating && hasUpper && hasLower && len(text) > 1 {
		return CaseAlternating
	}

	// Space-separated
	if hasSpace {
		if allUpper && hasUpper {
			return CaseTitle // Could be SCREAMING but spaces suggest title
		}
		// Check if all words start with uppercase
		words := strings.Fields(text)
		allStartUpper := true
		for _, w := range words {
			if len(w) > 0 && !unicode.IsUpper(rune(w[0])) {
				allStartUpper = false
				break
			}
		}
		if allStartUpper && len(words) > 1 {
			return CaseTitle
		}
		if firstUpper {
			return CaseSentence
		}
		return CaseSentence
	}

	// Underscore-separated
	if hasUnderscore {
		if allUpper && hasUpper {
			return CaseScreamingSnake
		}
		return CaseSnake
	}

	// Hyphen-separated
	if hasHyphen {
		if allUpper && hasUpper {
			return CaseScreamingKebab
		}
		// Check if all parts start with uppercase (Train-Case)
		parts := strings.Split(text, "-")
		allStartUpper := true
		for _, p := range parts {
			if len(p) > 0 && !unicode.IsUpper(rune(p[0])) {
				allStartUpper = false
				break
			}
		}
		if allStartUpper && len(parts) > 1 {
			return CaseTrain
		}
		return CaseKebab
	}

	// Dot-separated
	if hasDot {
		return CaseDot
	}

	// No separators — camelCase, PascalCase, or flatcase
	if hasUpper && hasLower {
		if firstUpper {
			return CasePascal
		}
		return CaseCamel
	}

	if allUpper && hasUpper {
		return CaseScreamingSnake // Single word all caps
	}

	return CaseFlat
}

// AllConversions converts text to all supported case formats and returns a map.
func AllConversions(text string) map[CaseType]string {
	result := make(map[CaseType]string)
	for _, c := range AllCases() {
		result[c] = Convert(text, c)
	}
	return result
}
