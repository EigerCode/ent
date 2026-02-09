package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Authentication holds the schema definition for the Authentication entity.
type Authentication struct {
	ent.Schema
}

// Fields of the Authentication.
func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("use_certificates").Optional().Default(true),
		field.Bool("allow_register").Optional().Default(true),
		field.Bool("use_OIDC").Optional().Default(false),
		field.String("OIDC_provider").Optional().Default(""),
		field.String("OIDC_issuer_url").Optional().Default(""),
		field.String("OIDC_client_id").Optional().Default(""),
		field.String("OIDC_role_admin").Optional().Default("").
			Comment("OIDC role/group that maps to admin (e.g. openuem_admin)"),
		field.String("OIDC_role_operator").Optional().Default("").
			Comment("OIDC role/group that maps to operator (e.g. openuem_operator)"),
		field.String("OIDC_role_user").Optional().Default("").
			Comment("OIDC role/group that maps to user (e.g. openuem_user)"),
		field.String("OIDC_cookie_encription_key").Optional().Default(""),
		field.String("OIDC_keycloak_public_key").Optional().Default(""),
		field.Bool("OIDC_auto_create_account").Optional().Default(true),
		field.Bool("OIDC_auto_approve").Optional().Default(true),
		field.Bool("use_passwd").Optional().Default(true),
	}
}
