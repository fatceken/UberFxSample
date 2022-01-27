package rediscache

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"os"
	"testing"
)

func TestModule(t *testing.T) {

	os.Setenv("MYPREFIX_RedisHost", "localhost:6379")
	os.Setenv("MYPREFIX_RedisPassword", "123456")

	fxtest.New(t,
		Module(),
		fx.Invoke(func(h *Helper) {
			require.NotNil(t, h.client)
		}),
	)

}
