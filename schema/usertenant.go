package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserTenant holds the schema definition for the User-Tenant relationship.
// This is the junction table that connects users to tenants with their role.
type UserTenant struct {
	ent.Schema
}

// Fields of the UserTenant.
func (UserTenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty(),
		field.Int("tenant_id"),
		field.Enum("role").
			Values("admin", "operator", "user").
			Default("user").
			Comment("The role of the user within this tenant"),
		field.Bool("is_default").
			Default(false).
			Comment("If true, this is the default tenant for the user after login"),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserTenant.
func (UserTenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tenant", Tenant.Type).
			Unique().
			Required().
			Field("tenant_id"),
	}
}

// Indexes of the UserTenant.
func (UserTenant) Indexes() []ent.Index {
	return []ent.Index{
		// Ensure a user can only be assigned to a tenant once
		index.Fields("user_id", "tenant_id").Unique(),
	}
}
