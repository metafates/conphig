package main

import (
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/metafates/conphig"
)

var PositiveInteger = conphig.New[int](
	"positive_integer",
	42,
	conphig.WithDescription[int]("Positive integer > 0"),
	conphig.WithValidateFunc[int](func(i int) error {
		if i <= 0 {
			return fmt.Errorf("expected >= 0, got %d", i)
		}

		return nil
	}),
)

var FooBar = conphig.New[string](
	"foobar",
	"${EDITOR}, $USER and $HOME",
	conphig.WithDescription[string]("String with env variables"),
	conphig.WithAdjustFunc[string](func(s string) (string, error) {
		return os.ExpandEnv(s), nil
	}),
)

func main() {
	err := conphig.Load(".",
		func(k *koanf.Koanf) error {
			// Load default values using the confmap provider.
			// We provide a flat map with the "." delimiter.
			// A nested map can be loaded by setting the delimiter to an empty string "".
			return k.Load(confmap.Provider(map[string]interface{}{
				"positive_integer": 11,
			}, "."), nil)
		},
		func(k *koanf.Koanf) error {
			// Load JSON config on top of the default values.
			return k.Load(file.Provider("conf.json"), json.Parser())
		},
		func(k *koanf.Koanf) error {
			// Load YAML config and merge into the previously loaded config (because we can).
			return k.Load(file.Provider("conf.yml"), yaml.Parser())
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(PositiveInteger.Value())
	fmt.Println(FooBar.Value())
}
