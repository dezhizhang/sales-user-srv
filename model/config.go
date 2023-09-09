package model

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name" json:"name"`
	MysqlConfig   MysqlConfig   `mapstructure:"mysql" json:"mysql"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}
