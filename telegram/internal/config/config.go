/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package config

import (
	"fmt"
	"time"

	"telegram/pkg/config"
	"telegram/pkg/errs"
	"telegram/pkg/log"
)

const (
	// DefaultPath - define default path to configuration file.
	DefaultPath = "./cmd/config.yaml"
)

type (
	// Config - define configuration model.
	Config struct {
		Delivery Delivery   `yaml:"delivery"`
		Logger   log.Config `yaml:"logger"`
	}

	// Delivery - define delivery configuration.
	Delivery struct {
		HTTPServer  HTTPServer  `yaml:"http-server"`
		TelegramBot TelegramBot `yaml:"telegram-bot"`
	}

	// HTTPServer - define http configuration.
	HTTPServer struct {
		LogRequests        bool          `yaml:"log-requests"`
		ListenAddress      string        `yaml:"listen-address"`
		ReadTimeout        time.Duration `yaml:"read-timeout"`
		WriteTimeout       time.Duration `yaml:"write-timeout"`
		BodySizeLimitBytes int           `yaml:"body-size-limit"`
		GracefulTimeout    int           `yaml:"graceful-timeout"`
	}

	// TelegramBot - define telegram bot configuration.
	TelegramBot struct {
		Debug   bool    `yaml:"debug"`
		Token   string  `yaml:"token"`
		Binance Binance `yaml:"binance"`
	}

	// Binance - define binance configuration.
	Binance struct {
		BinanceCourseURL string `yaml:"binance-course-url"`
	}
)

// NewConfig - create new configuration.
func NewConfig(appName, cfgFilePath string) (*Config, error) {
	cfg := new(Config)

	if cfgErr := cfg.loadFromFile(cfgFilePath); cfgErr != nil {
		return nil, fmt.Errorf("config loader: %s", cfgErr)
	}

	return cfg.valid()
}

// loadFromFile - load configuration from file.
func (c *Config) loadFromFile(configPath string) error {
	if err := config.LoadFromFile(configPath, c); err != nil {
		return err
	}

	return nil
}

// valid - validate configuration.
func (c *Config) valid() (*Config, error) {
	if errorsList := c.Validate(); len(errorsList) != 0 {
		return nil, &errs.FieldsValidation{Errors: errorsList}
	}

	return c, nil
}
