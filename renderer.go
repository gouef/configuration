package configuration

import "gopkg.in/yaml.v3"

type Renderer struct {
	Dir    string   `yaml:"dir"`
	Layout []string `yaml:"layout"`
	Custom Custom
}

func (c *Renderer) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Renderer
	var raw rawConfig
	custom, err := ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*c = Renderer(raw)
	c.Custom = custom
	return nil
}
