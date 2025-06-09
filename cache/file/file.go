package file

import "github.com/gouef/configuration/helper"

type File struct {
	Name   string `yaml:"name"`
	Dir    string `yaml:"dir"`
	Custom helper.Custom
}
