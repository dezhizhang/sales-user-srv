package model

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Host      string      `mapstructure:"host" json:"host"`
	Port      int         `mapstructure:"port" json:"port"`
	Name      string      `mapstructure:"name" json:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataId" json:"dataId"`
	Group     string `mapstructure:"group" json:"group"`
}
