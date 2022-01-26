package viperutils

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBindEnv(t *testing.T) {

	type innerConfig struct {
		InnerConfigElement string
	}

	type config struct {
		MyConfig    string
		InnerConfig *innerConfig
	}

	testEnvVars := make([]string, 2)

	testEnvVars[0] = "my config val"
	testEnvVars[1] = "my inner config val"

	os.Setenv("test_MyConfig", testEnvVars[0])
	os.Setenv("test_InnerConfig__InnerConfigElement", testEnvVars[1])

	runtimeViper := viper.New()
	BindEnv(runtimeViper, "test_")

	var c config
	runtimeViper.Unmarshal(&c)

	assert.Equal(t, c.MyConfig, testEnvVars[0])
	assert.Equal(t, c.InnerConfig.InnerConfigElement, testEnvVars[1])
}
