package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/language"
)

var (
	AmericanEnglish    = language.AmericanEnglish
	ArgentinianSpanish = language.Make("es-AR")
)

var (
	supported = []language.Tag{
		AmericanEnglish,
		ArgentinianSpanish,
	}
)

func main() {
	fmt.Println(resolveLanguageTag("zh-CN,zh;q=0.9,en;q=0.8"))

	var v float32
	v = 1
	t := test{a: &v}
	float(&t)

	tmp, err := strconv.ParseInt("", 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmp)
}

func resolveLanguageTag(prefs string) language.Tag {
	parts := strings.Split(prefs, ",")
	langs := make([]language.Tag, 0, len(parts))
	for _, e := range parts {
		langs = append(langs, language.Make(e))
	}

	matcher := language.NewMatcher(supported)
	matched, _, _ := matcher.Match(langs...)
	return matched
}

type test struct {
	a *float32
}

var a, b float32

func float(t *test) {
	a = 1
	b = 2
	*t.a = a + b
	fmt.Println(*t.a)
}
