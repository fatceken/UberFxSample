package configuration

import (
	"embed"
	"fmt"
	"go.uber.org/fx"
	"io"
	"io/fs"
	"log"
	"reflect"
	"uberfxsample/internal/viperutils"

	"github.com/spf13/viper"
)

type Options struct {
	ConfigurationFiles embed.FS
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			createOptions,
			createConfig,
		),
	)

}

func createOptions() *Options {
	return &Options{}
}

func BindConfigToOptions(configKey string, optionsType reflect.Type) fx.Option {
	var v *viper.Viper
	viperType := reflect.TypeOf(v)
	var err *error
	errType := reflect.TypeOf(err).Elem()

	funcType := reflect.FuncOf([]reflect.Type{viperType, optionsType}, []reflect.Type{errType}, false)
	invokeFunc := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		log.Printf("loading options for " + optionsType.String())
		optPtr := reflect.New(optionsType).Elem()
		optPtr.Set(args[1])
		args = []reflect.Value{
			args[0],
			reflect.ValueOf(configKey),
			optPtr,
		}
		log.Printf("args %v", args)
		return reflect.ValueOf(viperutils.UnmarshalSub).Call(args)
	})

	return fx.Invoke(invokeFunc.Interface())
}

func createConfig(options *Options) (*viper.Viper, error) {
	const stageKey = "STAGE"

	viperConfig := viper.New()
	viperConfig.SetDefault(stageKey, "local")
	err := viperConfig.BindEnv(stageKey)
	if err != nil {
		return nil, fmt.Errorf("failed to bind env key %v: %w", stageKey, err)
	}

	viperutils.BindEnv(viperConfig, "MYPREFIX_") // bind env variable that starts with "MYPREFIX_"

	files, err := openEmbeddedFiles(options.ConfigurationFiles,
		"config.yaml",
		fmt.Sprintf("config.%v.yaml", viperConfig.Get(stageKey)))

	if err != nil {
		return nil, err
	}
	err = mergeYAMLConfigurations(viperConfig, toIOReader(files))

	closeFiles(files)

	if err != nil {
		return nil, err
	}

	return viperConfig, nil
}

func toIOReader(files []fs.File) []io.Reader {
	var readers []io.Reader
	for _, f := range files {
		readers = append(readers, f)
	}
	return readers
}

func closeFiles(files []fs.File) {
	for _, f := range files {
		err := f.Close()
		if err != nil {
			log.Printf("warn: could not close file: %e\n", err)
		}
	}
}

func openEmbeddedFiles(configurations embed.FS, files ...string) ([]fs.File, error) {
	var fileStreams []fs.File

	for _, f := range files {
		fileStream, err := configurations.Open(f)
		if err != nil {
			return nil, fmt.Errorf("failed to load file %v with the following error: %w", f, err)
		}
		fileStreams = append(fileStreams, fileStream)
	}

	return fileStreams, nil
}

func mergeYAMLConfigurations(root *viper.Viper, in []io.Reader) error {
	for _, r := range in {
		viperConfig := viper.New()
		viperConfig.SetConfigType("yaml")

		err := viperConfig.ReadConfig(r)
		if err != nil {
			return fmt.Errorf("could not read configuration from file: %w", err)
		}

		err = root.MergeConfigMap(viperConfig.AllSettings())
		if err != nil {
			return fmt.Errorf("failed to merge YAML configurations: %w", err)
		}
	}
	return nil
}
