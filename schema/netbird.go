package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Netbird holds the schema definition for the Netbird entity.
type Netbird struct {
	ent.Schema
}

// Fields of the Netbird.
func (Netbird) Fields() []ent.Field {
	return []ent.Field{
		field.String("version").Default(""),
		field.Bool("installed").Default(false),
	}
}

// Edges of the Netbird.
func (Netbird) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("netbird").Required(),
	}
}
