package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SoftwareRepo holds the schema definition for the SoftwareRepo entity.
type SoftwareRepo struct {
	ent.Schema
}

// Fields of the SoftwareRepo.
func (SoftwareRepo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("repo_type").Values("global", "tenant"),
		field.String("endpoint").NotEmpty(),
		field.String("bucket").NotEmpty(),
		field.String("region").Optional().Default(""),
		field.String("access_key").Optional().Sensitive(),
		field.String("secret_key").Optional().Sensitive(),
		field.String("base_path").Optional().Default(""),
		field.Bool("use_presigned").Optional().Default(true),
		field.Int("presign_ttl_seconds").Optional().Default(14400).
			Comment("Pre-signed URL TTL in seconds, default 4 hours"),
		field.Bool("is_default").Optional().Default(false),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SoftwareRepo.
func (SoftwareRepo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Unique().Ref("software_repos"),
		edge.To("packages", ManagedPackage.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

// Indexes of the SoftwareRepo.
func (SoftwareRepo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_default"),
	}
}
