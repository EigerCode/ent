package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ManagedPackage holds the schema definition for the ManagedPackage entity.
type ManagedPackage struct {
	ent.Schema
}

// Fields of the ManagedPackage.
func (ManagedPackage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("display_name").Optional().Default(""),
		field.String("version").NotEmpty(),
		field.Enum("platform").Values("darwin", "windows"),
		field.String("installer_path").NotEmpty().
			Comment("Path within the S3 bucket, e.g. darwin/Firefox-130.0.pkg"),
		field.String("checksum_sha256").Optional().Default(""),
		field.Int64("size_bytes").Optional().Default(0),
		field.String("icon_name").Optional().Default(""),
		field.String("description").Optional().Default(""),
		field.String("category").Optional().Default(""),
		field.String("developer").Optional().Default(""),
		field.Text("pkginfo_data").Optional().Default("").
			Comment("Full pkgsinfo as JSON for Munki/CIMIAN compatibility"),
		field.Text("pre_install_script").Optional().Default(""),
		field.Text("post_install_script").Optional().Default(""),
		field.String("uninstall_method").Optional().Default(""),
		field.Text("installs_items").Optional().Default("").
			Comment("JSON array of items to check for installation status"),
		field.Text("receipts").Optional().Default("").
			Comment("JSON array of receipts to check"),
		field.Text("blocking_apps").Optional().Default("").
			Comment("JSON array of apps that block installation"),
		field.Enum("restart_action").Values("none", "RequireRestart", "RequireLogout", "RecommendRestart").
			Optional().Default("none"),
		field.String("min_os_version").Optional().Default(""),
		field.String("max_os_version").Optional().Default(""),
		field.Text("supported_architectures").Optional().Default("").
			Comment("JSON array, e.g. [\"x86_64\", \"arm64\"]"),
		field.Time("force_install_date").Optional().Nillable(),
		field.Bool("unattended_install").Optional().Default(false),
		field.Bool("unattended_uninstall").Optional().Default(false),
		field.Enum("status").Values("uploading", "ready", "error").Optional().Default("ready").
			Comment("Upload status: uploading = S3 transfer in progress, ready = available, error = upload failed"),
		field.Enum("source").Values("upload", "global", "global_subscription").Optional().Default("upload").
			Comment("upload = own package, global = legacy import copy, global_subscription = subscribed to global package"),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ManagedPackage.
func (ManagedPackage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repo", SoftwareRepo.Type).Unique().Ref("packages"),
		edge.From("catalogs", SoftwareCatalog.Type).Ref("packages"),
		edge.From("tenant", Tenant.Type).Unique().Ref("software_packages"),
		edge.To("install_logs", SoftwareInstallLog.Type),
		edge.To("requires", ManagedPackage.Type).
			Comment("Dependencies: packages that must be installed first"),
		edge.To("update_for", ManagedPackage.Type).
			Comment("This package is an update for the referenced packages"),
		edge.To("subscribers", ManagedPackage.Type).
			From("global_ref").Unique().
			Comment("Reference to the global package this subscription points to"),
	}
}

// Indexes of the ManagedPackage.
func (ManagedPackage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "version", "platform"),
		index.Fields("status"),
		index.Fields("platform"),
	}
}
