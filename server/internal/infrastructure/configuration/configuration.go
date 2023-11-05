package configuration

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Mqtt     MqttConf     `yaml:"mqtt"`
	Database DatabaseConf `yaml:"database"`
	Logging  LoggingConf  `yaml:"logging"`
}

type MqttConf struct {
	Host string `yaml:"host" env:"mqtt" env-default:"localhost"`
	Port int    `yaml:"port" env:"mqtt" env-default:"1883"`
}

type DatabaseConf struct {
	Path          string `yaml:"path" env:"DATABASE_PATH"`
	Name          string `yaml:"name" env:"DATABASE_NAME"`
	MigrationPath string `yaml:"migrationPath" env:"DATABASE_MIGRATION_PATH"`
}

type LoggingConf struct {
	LogLevel    string `yaml:"level" env:"LOG_LEVEL"`
	PrettyPrint bool   `yaml:"prettyPrint" env:"LOG_PRETTY_PRINT" env-default:"false"`
}

func CreateConfiguration() (*Configuration, error) {
	var configPath string
	var cfg Configuration

	f := flag.NewFlagSet("Room Read", 1)
	f.StringVar(&configPath, "config", "config.yml", "Path to configuration file")

	fu := f.Usage

	header := `This server can be configured using environment variables or with a config file.z`
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(&cfg, &header)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])

	var configError error
	if _, err := os.Stat(configPath); err == nil {
		configError = cleanenv.ReadConfig(configPath, &cfg)
	} else {
		configError = cleanenv.ReadEnv(&cfg)
	}

	if configError != nil {
		return nil, configError
	}

	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Configuration) Validate() error {
	return nil
}
