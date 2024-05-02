package valueObject

import (
	"errors"
	"strings"

	"golang.org/x/exp/slices"
)

type HtmlType string

var validHtmlTypes = []string{
	"checkbock",
	"color",
	"date",
	"datetime-local",
	"email",
	"image",
	"month",
	"number",
	"password",
	"radio",
	"range",
	"search",
	"tel",
	"text",
	"time",
	"url",
	"week",
}

func NewHtmlType(value string) (HtmlType, error) {
	value = strings.ToLower(value)
	if !slices.Contains(validHtmlTypes, value) {
		return "", errors.New("InvalidHtmlType")
	}

	return HtmlType(value), nil
}

func NewHtmlTypePanic(value string) HtmlType {
	ht, err := NewHtmlType(value)
	if err != nil {
		panic(err)
	}

	return ht
}

func (ht HtmlType) String() string {
	return string(ht)
}