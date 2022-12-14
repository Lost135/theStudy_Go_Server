package structs

type Conf struct {
	Port  string `yaml:"port"`
	Mysql string `yaml:"mysql"`
	Redis struct {
		Addr   string `yaml:"addr"`
		Passwd string `yaml:"passwd"`
		Db     int    `yaml:"db"`
	}
	Mongodb struct {
		Addr     string `yaml:"addr"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Log struct {
		ApiLogPath string `yaml:"apiLogPath"`
		ApiLogFile string `yaml:"apiLogFile"`
		SysLogPath string `yaml:"sysLogPath"`
		SysLogFile string `yaml:"sysLogFile"`
	}
}
