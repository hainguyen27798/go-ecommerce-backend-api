package setting

type Config struct {
	Server ServerSettings `mapstructure:"server"`
	Mysql  MysqlSettings  `mapstructure:"mysql"`
}

type ServerSettings struct {
	Port int `mapstructure:"port"`
}

type MysqlSettings struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"user"`
	Password        string `mapstructure:"pass"`
	Database        string `mapstructure:"db_name"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
}
