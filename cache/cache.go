package cache

import (
	"github.com/gouef/configuration/cache/file"
	"github.com/gouef/configuration/cache/memory"
	"github.com/gouef/configuration/cache/redis"
	"github.com/gouef/configuration/helper"
	"gopkg.in/yaml.v3"
)

type Cache struct {
	Storages []CacheStorageItem `yaml:"storages"`
	Custom   helper.Custom
}

type CacheStorageItem struct {
	Type     string        `yaml:"type"`
	Instance string        `yaml:"instance"`
	Name     string        `yaml:"name"`
	File     file.File     `yaml:"file"`
	Memory   memory.Memory `yaml:"memory"`
	Redis    redis.Redis   `yaml:"redis"`
	Custom   helper.Custom
}

func (r *Cache) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig Cache
	var raw rawConfig
	custom, err := helper.ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*r = Cache(raw)
	r.Custom = custom
	return nil
}

func (r *CacheStorageItem) UnmarshalYAML(value *yaml.Node) error {
	type rawConfig CacheStorageItem
	var raw rawConfig
	custom, err := helper.ParseKnownAndCustomAuto(value, &raw)
	if err != nil {
		return err
	}
	*r = CacheStorageItem(raw)
	r.Custom = custom
	return nil
}
