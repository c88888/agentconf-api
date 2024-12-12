package config

type target struct {
	TargetUp    string `mapstructure:"TargetUp" json:"TargetUp" yaml:"TargetUp"`
	LoadPerCore string `mapstructure:"LoadPerCore" json:"LoadPerCore" yaml:"LoadPerCore"`
	MemUtil     string `mapstructure:"MemUtil" json:"MemUtil" yaml:"MemUtil"`
	DiskUtil    string `mapstructure:"DiskUtil" json:"DiskUtil" yaml:"DiskUtil"`
}

type prom struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	User string `mapstructure:"user" json:"user" yaml:"user"`
	Auth string `mapstructure:"auth" json:"auth" yaml:"auth"`
}
