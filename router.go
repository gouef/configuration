package configuration

import "gopkg.in/yaml.v3"

type Router struct {
	Statics []RouterStatic `yaml:"statics"`
	Proxy   RouterProxy    `yaml:"proxy"`
	Custom  Custom
}

type RouterStatic struct {
	Path   string `yaml:"path"`
	Root   string `yaml:"root"`
	Custom Custom
}

type RouterProxy struct {
	Trust  []string `yaml:"trust"`
	Custom Custom
}

func (r *Router) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Router
	var raw rawConfig
	custom, err := ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*r = Router(raw)
	r.Custom = custom
	return nil
}
