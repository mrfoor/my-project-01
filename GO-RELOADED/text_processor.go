package main

import (
	"regexp"
	"strconv"
	"strings"
)

func ProcessText(text string) string {
	text = ReplaceHex(text)
	text = ReplaceBin(text)
	text = ApplyModifiers(text)
	text = FixPunctuation(text)
	text = FixArticles(text)
	return text
}

func ReplaceHex(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(num, 10)
	})
}

func ReplaceBin(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 2, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(num, 10)
	})
}

func ApplyModifiers(text string) string {
	text = applyModifier(text, "up", strings.ToUpper)
	text = applyModifier(text, "low", strings.ToLower)
	text = applyModifier(text, "cap", capitalize)
	return text
}

func applyModifier(text, modifier string, transform func(string) string) string {
	re := regexp.MustCompile(`(\w+)(?:,\s*(\d+))?\s\(` + modifier + `\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		word := parts[1]
		return transform(word)
	})
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func FixPunctuation(text string) string {
	re := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	text = re.ReplaceAllString(text, "$1 ")
	text = strings.ReplaceAll(text, " ... ", "...")
	return text
}

func FixArticles(text string) string {
	re := regexp.MustCompile(`\ba\s([aeiouhAEIOUH])`)
	return re.ReplaceAllString(text, "an $1")
}
