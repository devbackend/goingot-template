package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	Port uint16 `mapstructure:"PORT" valid:"required"`
}

// Validate check config on errors
func (c *Config) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return fmt.Errorf("config: validate service config err - %v", err)
	}

	return nil
}

// Addr return server listen address
func (c *Config) Addr() string {
	return strconv.Itoa(int(c.Port))
}

// Load fill config to param
func Load(conf *Config) error {
	_ = godotenv.Load()

	govalidator.SetFieldsRequiredByDefault(true)

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("config.Load: not found homedir - %v", err)
	}
	// Search config in home directory with name ".filepath" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".config")

	viper.AutomaticEnv() // read in environment variables that match

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err = bindEnvs(*conf)
	if err != nil {
		return fmt.Errorf("config.Load: failed bind env conf - %v", err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		return fmt.Errorf("config.Load: failed viper.Unmarshal - %v", err)
	}

	err = conf.Validate()
	if err != nil {
		return fmt.Errorf("config.Load: invalid config - %v", err)
	}

	return nil
}

func bindEnvs(i interface{}, parts ...string) error {
	ifv := reflect.ValueOf(i)
	ift := reflect.TypeOf(i)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			err := bindEnvs(v.Interface(), append(parts, tv)...)
			if err != nil {
				return err
			}
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				return err
			}
		}
	}
	return nil
}
