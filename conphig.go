package conphig

import (
	"fmt"

	"github.com/knadh/koanf/v2"
)

var koanfInstance *koanf.Koanf

type Loader func(k *koanf.Koanf) error

func Load(delim string, loaders ...Loader) error {
	koanfInstance = koanf.NewWithConf(koanf.Conf{
		Delim:       delim,
		StrictMerge: true,
	})

	for _, field := range registeredFields {
		if err := koanfInstance.Set(field.key, field.defaultValue); err != nil {
			return err
		}
	}

	for _, loader := range loaders {
		loader(koanfInstance)
	}

	for _, field := range registeredFields {
		if err := field.adjust(); err != nil {
			return fmt.Errorf("%s: %w", field.key, err)
		}

		if err := field.validate(); err != nil {
			return fmt.Errorf("%s: %w", field.key, err)
		}
	}

	return nil
}
