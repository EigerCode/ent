package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RecoveryCode holds the schema definition for the RecoveryCode entity.
type RecoveryCode struct {
	ent.Schema
}

// Fields of the RecoveryCode.
func (RecoveryCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code1").NotEmpty(),
		field.Bool("used1").Default(false),
		field.String("code2").NotEmpty(),
		field.Bool("used2").Default(false),
		field.String("code3").NotEmpty(),
		field.Bool("used3").Default(false),
		field.String("code4").NotEmpty(),
		field.Bool("used4").Default(false),
		field.String("code5").NotEmpty(),
		field.Bool("used5").Default(false),
		field.String("code6").NotEmpty(),
		field.Bool("used6").Default(false),
		field.String("code7").NotEmpty(),
		field.Bool("used7").Default(false),
		field.String("code8").NotEmpty(),
		field.Bool("used8").Default(false),
		field.String("code9").NotEmpty(),
		field.Bool("used9").Default(false),
		field.String("code10").NotEmpty(),
		field.Bool("used10").Default(false),
	}
}

func (RecoveryCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("recoverycodes"),
	}
}
