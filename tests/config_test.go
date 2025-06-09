package tests

import (
	"github.com/gouef/configuration"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := configuration.LoadConfig("./config/config.yml")

	assert.NoError(t, err)
	assert.Equal(t, "./views/templates", cfg.Renderer.Dir)
	assert.Equal(t, "lalala", cfg.Renderer.Custom["test"])
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := configuration.LoadConfig("nonexistent.yaml")
	assert.Error(t, err)
}

func TestUnmarshalConfig_InvalidYAML(t *testing.T) {
	var c configuration.Config
	err := yaml.Unmarshal([]byte(`- invalid`), &c)
	assert.Error(t, err)
}

func TestParseKnownAndCustomAuto_InvalidKind(t *testing.T) {
	node := &yaml.Node{Kind: yaml.SequenceNode}
	_, err := configuration.ParseKnownAndCustomAuto(node, &configuration.Router{})
	assert.Error(t, err)
}

func TestParseScalarValue_StringFallback(t *testing.T) {
	assert.Equal(t, "hello", configuration.ParseScalarValue("hello"))
}
func TestIndexComma(t *testing.T) {
	assert.Equal(t, 3, configuration.IndexComma("abc,def"))
	assert.Equal(t, -1, configuration.IndexComma("abcdef"))
}
func TestParseKnownAndCustomAuto_NotPointer(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	_, err := configuration.ParseKnownAndCustomAuto(node, configuration.Router{}) // není pointer
	assert.Error(t, err)
}

func TestParseKnownAndCustomAuto_NotStruct(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	x := 123 // není struct
	_, err := configuration.ParseKnownAndCustomAuto(node, &x)
	assert.Error(t, err)
}
func TestParseKnownAndCustom_NotPointer(t *testing.T) {
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "dir"},
			{Kind: yaml.ScalarNode, Value: "./views"},
		},
	}
	_, err := configuration.ParseKnownAndCustom(node, configuration.Renderer{}, []string{"dir"}) // není pointer
	assert.Error(t, err)
}
func TestValueParse_UnknownKind(t *testing.T) {
	node := &yaml.Node{Kind: 999}
	val := configuration.ValueParse("test", node)
	assert.Equal(t, node, val)
}
func TestParseScalarValue_Types(t *testing.T) {
	assert.Equal(t, true, configuration.ParseScalarValue("true"))
	assert.Equal(t, int64(123), configuration.ParseScalarValue("123"))
	assert.Equal(t, 3.14, configuration.ParseScalarValue("3.14"))
	assert.Equal(t, "text", configuration.ParseScalarValue("text"))
}
func TestRendererConfig_CustomField(t *testing.T) {
	var cfg configuration.Renderer
	err := yaml.Unmarshal([]byte(`dir: "./views/templates"
unknown: "something"`), &cfg)

	assert.NoError(t, err)
	assert.Equal(t, "./views/templates", cfg.Dir)
	assert.Equal(t, "something", cfg.Custom["unknown"])
}
