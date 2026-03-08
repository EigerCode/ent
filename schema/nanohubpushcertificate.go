package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoHubPushCertificate holds the schema definition for the NanoHubPushCertificate entity.
// This stores the Apple Push Certificate per tenant for MDM push notifications.
type NanoHubPushCertificate struct {
	ent.Schema
}

// Fields of the NanoHubPushCertificate.
func (NanoHubPushCertificate) Fields() []ent.Field {
	return []ent.Field{
		field.Text("private_key_pem").Optional().Default("").
			Comment("RSA private key in PEM format (generated during CSR creation)"),
		field.Text("csr_pem").Optional().Default("").
			Comment("Certificate Signing Request in PEM format"),
		field.Text("certificate_pem").Optional().Default("").
			Comment("Apple Push Certificate in PEM format (uploaded by admin)"),
		field.String("apns_topic").Optional().Default("").
			Comment("APNs topic extracted from the push certificate"),
		field.Time("expires_at").Optional().Nillable().
			Comment("Push certificate expiration date"),
		field.Time("uploaded_at").Optional().Nillable().
			Comment("When the push certificate was last uploaded"),
	}
}

// Edges of the NanoHubPushCertificate.
func (NanoHubPushCertificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("nanohub_push"),
	}
}
