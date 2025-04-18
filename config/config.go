package config

import "github.com/spf13/viper"

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Port            string `mapstructure:"port"`
	DB              DBConfig
	Redis           RedisConfig
	Workers         int    `mapstructure:"workers"`
	MaxRetries      int    `mapstructure:"max_retries"`
	TopicPartitions int    `mapstructure:"topic_partitions"`
	MetricsPort     string `mapstructure:"metrics_port"`
	APIKey          string `mapstructure:"api_key"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
