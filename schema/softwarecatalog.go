package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SoftwareCatalog holds the schema definition for the SoftwareCatalog entity.
type SoftwareCatalog struct {
	ent.Schema
}

// Fields of the SoftwareCatalog.
func (SoftwareCatalog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Comment("Ring name: test, first, fast, broad"),
		field.String("description").Optional().Default(""),
		field.Int("ring_order").
			Comment("Rollout order: 0=test, 1=first, 2=fast, 3=broad"),
		field.Bool("is_default").Optional().Default(false),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SoftwareCatalog.
func (SoftwareCatalog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Unique().Ref("software_catalogs"),
		edge.To("packages", ManagedPackage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

// Indexes of the SoftwareCatalog.
func (SoftwareCatalog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("ring_order"),
	}
}
