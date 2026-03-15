package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SoftwareAssignment holds the schema definition for the SoftwareAssignment entity.
type SoftwareAssignment struct {
	ent.Schema
}

// Fields of the SoftwareAssignment.
func (SoftwareAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("package_name").NotEmpty().
			Comment("Munki/CIMIAN package name (e.g. Firefox)"),
		field.Enum("package_platform").Values("darwin", "windows").
			Comment("Platform of the package"),
		field.Enum("assignment_type").
			Values("managed_install", "managed_uninstall", "optional_install", "managed_update"),
		field.Enum("target_type").Values("site", "tag", "agent").
			Comment("What kind of target this assignment applies to"),
		field.String("target_id").NotEmpty().
			Comment("ID of the target (site ID, tag ID, or agent ID)"),
		field.Int("priority").Optional().Default(0).
			Comment("Higher priority wins when there are conflicts"),
		field.Text("condition_predicate").Optional().Default("").
			Comment("Optional JSON condition, e.g. OS version, architecture"),
		field.Bool("active").Optional().Default(true),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SoftwareAssignment.
func (SoftwareAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Unique().Ref("software_assignments"),
	}
}

// Indexes of the SoftwareAssignment.
func (SoftwareAssignment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("target_type", "target_id"),
		index.Fields("package_name", "package_platform", "target_type", "target_id"),
	}
}
