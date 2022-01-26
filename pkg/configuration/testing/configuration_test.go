package testing

import (
	"embed"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"reflect"
	"testing"
	"uberfxsample/pkg/configuration"
)

//go:embed config.yaml config.*.yaml
var configurationFiles embed.FS

func TestModule(t *testing.T) {

	app := fxtest.New(t,

		fx.Invoke(func(o *configuration.Options) {
			o.ConfigurationFiles = configurationFiles
		}),

		configuration.Module(),
		fx.Invoke(func(o *configuration.Options, v *viper.Viper) {

		}),
	)
	app.RequireStart()
	app.RequireStop()
}

func TestBindConfigToOptions(t *testing.T) {

	var o *testOption
	fxtest.New(t,
		fx.Provide(func() *viper.Viper {
			v := viper.New()
			v.Set("test", map[string]interface{}{
				"prop1": "val",
			})
			return v
		}),
		fx.Provide(func() *testOption {
			return &testOption{}
		}),
		configuration.BindConfigToOptions("test", reflect.TypeOf(o)),
		fx.Invoke(func(o *testOption) {
			require.Equal(t, "val", o.Prop1)
		}),
	)
}

type testOption struct {
	Prop1 string
}
