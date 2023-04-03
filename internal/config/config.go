package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Mode Mode
	Host string
	Port int
}

/*
Each item takes precedence over the item below it:

- explicit call to Set

- flag

- env

- config

- key/value store

- default
*/
func Read(filepath string) Config {
	var c Config

	v := viper.New()

	// set the config file
	v.SetConfigFile(filepath)

	// use environment
	for _, env := range os.Environ() {
		key := strings.SplitN(env, "=", 2)[0]
		v.MustBindEnv(strings.ReplaceAll(strings.ToLower(key), "_", "."), key)
	}

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = v.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	if c.App.Mode.Dev() {
		log.Info().Interface("config", c).Msg("the config")
	} else {
		log.Info().Str("app_mode", string(c.App.Mode)).Msg("the config")
	}

	return c
}
