package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NanoHubSettings holds the schema definition for the NanoHubSettings entity.
// This stores global NanoHub server configuration (not per-tenant).
type NanoHubSettings struct {
	ent.Schema
}

// Fields of the NanoHubSettings.
func (NanoHubSettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("server_url").Optional().Default("").
			Comment("NanoHub API base URL (e.g. https://nanohub.example.com:10000/api/v1)"),
		field.String("username").Optional().Default("").
			Comment("NanoHub API basic auth username"),
		field.String("password").Optional().Default("").
			Comment("NanoHub API basic auth password"),
		field.Text("ca_cer_file").Optional().Default("").
			Comment("CA certificate PEM content used by NanoHub for webhook TLS verification"),
		field.String("scep_url").Optional().Default("").
			Comment("SCEP server URL for Apple device certificate enrollment"),
		field.String("scep_challenge").Optional().Default("").
			Comment("SCEP challenge password"),
		field.String("mdm_url").Optional().Default("").
			Comment("NanoHub MDM endpoint URL that Apple devices connect to (e.g. https://mdm.example.com:9004/mdm)"),
		field.Text("vendor_private_key_pem").Optional().Default("").Sensitive().
			Comment("MDM Vendor RSA private key in PEM format (provided by OpenUEM)"),
		field.Text("vendor_cert_pem").Optional().Default("").
			Comment("MDM Vendor certificate in PEM format (provided by OpenUEM)"),
		field.String("enrollment_profile_id").Optional().Default("com.openuem.mdm.enrollment").
			Comment("MDM enrollment profile PayloadIdentifier used in .mobileconfig and RemoveProfile command"),
	}
}

// Edges of the NanoHubSettings.
func (NanoHubSettings) Edges() []ent.Edge {
	return nil
}
