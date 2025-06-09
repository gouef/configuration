package configuration

import (
	"github.com/gouef/configuration/helper"
	"gopkg.in/yaml.v3"
)

type Router struct {
	Statics []RouterStatic `yaml:"statics"`
	Proxy   RouterProxy    `yaml:"proxy"`
	Custom  helper.Custom
}

type RouterStatic struct {
	Path   string `yaml:"path"`
	Root   string `yaml:"root"`
	Custom helper.Custom
}

type RouterProxy struct {
	Trust  []string `yaml:"trust"`
	Custom helper.Custom
}

func (r *Router) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Router
	var raw rawConfig
	custom, err := helper.ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*r = Router(raw)
	r.Custom = custom
	return nil
}
