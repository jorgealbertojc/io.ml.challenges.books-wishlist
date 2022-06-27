package config

type App struct {
	ConfigVersion string     `mapstructure:"configversion" json:"configversion"`
	SecretVersion string     `mapstructure:"secretversion" json:"secretversion"`
	AppService    AppService `mapstructure:"service" json:"service"`
}

type AppService struct {
	Port int    `mapstructure:"port" json:"port"`
	Host string `mapstructure:"host" json:"host"`
}
