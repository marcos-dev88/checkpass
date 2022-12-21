package service

import "github.com/marcos-dev88/checkpass/internal/entity"

const (
	regexUpper            = "[A-Z]"
	regexLower            = "[a-z]"
	regexDigit            = "[0-9]"
	regexSpecialCharacter = "[^a-zA-Z0-9]"
)

const trueNoRepeat = 1

type verify struct {
	Password *entity.PasswordType
}
