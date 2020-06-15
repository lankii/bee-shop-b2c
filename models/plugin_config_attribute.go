package models

type PluginConfigAttribute struct {
	PluginConfig *PluginConfig `orm:"column(plugin_config);rel(fk)"`
	Attributes   string        `orm:"column(attributes);size(255);null"`
	Name         string        `orm:"column(name);size(100)"`
}
