package application

import (
	"encoding/json"

	"github.com/marcos-dev88/checkpass/internal/entity"
)

var lowSizePassword = `{"password":"lowpass","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted", "value":1},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

// upperCases JSONs
var (
	successUpperCase = `{"password":"11124ASf***$g*&lmfd%fFG#HHs%3301C!@O","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	successUpperCaseMore = `{"password":"1ASfgJJgmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	errorUpperCase = `{"password":"1ASfggmfdwfFGsHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`
)

// lowerCases JSONs
var (
	successLowerCase = `{"password":"1ASfggmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	successLowerCaseMore = `{"password":"1ASfgJJgmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	errorLowerCase = `{"password":"1ASffFGsHs%3COO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`
)

// digits JSONs
var (
	successDigit = `{"password":"1111111ggmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	successDigitMore = `{"password":"1A0101011001SfgJJgmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	errorDigit = `{"password":"1ASffFGsHs%3COO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`
)

// specialChars JSONs
var (
	successSpecialChar = `{"password":"$%)([])ggmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	successSpecialCharMore = `{"password":"1A%%!##$%)([])JgmfdwfFGsHHs%3CO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`

	errorSpecialChar = `{"password":"1ASffFGsHs%3COO","rules":[{"rule":"minSize","value":8},{"rule":"minSpecialChars","value":8},{"rule":"noRepeted"},{"rule":"minDigit","value":8},{"rule":"minUppercase","value":8},{"rule":"minLowercase","value":8}]}`
)

func mockPasswordStruct(pasStr string) *entity.PasswordType {

	pass := new(entity.PasswordType)
	json.Unmarshal([]byte(pasStr), pass)

	return pass
}
