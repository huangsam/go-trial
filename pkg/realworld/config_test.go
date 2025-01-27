package realworld_test

import (
	"testing"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	koanf "github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

// appConfig represents the data model for app config.
type appConfig struct {
	App   string
	Roles []string
	Extra map[string]string
}

var (
	// testJsonLocation is a JSON path for app config.
	testJsonLocation string = "testdata/config.json"

	// testYamlLocation is a YAML path for app config.
	testYamlLocation string = "testdata/config.yml"
)

func assertAppConfig(t *testing.T, cfg appConfig) {
	assert.Equal(t, "jenkins", cfg.App)
	assert.Equal(t, []string{"ci", "cd"}, cfg.Roles)
	assert.Equal(t, "-Xms256m -Xmx2048m", cfg.Extra["jvmopt"])
}

func TestJsonConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider(testJsonLocation), json.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg appConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, cfg)
	})
}

func TestYamlConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider(testYamlLocation), yaml.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg appConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, cfg)
	})
}

func BenchmarkJsonConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider(testJsonLocation), json.Parser())
			var cfg appConfig
			_ = k.Unmarshal("", &cfg)
		}
	})
}

func BenchmarkYamlConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider(testYamlLocation), json.Parser())
			var cfg appConfig
			_ = k.Unmarshal("", &cfg)
		}
	})
}
