# ⚙️ Conφg

> Spells like "con-phi-g"

A better way to manage Go apps configuration

```go
var PositiveInteger = conphig.New[int, int](
	"positive_integer",
	42,
	conphig.ID[int],
	conphig.WithDescription[int, int]("Positive integer > 0"),
	conphig.WithValidateFunc[int](func(i int) error {
		if i <= 0 {
			return fmt.Errorf("expected >= 0, got %d", i)
		}

		return nil
	}),
)

var FooBar = conphig.New[string, string](
	"foobar",
	"${EDITOR}, $USER and $HOME",
	conphig.ID[string],
	conphig.WithDescription[string, string]("String with env variables"),
	conphig.WithAdjustFunc[string, string](func(s string) (string, error) {
		return os.ExpandEnv(s), nil
	}),
)

var Time = conphig.New[string, time.Time](
	"time",
	time.Now(),
	func(s string) (time.Time, error) {
		return time.Parse(time.Kitchen, s)
	},
)
```

> [!NOTE]  
> Work in progress, API may change

## How to use

Conphig extends awesome [koanf](https://github.com/knadh/koanf) library.

See [examples](./examples) for examples how to use this library.