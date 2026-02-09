package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Branding holds the schema definition for the Branding entity.
// This stores global provider branding settings that apply across all tenants.
type Branding struct {
	ent.Schema
}

// Fields of the Branding.
func (Branding) Fields() []ent.Field {
	return []ent.Field{
		// Logo fields - stored as base64 data URLs
		field.Text("logo_light").
			Optional().
			Comment("Logo as base64 data URL"),
		field.Text("logo_small").
			Optional().
			Comment("Favicon as base64 data URL"),

		// Color customization
		field.String("primary_color").
			Optional().
			Default("#16a34a").
			Comment("Primary brand color (hex)"),

		// Product name
		field.String("product_name").
			Optional().
			Default("OpenUEM").
			Comment("Custom product name to display"),

		// Login page customization
		field.Text("login_background_image").
			Optional().
			Comment("Login page background image as base64 data URL"),
		field.String("login_welcome_text").
			Optional().
			Comment("Welcome text shown on login page"),
	}
}

// Edges of the Branding.
func (Branding) Edges() []ent.Edge {
	return nil
}
