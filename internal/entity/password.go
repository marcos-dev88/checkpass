package entity

import (
	"errors"
	"strings"
)

// validation names
const (
	MinSize         = "minSize"
	MinSpecialChars = "minSpecialChars"
	NoRepeted       = "noRepeted"
	MinDigit        = "minDigit"
	MinUppercase    = "minUppercase"
	MinLowercase    = "minLowercase"
)

// validation name lower case
const (
	MinSizeL         = "minsize"
	MinSpecialCharsL = "minspecialchars"
	NoRepetedL       = "norepeted"
	MinDigitL        = "mindigit"
	MinUppercaseL    = "minuppercase"
	MinLowercaseL    = "minlowercase"
)

// NewPassword: Creates a new password structure, this is what comes from request
func NewPassword(password string, rules []RuleType) (*PasswordType, error) {
	rs, err := checkRuleList(rules)
	if err != nil {
		return nil, err
	}

	return &PasswordType{
		Password: password,
		Rules:    rs,
	}, nil
}

type PasswordType struct {
	Password string     `json:"password,omitempty"`
	Rules    []RuleType `json:"rules,omitempty"`
}

type RuleType struct {
	Rule  string  `json:"rule,omitempty"`
	Value *uint32 `json:"value,omitempty"`
}

// checkRuleList: This function check if rule exists in Password Validator project
func checkRuleList(r []RuleType) ([]RuleType, error) {

	if len(r) == 0 {
		return nil, errors.New(`you didn't set any rule to verify your password. \ntry use one of these validations: "minSize", "minSpecialChars", "noRepeted", "minDigit"`)
	}

	for c := 0; c < len(r); c++ {
		switch strings.ToLower(r[c].Rule) {
		case MinSizeL:
			r[c].Rule = MinSize
		case MinSpecialCharsL:
			r[c].Rule = MinSpecialChars
		case NoRepetedL:
			r[c].Rule = NoRepeted
		case MinDigitL:
			r[c].Rule = MinDigit
		case MinLowercaseL:
			r[c].Rule = MinLowercase
		case MinUppercaseL:
			r[c].Rule = MinUppercase

		default:
			return nil, errors.New(`this rule is not valid. \ntry use one of these validations: "minSize", "minSpecialChars", "noRepeted", "minDigit"`)
		}
	}

	return r, nil
}
