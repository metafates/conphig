package conphig

import (
	"fmt"
)

type Option[T Value, G any] func(*Field[T, G])

// WithAdjustFunc sets adjust function for this field.
// Adjust func "adjusts" the value. You can use to modify the underlying value.
// It is called before any validation happens.
func WithAdjustFunc[T Value, G any](adjust func(G) (G, error)) Option[T, G] {
	return func(f *Field[T, G]) {
		f.adjust = adjust
	}
}

// WithValidateFunc sets validate function for this field.
// Validate func is called after the config has been loaded.
func WithValidateFunc[T Value, G any](validate func(G) error) Option[T, G] {
	return func(f *Field[T, G]) {
		f.validate = validate
	}
}

// WithDescription sets description for this field.
func WithDescription[T Value, G any](description string) Option[T, G] {
	return func(f *Field[T, G]) {
		f.description = description
	}
}

// New creates a new field
func New[T Value, G any](
	key string,
	defaultValue G,
	convert func(T) (G, error),
	options ...Option[T, G],
) *Field[T, G] {
	field := &Field[T, G]{
		key:          key,
		defaultValue: defaultValue,
		value:        defaultValue,
		convert:      convert,
		adjust: func(value G) (G, error) {
			return value, nil
		},
		validate: func(value G) error {
			return nil
		},
	}

	for _, option := range options {
		option(field)
	}

	registeredFields = append(registeredFields, RegisteredField{
		defaultValue: field.defaultValue,
		description:  field.description,
		key:          field.key,
		convert: func() error {
			if koanfInstance == nil {
				return nil
			}

			raw, ok := koanfInstance.Get(field.key).(T)
			if !ok {
				var t T
				return fmt.Errorf("expected %T, got %T", t, raw)
			}

			value, err := field.convert(raw)
			if err != nil {
				return err
			}

			field.value = value
			return nil
		},
		validate: func() error {
			return field.validate(field.Value())
		},
		adjust: func() error {
			value, err := field.adjust(field.Value())
			if err != nil {
				return err
			}

			field.value = value
			return nil
		},
	})

	return field
}
