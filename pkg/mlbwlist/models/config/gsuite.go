package config

type GSuite struct {
	API      string `mapstructure:"api" json:"api"`
	Token    string `mapstructure:"token" json:"-"`
	Username string `mapstructure:"username" json:"-"`
	Password string `mapstructure:"password" json:"-"`
}
