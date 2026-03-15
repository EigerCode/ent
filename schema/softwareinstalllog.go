package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SoftwareInstallLog holds the schema definition for the SoftwareInstallLog entity.
type SoftwareInstallLog struct {
	ent.Schema
}

// Fields of the SoftwareInstallLog.
func (SoftwareInstallLog) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("action").Values("install", "update", "uninstall"),
		field.Enum("status").Values("pending", "downloading", "installing", "success", "failed"),
		field.String("error_message").Optional().Default(""),
		field.String("installed_version").Optional().Default(""),
		field.Time("started_at").Optional().Nillable(),
		field.Time("completed_at").Optional().Nillable(),
		field.Time("created").Optional().Default(time.Now),
	}
}

// Edges of the SoftwareInstallLog.
func (SoftwareInstallLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).Unique().Ref("software_install_logs").Required(),
		edge.From("package", SoftwarePackage.Type).Unique().Ref("install_logs"),
	}
}

// Indexes of the SoftwareInstallLog.
func (SoftwareInstallLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
		// Ensure only one record per agent+package (for upsert via ON CONFLICT)
		index.Edges("agent", "package").Unique(),
	}
}
