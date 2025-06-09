package memory

import "github.com/gouef/configuration/helper"

type Memory struct {
	Name   string `yaml:"name"`
	Custom helper.Custom
}
