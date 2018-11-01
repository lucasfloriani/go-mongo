package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations and
// can be used in any place with app.Config.Variable
var Config appConfig

type appConfig struct {
	// Environment select the environment that will be used. Defaults to "production"
	Environment string `mapstructure:"environment"`
	// ServerPort is the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// Database gets info to connect to db
	Database struct {
		Test struct {
			Connection string `required:"true"`
			Database   string `required:"true"`
		}
		Development struct {
			Connection string `required:"true"`
			Database   string `required:"true"`
		}
		Production struct {
			Connection string `required:"true"`
			Database   string `required:"true"`
		}
	}
}

// Validate check if the required config about the aplication is filled.
// Emmits a panic error if doesn't
func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.Database, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.SetDefault("environment", "production")
	v.SetDefault("server_port", 8080)
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
