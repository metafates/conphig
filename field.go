package conphig

type Value interface {
	~bool | ~int | ~string
}

type Field[T Value] struct {
	key          string
	defaultValue T
	adjust       func(value T) (T, error)
	validate     func(value T) error
	description  string
}

func (f Field[T]) Key() string {
	return f.key
}

func (f Field[T]) Value() T {
	if koanfInstance == nil {
		return f.defaultValue
	}

	return koanfInstance.Get(f.key).(T)
}

func (f Field[T]) DefaultValue() T {
	return f.defaultValue
}

func (f Field[T]) Description() string {
	return f.description
}
