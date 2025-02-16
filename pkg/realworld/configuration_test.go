package realworld_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/realworld"
	kjson "github.com/knadh/koanf/parsers/json"
	kyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

const (
	testJsonLocation string = fixturesPath + "/config.json"
	testYamlLocation string = fixturesPath + "/config.yml"
)

func assertAppConfig(t *testing.T, cfg *realworld.AppConfig) {
	assert.Equal(t, "jenkins", cfg.App)
	assert.Equal(t, []string{"ci", "cd"}, cfg.Roles)
	assert.Equal(t, "-Xms256m -Xmx2048m", cfg.Extra["jvmopt"])
}

func TestJsonConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider(testJsonLocation), kjson.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg realworld.AppConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, &cfg)
	})

	t.Run("Raw", func(t *testing.T) {
		cfg, err := realworld.ReadJSONConfigRaw(testJsonLocation)
		assert.Nil(t, err)
		assertAppConfig(t, cfg)
	})
}

func TestYamlConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider(testYamlLocation), kyaml.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg realworld.AppConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, &cfg)
	})

	t.Run("Raw", func(t *testing.T) {
		cfg, err := realworld.ReadYAMLConfigRaw(testYamlLocation)
		assert.Nil(t, err)
		assertAppConfig(t, cfg)
	})
}

func BenchmarkJsonConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider(testJsonLocation), kjson.Parser())
			var cfg realworld.AppConfig
			_ = k.Unmarshal("", &cfg)
		}
	})

	b.Run("Raw", func(b *testing.B) {
		if _, err := realworld.ReadJSONConfigRaw(testJsonLocation); err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	})
}

func BenchmarkYamlConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider(testYamlLocation), kjson.Parser())
			var cfg realworld.AppConfig
			_ = k.Unmarshal("", &cfg)
		}
	})

	b.Run("Raw", func(b *testing.B) {
		if _, err := realworld.ReadYAMLConfigRaw(testYamlLocation); err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	})
}
