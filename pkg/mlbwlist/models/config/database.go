package config

type Database struct {
	Type     string `mapstructure:"type" json:"type"`
	Filepath string `mapstructure:"filepath" json:"filepath"`
	Username string `mapstructure:"username" json:"-"`
	Password string `mapstructure:"password" json:"-"`
}
