package util

import "os"

// GetOptionalEnvironmentVar get environment variabe.
// You can return default value if the key has no value.
func GetOptionalEnvironmentVar(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}
	return v
}
