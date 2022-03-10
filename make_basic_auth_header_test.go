package isgofirst

import (
	"fmt"
	"os"
	"testing"
)

type env struct {
	name     string
	value    string
	isNotSet bool
}

func updateEnvs(envs []env) error {
	for _, env := range envs {
		if env.isNotSet {
			if err := os.Unsetenv(env.name); err != nil {
				return fmt.Errorf("unable to unset an environment variable: %w", err)
			}
		} else {
			if err := os.Setenv(env.name, env.value); err != nil {
				return fmt.Errorf("unable to set an environment variable: %w", err)
			}
		}
	}

	return nil
}

func Test_makeBasicAuthHeader(test *testing.T) {
	type args struct {
		usernameEnv env
		passwordEnv env
	}
	type data struct {
		name   string
		args   args
		wanted string
	}

	tests := []data{
		data{
			name: "without the username and the password",
			args: args{
				usernameEnv: env{name: "TEST_USERNAME", isNotSet: true},
				passwordEnv: env{name: "TEST_PASSWORD", isNotSet: true},
			},
			wanted: "",
		},
		data{
			name: "with the empty username and the empty password",
			args: args{
				usernameEnv: env{name: "TEST_USERNAME", value: ""},
				passwordEnv: env{name: "TEST_PASSWORD", value: ""},
			},
			wanted: "",
		},
		data{
			name: "with the nonempty username and the nonempty password",
			args: args{
				usernameEnv: env{name: "TEST_USERNAME", value: "username"},
				passwordEnv: env{name: "TEST_PASSWORD", value: "password"},
			},
			wanted: "Basic dXNlcm5hbWU6cGFzc3dvcmQ=",
		},
	}
	for _, testData := range tests {
		err := updateEnvs([]env{testData.args.usernameEnv, testData.args.passwordEnv})
		if err != nil {
			test.Logf("failed %q: %s", testData.name, err)
			test.FailNow()
		}
		defer updateEnvs([]env{
			env{name: testData.args.usernameEnv.name, isNotSet: true},
			env{name: testData.args.passwordEnv.name, isNotSet: true},
		})

		received := makeBasicAuthHeader(
			testData.args.usernameEnv.name,
			testData.args.passwordEnv.name,
		)

		if received != testData.wanted {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wanted,
				received,
			)
			test.Fail()
		}
	}
}
