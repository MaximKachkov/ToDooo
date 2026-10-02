package core_http_types

import (
	"encoding/json"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true
	if string(b) == "null" {

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	n.Value = &value
	return nil
}

func (n *Nullable[T]) ToDomainNullable() domain.Nullable[T] {
	var domainNullable domain.Nullable[T]

	domainNullable.Set = n.Set
	domainNullable.Value = n.Value

	return domainNullable
}
