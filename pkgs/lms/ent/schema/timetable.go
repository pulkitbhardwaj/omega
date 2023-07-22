package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TimeTable holds the schema definition for the TimeTable entity.
type TimeTable struct {
	ent.Schema
}

// Fields of the TimeTable.
func (TimeTable) Fields() []ent.Field {
	return []ent.Field{
		field.Int("period").
			Min(1).
			Max(9),

		field.String("monday"),
		
		field.String("tuesday"),

		field.String("wednesday"),

		field.String("thursday"),

		field.String("friday"),

	}
}

// Edges of the TimeTable.
func (TimeTable) Edges() []ent.Edge {
	return nil
}
