package appsettings

import (
	"go.uber.org/fx"
	"reflect"
)

type AppSettings struct {
	FooSettings FooSettings
	BarSettings BarSettings
	Address     string
}

type FooSettings struct {
	Name       string
	Descripton string
}

type BarSettings struct {
	Port      string
	IsEnabled string
}

func GetType() reflect.Type {
	var o *AppSettings
	return reflect.TypeOf(o)
}

func Module() fx.Option {
	return fx.Provide(
		createSettings,
	)
}
func createSettings() *AppSettings {
	return &AppSettings{}
}
