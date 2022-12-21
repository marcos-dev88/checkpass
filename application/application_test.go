package application

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/marcos-dev88/checkpass/internal/entity"
	"github.com/marcos-dev88/checkpass/internal/service"
)

func uint32ToPointer(val uint32) *uint32 {
	return &val
}

// go test -v -failfast -count=1 -cover -run ^Test_Check$
func Test_Check(t *testing.T) {

	type fields struct {
		password *entity.PasswordType
		service  service.Verify
	}
	tests := []struct {
		name          string
		fields        fields
		isWantedErr   bool
		wantedNoMatch []string
	}{
		{
			name: "success",
			fields: fields{
				password: mockPasswordStruct(successUpperCase),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},

			wantedNoMatch: []string{},
			isWantedErr:   false,
		},
		{
			name: "error_min_size",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         errors.New("error verify: your"),
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string{"minSize"},
			isWantedErr:   false,
		},
		{
			name: "error_upper_case",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       errors.New("error verify: your"),
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string{"minUppercase"},
			isWantedErr:   false,
		},
		{
			name: "error_lower_case",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    errors.New("error verify: your"),
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string{"minLowercase"},
			isWantedErr:   false,
		},
		{
			name: "error_digit",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        errors.New("error verify: your"),
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string{"minDigit"},
			isWantedErr:   false,
		},
		{
			name: "error_spcial_chars",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: errors.New("error verify: your"),
				},
			},
			wantedNoMatch: []string{"minSpecialChars"},
			isWantedErr:   false,
		},
		{
			name: "error_go_in_min_size",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         errors.New("some_err"),
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string(nil),
			isWantedErr:   true,
		},
		{
			name: "error_go_in_upper_cases",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       errors.New("some_err"),
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string(nil),
			isWantedErr:   true,
		},
		{
			name: "error_go_in_lower_cases",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    errors.New("some_err"),
					errMinDigit:        nil,
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string(nil),
			isWantedErr:   true,
		},
		{
			name: "error_go_in_digit",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        errors.New("some_err"),
					errMinSpecialChars: nil,
				},
			},
			wantedNoMatch: []string(nil),
			isWantedErr:   true,
		},
		{
			name: "error_go_in_special_chars",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
				service: &verifyMock{
					errMinSize:         nil,
					errUpperCase:       nil,
					errMinLowerCase:    nil,
					errMinDigit:        nil,
					errMinSpecialChars: errors.New("some_err"),
				},
			},
			wantedNoMatch: []string(nil),
			isWantedErr:   true,
		},
		{
			name: "error_all_cases",
			fields: fields{
				password: mockPasswordStruct(lowSizePassword),
			},
			wantedNoMatch: []string{"minSize", "minSpecialChars", "noRepeted", "minDigit", "minUppercase", "minLowercase"},
			isWantedErr:   false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			a := &app{
				service:  testCase.fields.service,
				password: testCase.fields.password,
			}

			gotWantedNoMatch, err := a.RunCheck()

			if (testCase.isWantedErr) && err == nil {
				t.Errorf("was not suppose to have an error here and got %v", err)
			}

			if !reflect.DeepEqual(gotWantedNoMatch, testCase.wantedNoMatch) {
				t.Errorf("got %v and wanted %v", gotWantedNoMatch, testCase.wantedNoMatch)

			}

		})

	}

}

// go test -v -failfast -count=1 -cover -run ^Test_GetPassword$
func Test_GetPassword(t *testing.T) {

	p, err := NewPassword("oooooooo", []entity.RuleType{
		{
			Rule:  "minUppercase",
			Value: uint32ToPointer(8),
		},
	})

	if err != nil {
		t.Errorf("data: %v", err)
	}

	pApp := NewApp(p)

	gotPass := pApp.GetPassword()

	err = pApp.SetPassword(nil)

	if err != nil {
		t.Errorf("data: %v", err)
	}

	log.Printf("data filled: %v", gotPass)

	gotPassNil := pApp.GetPassword()

	log.Printf("data filled: %v", gotPassNil)

}

// go test -v -failfast -count=1 -cover -run ^Test_SetPassword$
func Test_SetPassword(t *testing.T) {

	p, err := NewPassword("oooooooo", []entity.RuleType{
		{
			Rule:  "minUppercase",
			Value: uint32ToPointer(8),
		},
	})

	if err != nil {
		t.Errorf("data: %v", err)
	}

	pApp := NewApp(p)

	oldPass := pApp.GetPassword()

	log.Printf("old Password = %v", oldPass)

	newPass, err := NewPassword("newPasswordBro", []entity.RuleType{
		{
			Rule:  "minUppercase",
			Value: uint32ToPointer(10),
		},
		{
			Rule:  "minLowerCase",
			Value: uint32ToPointer(8),
		},
	})

	err = pApp.SetPassword(&Password{
		Password: newPass.Password,
		Rules:    newPass.Rules,
	})

	if err != nil {
		t.Errorf("data: %v", err)
	}

	newPassword := pApp.GetPassword()

	log.Printf("new Password = %v", newPassword)

}
