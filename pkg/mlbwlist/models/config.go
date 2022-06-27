package models

type Config struct {
	UUID        string            `json:"uuid"`
	Application ConfigApplication `mapstructure:"app" json:"app"`
	Database    ConfigDatabase    `mapstructure:"database" json:"database"`
	GSuite      ConfigGsuite      `mapstructure:"gsuite" json:"gsuite"`
}

type ConfigApplication struct {
	ConfigVersion string                   `mapstructure:"configversion" json:"configversion"`
	SecretVersion string                   `mapstructure:"secretversion" json:"secretversion"`
	AppService    ConfigApplicationService `mapstructure:"service" json:"service"`
}

type ConfigApplicationService struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConfigDatabase struct {
	Type     string `mapstructure:"type" json:"type"`
	Filepath string `mapstructure:"filepath" json:"filepath"`
	Username string `mapstructure:"username" json:"-"`
	Password string `mapstructure:"password" json:"-"`
}

type ConfigGsuite struct {
	API      string `mapstructure:"api" json:"api"`
	Token    string `mapstructure:"token" json:"-"`
	Username string `mapstructure:"username" json:"-"`
	Password string `mapstructure:"password" json:"-"`
}
