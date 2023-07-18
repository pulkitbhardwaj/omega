package schema

import "entgo.io/ent"

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Node{},
	}
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return nil
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return nil
}
