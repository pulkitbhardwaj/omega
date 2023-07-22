package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Obervation holds the schema definition for the Obervation entity.
type Obervation struct {
	ent.Schema
}

// Fields of the Obervation.
func (Obervation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("quarter").
		NamedValues(
			"1", "APR-JUN",
			"2", "SEP-OCT",
			"3", "NOV-DEC",
			"4", "JAN-FEB",
		),

		
	}
}

// Edges of the Obervation.
func (Obervation) Edges() []ent.Edge {
	return nil
}
