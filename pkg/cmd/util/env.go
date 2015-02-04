package util

import (
	"fmt"
	"os"
	"regexp"
)

// Env returns an environment variable or a default value if not specified.
func Env(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

// GetEnv returns an environment value if specified
func GetEnv(key string) (string, bool) {
	val := os.Getenv(key)
	if len(val) == 0 {
		return "", false
	}
	return val, true
}

type Environment map[string]string

var argumentEnvironment = regexp.MustCompile("^([\\w\\-]+)\\=(.*)$")

func IsEnvironmentArgument(s string) bool {
	return argumentEnvironment.MatchString(s)
}

func ParseEnvironmentArguments(s []string) (Environment, []string, []error) {
	errs := []error{}
	duplicates := []string{}
	env := make(Environment)
	for _, s := range s {
		switch matches := argumentEnvironment.FindStringSubmatch(s); len(matches) {
		case 3:
			k, v := matches[1], matches[2]
			if exist, ok := env[k]; ok {
				duplicates = append(duplicates, fmt.Sprintf("%s=%s", k, exist))
			}
			env[k] = v
		default:
			errs = append(errs, fmt.Errorf("environment variables must be of the form key=value: %s", s))
		}
	}
	return env, duplicates, errs
}
