package application

import "github.com/marcos-dev88/checkpass/internal/entity"

type verifyMock struct {
	password           entity.PasswordType
	errMinSize         error
	errUpperCase       error
	errMinLowerCase    error
	errMinDigit        error
	errMinSpecialChars error
	errNoRepeat        error
}

func (v *verifyMock) MinSize() error {
	return v.errMinSize
}

func (v *verifyMock) MinUpperCase() error {
	return v.errUpperCase
}

func (v *verifyMock) MinLowerCase() error {
	return v.errMinLowerCase
}

func (v *verifyMock) MinDigit() error {
	return v.errMinDigit
}

func (v *verifyMock) MinSpecialChars() error {
	return v.errMinSpecialChars
}

func (v *verifyMock) NoRepeat() error {
	return v.errNoRepeat
}
