package application

import (
	"strings"

	"github.com/marcos-dev88/checkpass/internal/entity"
)

type Application interface {
	RunCheck() ([]string, error)
	GetPassword() *Password
	SetPassword(*Password) error
	SetVerify(Password)
}

func NewApp(password *entity.PasswordType) Application {
	return &app{password: password}
}

func (a *app) RunCheck() ([]string, error) {
	if a.service == nil {
		a.service = NewVerify(*a.password)
	}

	out := make([]string, 0)

	if a.password != nil {
		for _, v := range a.password.Rules {
			switch checkCases[v.Rule] {
			case MinSizeCheck:
				if err := a.minSize(); err != nil {
					if !isVerifyError(err) {

						return nil, err
					}

					out = append(out, entity.MinSize)
				}

			case MinDigitCheck:
				if err := a.minDigit(); err != nil {
					if !isVerifyError(err) {
						return nil, err
					}

					out = append(out, entity.MinDigit)
				}

			case MinLowercaseCheck:
				if err := a.minLowerCase(); err != nil {
					if !isVerifyError(err) {
						return nil, err
					}

					out = append(out, entity.MinLowercase)
				}

			case MinUppercaseCheck:
				if err := a.minUpperCase(); err != nil {
					if !isVerifyError(err) {
						return nil, err
					}

					out = append(out, entity.MinUppercase)
				}

			case MinSpecialCharsCheck:
				if err := a.minSpecialChars(); err != nil {
					if !isVerifyError(err) {
						return nil, err
					}
					out = append(out, entity.MinSpecialChars)
				}

			case NoRepetedCheck:
				if err := a.noRepeat(); err != nil {
					if !isVerifyError(err) {
						return nil, err
					}

					out = append(out, entity.NoRepeted)
				}
			}
		}
	}

	return out, nil
}

func (a *app) GetPassword() *Password {
	if a.password != nil {
		return &Password{Password: a.password.Password, Rules: a.password.Rules}
	}
	return nil
}

func (a *app) SetPassword(password *Password) error {

	if password != nil {
		newPassword, err := NewPassword(password.Password, password.Rules)

		if err != nil {
			return err
		}

		a.password = newPassword
	} else {
		a.password = nil
	}

	return nil
}

func (a *app) SetVerify(password Password) {
	a.service = NewVerify(entity.PasswordType{Password: password.Password, Rules: password.Rules})
}

func (a *app) minSize() error {
	return a.service.MinSize()
}

func (a *app) minUpperCase() error {
	return a.service.MinUpperCase()
}
func (a *app) minLowerCase() error {
	return a.service.MinLowerCase()
}

func (a *app) minDigit() error {
	return a.service.MinDigit()
}

func (a *app) minSpecialChars() error {
	return a.service.MinSpecialChars()
}

func (a *app) noRepeat() error {
	return a.service.NoRepeat()
}

func isVerifyError(err error) bool {
	return strings.Contains(err.Error(), "error verify:")

}
