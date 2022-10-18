package utils

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateNumeric() *validation.MatchRule {
	return validation.Match(regexp.MustCompile("^[0-9]*$")).Error(FIELD_NUMERIC)
}
