package application

import (
	"github.com/marcos-dev88/checkpass/internal/entity"
	"github.com/marcos-dev88/checkpass/internal/service"
)

const (
	MinSizeCheck int8 = iota
	MinSpecialCharsCheck
	NoRepetedCheck
	MinDigitCheck
	MinUppercaseCheck
	MinLowercaseCheck
)

var checkCases = map[string]int8{
	entity.MinSize:         MinSizeCheck,
	entity.MinSpecialChars: MinSpecialCharsCheck,
	entity.NoRepeted:       NoRepetedCheck,
	entity.MinDigit:        MinDigitCheck,
	entity.MinUppercase:    MinUppercaseCheck,
	entity.MinLowercase:    MinLowercaseCheck,
}

type (
	Password entity.PasswordType
	Rules    entity.RuleType
)

var (
	NewPassword = entity.NewPassword
	NewVerify   = service.NewVerify
)

type app struct {
	password *entity.PasswordType
	service  service.Verify
}
