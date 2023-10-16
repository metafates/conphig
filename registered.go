package conphig

import (
	"fmt"
	"strings"
)

var registeredFields []RegisteredField

func Registered() []RegisteredField {
	return registeredFields
}

type RegisteredField struct {
	defaultValue any
	validate     func() error
	adjust       func() error
	key          string
	description  string
}

func (r RegisteredField) String() string {
	var sb strings.Builder

	sb.WriteString(r.key)
	sb.WriteString(" -> ")
	sb.WriteString(fmt.Sprint(r.Value()))

	return sb.String()
}

func (r RegisteredField) Value() any {
	return nil
}

func (r RegisteredField) DefaultValue() any {
	return nil
}

func (r RegisteredField) Key() string {
	return r.key
}

func (r RegisteredField) Description() string {
	return r.description
}
