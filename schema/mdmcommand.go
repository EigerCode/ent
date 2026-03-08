package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MDMCommand holds the schema definition for the MDMCommand entity.
// Tracks MDM commands sent via the console to NanoHub.
type MDMCommand struct {
	ent.Schema
}

// Fields of the MDMCommand.
func (MDMCommand) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("uuid").
			Comment("Command UUID"),
		field.Time("when").Optional().Nillable().
			Comment("Scheduled execution time"),
		field.String("type").Default("DeviceInformation").
			Comment("Command type (DeviceInformation, InstalledApplicationList, UserList, etc.)"),
		field.String("agent_id").NotEmpty().
			Comment("Agent UDID"),
	}
}

// Edges of the MDMCommand.
func (MDMCommand) Edges() []ent.Edge {
	return nil
}
