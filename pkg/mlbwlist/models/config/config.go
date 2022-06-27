package config

type Config struct {
	UUID     string   `json:"uuid"`
	App      App      `mapstructure:"app" json:"app"`
	Database Database `mapstructure:"database" json:"database"`
	GSuite   GSuite   `mapstructure:"gsuite" json:"gsuite"`
}
