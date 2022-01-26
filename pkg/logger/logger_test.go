package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"testing"
)

func TestModule(t *testing.T) {
	app := fxtest.New(t,
		Module(),
		fx.Invoke(func(logger *zap.SugaredLogger) {

		}),
	)
	app.RequireStart()
	app.RequireStop()
}
