package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoHubUser holds the schema definition for the NanoHubUser entity.
// Stores MDM-reported local users on Apple devices.
type NanoHubUser struct {
	ent.Schema
}

// Fields of the NanoHubUser.
func (NanoHubUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("data_quota").Optional().Nillable().
			Comment("User data quota"),
		field.Int64("data_used").Optional().Nillable().
			Comment("User data used"),
		field.Bool("has_data_to_sync").Optional().Nillable().
			Comment("User has data to sync"),
		field.Bool("has_secure_token").Optional().Nillable().
			Comment("User has SecureToken"),
		field.Bool("is_logged_in").Optional().Nillable().
			Comment("User is currently logged in"),
		field.String("username").Optional().Nillable().
			Comment("Short username"),
		field.String("fullname").Optional().Nillable().
			Comment("Full display name"),
		field.Bool("mobile_account").Optional().Nillable().
			Comment("Mobile account"),
		field.Int64("uid").Optional().Nillable().
			Comment("Unix UID"),
		field.String("user_guid").Optional().Nillable().
			Comment("User GUID"),
	}
}

// Edges of the NanoHubUser.
func (NanoHubUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("nanohubusers").Required(),
	}
}
