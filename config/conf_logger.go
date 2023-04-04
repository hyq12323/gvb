package config

type Logger struct {
	level        string `yaml:"level"`
	prefix       string `yaml:"prefix"`
	director     string `yaml:"director"`
	ShoLine      bool   `yaml:"show_Line"`      // 是否显示行号
	LogInConsole bool   `yaml:"log_in_console"` //是否显示打印的路径
}
