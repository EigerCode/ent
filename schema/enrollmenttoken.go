package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EnrollmentToken holds the schema definition for the EnrollmentToken entity.
// Enrollment tokens allow tenant admins to securely register agents to their tenant.
type EnrollmentToken struct {
	ent.Schema
}

// Fields of the EnrollmentToken.
func (EnrollmentToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			NotEmpty().
			Unique().
			Comment("The enrollment token value (UUID)"),
		field.String("description").
			Optional().
			Comment("Human-readable description for this token"),
		field.Int("max_uses").
			Default(0).
			Comment("Maximum number of times this token can be used (0 = unlimited)"),
		field.Int("current_uses").
			Default(0).
			Comment("Number of times this token has been used"),
		field.Time("expires_at").
			Optional().
			Nillable().
			Comment("When this token expires (nil = never)"),
		field.Bool("active").
			Default(true).
			Comment("Whether this token is currently active"),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the EnrollmentToken.
func (EnrollmentToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("enrollment_tokens").
			Unique().
			Required(),
		edge.From("site", Site.Type).
			Ref("enrollment_tokens").
			Unique(),
	}
}
