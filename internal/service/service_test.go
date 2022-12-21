package service

import (
	"log"
	"testing"
)

// go test -v -failfast -count=1 -cover -run ^Test_CheckNoRepeat$
func Test_CheckNoRepeat(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(noRepeatPasswordSuccess),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successDigitMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(noRepeatPasswordError),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.NoRepeat()
			log.Printf("data: %v", err)
			if (testCase.isWantedErr == true) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}
		})

	}

}

// go test -v -failfast -count=1 -cover -run ^Test_CheckDigit$
func Test_CheckSize(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(successDigit),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successDigitMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(lowSizePassword),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.MinSize()
			if (testCase.isWantedErr) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}
		})

	}

}

// go test -v -failfast -count=1 -cover -run ^Test_CheckUpperCase$
func Test_CheckUpperCase(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(successUpperCase),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successUpperCaseMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(errorUpperCase),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.MinUpperCase()

			if (testCase.isWantedErr) && err == nil {

				t.Errorf("was not suppose to have an error here and got %v", err)
			}

		})

	}
}

// go test -v -failfast -count=1 -cover -run ^Test_CheckLowerCase$
func Test_CheckLowerCase(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(successLowerCase),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successLowerCaseMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(errorLowerCase),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.MinLowerCase()
			if (testCase.isWantedErr) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}
		})
	}

}

// go test -v -failfast -count=1 -cover -run ^Test_CheckDigit$
func Test_CheckDigit(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(successDigit),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successDigitMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(errorDigit),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.MinDigit()
			if (testCase.isWantedErr) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}
		})

	}

}

// go test -v -failfast -count=1 -cover -run ^Test_CheckDigit$
func Test_CheckSpecialCharacters(t *testing.T) {

	tests := []struct {
		name        string
		v           Verify
		isWantedErr bool
	}{
		{
			name:        "success",
			v:           mockPassword(successSpecialChar),
			isWantedErr: false,
		},
		{
			name:        "success_more",
			v:           mockPassword(successSpecialCharMore),
			isWantedErr: false,
		},
		{
			name:        "error",
			v:           mockPassword(errorSpecialChar),
			isWantedErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.v.MinSpecialChars()
			if (testCase.isWantedErr) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}
		})

	}

}
