package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NanoHubInfo holds the schema definition for the NanoHubInfo entity.
// Stores MDM DeviceInformation response data per Apple device.
type NanoHubInfo struct {
	ent.Schema
}

// Fields of the NanoHubInfo.
func (NanoHubInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Float("available_device_capacity").Optional().Default(0).
			Comment("Available storage in GB"),
		field.Bool("awaiting_configuration").Optional().Default(false).
			Comment("DEP awaiting configuration"),
		field.Float("battery_level").Optional().Default(-1).
			Comment("Battery level (0-1, -1 = no battery)"),
		field.String("bluetooth_mac").Optional().Default("").
			Comment("Bluetooth MAC address"),
		field.String("build_version").Optional().Default("").
			Comment("OS build version (e.g. 25D122)"),
		field.String("current_console_managed_user").Optional().Default("").
			Comment("UUID of current managed user"),
		field.Float("device_capacity").Optional().Default(-1).
			Comment("Total storage in GB"),
		field.String("device_name").Optional().Default("").
			Comment("User-assigned device name"),
		field.String("eacs_preflight").Optional().Default("").
			Comment("EACS preflight status"),
		field.String("ethernet_mac").Optional().Default("").
			Comment("Ethernet MAC address"),
		field.String("wifi_mac").Optional().Default("").
			Comment("WiFi MAC address"),
		field.Bool("has_battery").Optional().Default(false).
			Comment("Device has battery"),
		field.String("hostname").Optional().Default("").
			Comment("Full hostname"),
		field.Bool("is_activation_lock_enabled").Optional().Default(false).
			Comment("Activation Lock enabled"),
		field.Bool("is_activation_lock_supported").Optional().Default(false).
			Comment("Activation Lock supported"),
		field.Bool("is_apple_silicon").Optional().Default(false).
			Comment("Apple Silicon chip"),
		field.Bool("is_supervised").Optional().Default(false).
			Comment("Device is supervised"),
		field.String("localhostname").Optional().Default("").
			Comment("Bonjour local hostname"),
		field.String("model").Optional().Default("").
			Comment("Model identifier (e.g. Mac14,15)"),
		field.String("model_name").Optional().Default("").
			Comment("Human-readable model name"),
		field.Bool("auto_check_enabled").Optional().Default(false).
			Comment("OS update auto-check"),
		field.Bool("automatic_app_installation_enabled").Optional().Default(false).
			Comment("Auto app install"),
		field.Bool("automatic_os_installation_enabled").Optional().Default(false).
			Comment("Auto OS install"),
		field.Bool("automatic_security_updates_enabled").Optional().Default(false).
			Comment("Auto security updates"),
		field.Bool("background_download_enabled").Optional().Default(false).
			Comment("Background downloads"),
		field.String("catalog_url").Optional().Default("").
			Comment("Software update catalog URL"),
		field.Bool("is_default_catalog").Optional().Default(false).
			Comment("Using default catalog"),
		field.Time("previous_scan_date").Optional().Nillable().
			Comment("Last update scan date"),
		field.Int64("previous_scan_result").Optional().Default(0).
			Comment("Last scan result code"),
		field.String("os_version").Optional().Default("").
			Comment("OS version (e.g. 26.3)"),
		field.Bool("pin_required_for_device_lock").Optional().Default(false).
			Comment("PIN required for lock"),
		field.Bool("pin_required_for_erase_device").Optional().Default(false).
			Comment("PIN required for erase"),
		field.String("product_name").Optional().Default("").
			Comment("Product name (e.g. macOS, Mac14,15)"),
		field.String("provisioning_udid").Optional().Default("").
			Comment("Provisioning UDID"),
		field.String("serial_number").Optional().Default("").
			Comment("Hardware serial number"),
		field.String("software_update_device_id").Optional().Default("").
			Comment("Software update device ID"),
		field.String("supplemental_build_version").Optional().Default("").
			Comment("Supplemental build version"),
		field.Bool("supports_lom_device").Optional().Default(false).
			Comment("Supports LOM"),
		field.Bool("supports_ios_app_installs").Optional().Default(false).
			Comment("Supports iOS app installs"),
		field.Bool("system_integrity_protection_enabled").Optional().Default(false).
			Comment("SIP enabled"),
		field.String("udid").Optional().Default("").
			Comment("Device UDID"),
	}
}

// Edges of the NanoHubInfo.
func (NanoHubInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("nanohubinfo").Required(),
	}
}
