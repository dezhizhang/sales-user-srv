package model

type SalesUserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SalesServerConfig struct {
	Name               string             `mapstructure:"name"`
	SalesUserSrvConfig SalesUserSrvConfig `mapstructure:"sales_user_srv"`
}
