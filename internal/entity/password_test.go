package entity

import (
	"reflect"
	"testing"
)

// go test -v -failfast -count=1 -cover -run ^Test_NewPassword$
func Test_NewPassword(t *testing.T) {
	type args struct {
		Password string
		Rules    []RuleType
	}
	tests := []struct {
		testName string
		args     args
		wantOut  *PasswordType
		wantErr  bool
	}{
		{
			testName: "success",
			args: args{
				Password: "marcos-dev88",
				Rules:    successRules,
			},
			wantOut: &PasswordType{
				Password: "marcos-dev88",
				Rules:    successRules,
			},
			wantErr: false,
		},
		{
			testName: "invalid_rule",
			args: args{
				Password: "marcos-dev88",
				Rules:    invalidRule,
			},
			wantOut: nil,
			wantErr: true,
		},
		{
			testName: "zero_rules",
			args: args{
				Password: "marcos-dev88",
				Rules:    nil,
			},
			wantOut: nil,
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			gotPassword, gotError := NewPassword(testCase.args.Password, testCase.args.Rules)

			if (!testCase.wantErr) && gotError != nil {
				t.Errorf("not wanted error and got: %v", gotError)
			}

			if !reflect.DeepEqual(gotPassword, testCase.wantOut) {
				t.Errorf("wanted %v and got: %v", testCase.wantOut, gotPassword)
			}
		})
	}
}
