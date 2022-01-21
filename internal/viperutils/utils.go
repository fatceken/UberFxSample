package viperutils

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// BindEnv will load all environment variables with the specified prefix, and replace double underscores with "."
// This will enable environment variable support for viper.Unmarshal
func BindEnv(v *viper.Viper, prefix string) {
	envKeyMap := getEnvKeyMap(prefix)

	for envKey, viperKey := range envKeyMap {
		v.BindEnv(viperKey, envKey)
	}
}

// UnmarshalSub is a wrapper function for Viper.Unmarshal which does proper error checking when sub key does not exist
func UnmarshalSub(v *viper.Viper, key string, rawVal interface{}) error {
	v = v.Sub(key)
	if v == nil {
		return fmt.Errorf("could not load key \"%v\"", key)
	}
	return v.Unmarshal(&rawVal)
}

func getEnvKeyMap(prefix string) map[string]string {
	filterKey := func(key string) bool {
		return strings.HasPrefix(key, prefix)
	}

	translateToViperKey := func(key string) string {
		key = removePrefix(prefix, key)
		key = normalizeKey(key)
		return key
	}

	rawEnvKeys := getEnvKeys()
	envKeyMap := make(map[string]string)

	for _, k := range rawEnvKeys {
		if filterKey(k) {
			viperKey := translateToViperKey(k)
			envKeyMap[k] = viperKey
		}
	}

	return envKeyMap
}

func removePrefix(prefix string, key string) string {
	prefixLength := len(prefix)
	keyLen := len(key)
	return string(key[prefixLength:keyLen])
}

func normalizeKey(key string) string {
	const environmentDelimeter = "__"
	const viperDelimeter = "."
	const wordDelimeter = "_"

	key = strings.ReplaceAll(key, environmentDelimeter, viperDelimeter)
	key = strings.ReplaceAll(key, wordDelimeter, "")
	return key
}

func getEnvKeys() []string {
	envVars := os.Environ()
	envVarLen := len(envVars)
	envKeys := make([]string, envVarLen)

	const kvDelimeter = "="

	for i, k := range envVars {
		pair := strings.SplitN(k, kvDelimeter, 2)
		envKeys[i] = strings.TrimSpace(pair[0])
	}

	return envKeys
}
