package conphig

type Option[T Value] func(*Field[T])

// WithAdjustFunc sets adjust function for this field.
// Adjust func "adjusts" the value. You can use to modify the underlying value.
// It is called before any validation happens.
func WithAdjustFunc[T Value](adjust func(T) (T, error)) Option[T] {
	return func(f *Field[T]) {
		f.adjust = adjust
	}
}

// WithValidateFunc sets validate function for this field.
// Validate func is called after the config has been loaded.
func WithValidateFunc[T Value](validate func(T) error) Option[T] {
	return func(f *Field[T]) {
		f.validate = validate
	}
}

// WithDescription sets description for this field.
func WithDescription[T Value](description string) Option[T] {
	return func(f *Field[T]) {
		f.description = description
	}
}

// New creates a new field
func New[T Value](key string, defaultValue T, options ...Option[T]) Field[T] {
	field := Field[T]{
		key:          key,
		defaultValue: defaultValue,
		adjust: func(value T) (T, error) {
			return value, nil
		},
		validate: func(value T) error {
			return nil
		},
	}

	for _, option := range options {
		option(&field)
	}

	registeredFields = append(registeredFields, RegisteredField{
		defaultValue: field.defaultValue,
		description:  field.description,
		key:          field.key,
		validate: func() error {
			return field.validate(field.Value())
		},
		adjust: func() error {
			value, err := field.adjust(field.Value())
			if err != nil {
				return err
			}

			if koanfInstance == nil {
				field.defaultValue = value
				return nil
			}

			return koanfInstance.Set(field.key, value)
		},
	})

	return field
}
