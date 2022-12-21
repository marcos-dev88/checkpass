package service

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/marcos-dev88/checkpass/internal/entity"
)

type Verify interface {
	MinSize() error
	MinUpperCase() error
	MinLowerCase() error
	MinDigit() error
	MinSpecialChars() error
	NoRepeat() error
}

func NewVerify(p entity.PasswordType) Verify {
	return &verify{Password: &p}
}

func (v *verify) MinSize() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.MinSize)
		if err != nil {
			return err
		}

		lPass := len(v.Password.Password)
		if uint32(lPass) < *value {
			return fmt.Errorf("error verify: your password has %d of length and minimum size is %d", lPass, *value)
		}
	}
	return nil
}

func (v *verify) MinUpperCase() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.MinUppercase)
		if err != nil {
			return err
		}

		r, err := regexp.Compile(regexUpper)

		if err != nil {
			return err
		}

		upperCasesStr := r.FindAllString(v.Password.Password, -1)
		upperCases := uint32(len(upperCasesStr))

		if upperCases < *value {
			return fmt.Errorf("error verify: your password has %d uppercases and the minimum is %d", upperCases, *value)
		}
	}

	return nil
}

func (v *verify) MinLowerCase() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.MinLowercase)
		if err != nil {
			return err
		}

		r, err := regexp.Compile(regexLower)

		if err != nil {
			return err
		}

		lowerCasesStr := r.FindAllString(v.Password.Password, -1)
		lowerCases := uint32(len(lowerCasesStr))

		if lowerCases < *value {
			return fmt.Errorf("error verify: your password has %d lowercases and the minimum is %d", lowerCases, *value)
		}
	}

	return nil
}

func (v *verify) MinDigit() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.MinDigit)
		if err != nil {
			return err
		}

		r, err := regexp.Compile(regexDigit)

		if err != nil {
			return err
		}

		digitStr := r.FindAllString(v.Password.Password, -1)
		digits := uint32(len(digitStr))

		if digits < *value {
			return fmt.Errorf("error verify: your password has %d numeric values and the minimum is %d", digits, *value)
		}
	}

	return nil
}
func (v *verify) MinSpecialChars() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.MinSpecialChars)
		if err != nil {
			return err
		}

		r, err := regexp.Compile(regexSpecialCharacter)

		if err != nil {
			return err
		}

		specialCharsStr := r.FindAllString(v.Password.Password, -1)
		specialChars := uint32(len(specialCharsStr))

		if specialChars < *value {
			return fmt.Errorf("error verify: your password has %d special characters and the minimum is %d", specialChars, *value)
		}
	}

	return nil
}

func (v *verify) NoRepeat() error {
	if v.Password != nil {
		value, err := v.getRuleValue(entity.NoRepeted)

		if err != nil {
			return err
		}

		if *value == trueNoRepeat {
			bPassword := []byte(v.Password.Password)

			for c := 0; c < len(bPassword); c++ {
				if c > 0 {
					if bPassword[c] == bPassword[c-1] {
						return errors.New("error verify: your password is not able to use repeated letters")
					}
				}
			}
		}

	}

	return nil
}

func (v *verify) getRuleValue(ruleName string) (*uint32, error) {
	if v.Password != nil {
		for c := 0; c < len(v.Password.Rules); c++ {
			if v.Password.Rules[c].Rule == ruleName {
				if v.Password.Rules[c].Value == nil {
					return nil, fmt.Errorf("error: value of rule %s is null or invalid", ruleName)
				}
				return v.Password.Rules[c].Value, nil
			}
		}
	}

	return nil, errors.New("error verify: no rule found")
}
