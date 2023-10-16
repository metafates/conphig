# ⚙️ Conφg

> Spells like "con-phi-g"

A better way to manage Go apps configuration

```go
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
```

## How to use

Conphig extends awesome [koanf](https://github.com/knadh/koanf) library.

See [examples](./examples) for examples how to use this library.