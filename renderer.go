package configuration

import (
	"github.com/gouef/configuration/helper"
	"gopkg.in/yaml.v3"
)

type Renderer struct {
	Dir    string   `yaml:"dir"`
	Layout []string `yaml:"layout"`
	Custom helper.Custom
}

func (c *Renderer) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Renderer
	var raw rawConfig
	custom, err := helper.ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*c = Renderer(raw)
	c.Custom = custom
	return nil
}
