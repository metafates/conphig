package conphig

type Value interface {
	~bool | ~int | ~string
}

type Field[T Value, G any] struct {
	key          string
	defaultValue G
	convert      func(T) (G, error)
	adjust       func(value G) (G, error)
	validate     func(value G) error
	description  string
	value        G
}

func (f Field[T, G]) Key() string {
	return f.key
}

func (f Field[T, G]) Value() G {
	if koanfInstance == nil {
		return f.defaultValue
	}

	return f.value
}

func (f Field[T, G]) DefaultValue() G {
	return f.defaultValue
}

func (f Field[T, G]) Description() string {
	return f.description
}
