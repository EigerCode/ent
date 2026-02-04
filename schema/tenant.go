package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional(),
		field.Bool("is_default").Optional(),
		field.String("oidc_org_id").
			Optional().
			Comment("OIDC organization ID (e.g. Zitadel Org-ID) for automatic tenant mapping"),
		field.Enum("oidc_default_role").
			Values("admin", "operator", "user").
			Default("user").
			Optional().
			Comment("Default role for users auto-assigned via OIDC"),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sites", Site.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("settings", Settings.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("tags", Tag.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("metadata", OrgMetadata.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("rustdesk", Rustdesk.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("netbird", NetbirdSettings.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("user_tenants", UserTenant.Type).Ref("tenant"),
		edge.To("enrollment_tokens", EnrollmentToken.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
