package configuration

import (
	"github.com/gouef/configuration/cache"
	"github.com/gouef/configuration/helper"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type ConfigInterface interface {
	UnmarshalYAML(value *yaml.Node) error
}

type Config struct {
	Parameters map[string]any `yaml:"parameters"`
	Renderer   Renderer       `yaml:"renderer"`
	Router     Router         `yaml:"router"`
	Cache      cache.Cache    `yaml:"cache"`
	Diago      Diago          `yaml:"diago"`
	Custom     helper.Custom
}

func DefaultConfig() *Config {
	rootDir, _ := filepath.Abs(".")
	cfg := Config{}
	cfg.Parameters = map[string]any{
		"rootDir": rootDir,
	}
	cfg.Renderer = Renderer{Dir: "./views/templates", Layout: []string{"@layout", "base", "layout"}}
	cfg.Router.Statics = []RouterStatic{
		{Path: "/static", Root: "./static"},
		{Path: "/assets", Root: "./static/assets"},
	}
	cfg.Diago.Enabled = true
	return &cfg
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := DefaultConfig()
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Config
	var raw rawConfig
	custom, err := helper.ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*c = Config(raw)
	c.Custom = custom
	return nil
}
