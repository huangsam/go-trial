package realworld_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/realworld"
	kjson "github.com/knadh/koanf/parsers/json"
	kyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testJSONPath = fixturesPath + "/config.json"
	testYAMLPath = fixturesPath + "/config.yml"
)

func assertAppConfig(t *testing.T, cfg *realworld.AppConfig) {
	assert.Equal(t, "jenkins", cfg.App)
	assert.Equal(t, []string{"ci", "cd"}, cfg.Roles)
	assert.Equal(t, "-Xms256m -Xmx2048m", cfg.Extra["jvmopt"])
}

func TestJsonConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		require.NoError(t, k.Load(file.Provider(testJSONPath), kjson.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg realworld.AppConfig
		require.NoError(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, &cfg)
	})

	t.Run("Raw", func(t *testing.T) {
		cfg, err := realworld.ReadJSONConfigRaw(testJSONPath)
		require.NoError(t, err)
		assertAppConfig(t, cfg)
	})
}

func TestYamlConfig(t *testing.T) {
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		require.NoError(t, k.Load(file.Provider(testYAMLPath), kyaml.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg realworld.AppConfig
		require.NoError(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, &cfg)
	})

	t.Run("Raw", func(t *testing.T) {
		cfg, err := realworld.ReadYAMLConfigRaw(testYAMLPath)
		require.NoError(t, err)
		assertAppConfig(t, cfg)
	})
}

func BenchmarkJsonConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for b.Loop() {
			_ = k.Load(file.Provider(testJSONPath), kjson.Parser())
			var cfg realworld.AppConfig
			_ = k.Unmarshal("", &cfg)
		}
	})

	b.Run("Raw", func(b *testing.B) {
		if _, err := realworld.ReadJSONConfigRaw(testJSONPath); err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	})
}

func BenchmarkYamlConfig(b *testing.B) {
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for b.Loop() {
			_ = k.Load(file.Provider(testYAMLPath), kjson.Parser())
			var cfg realworld.AppConfig
			_ = k.Unmarshal("", &cfg)
		}
	})

	b.Run("Raw", func(b *testing.B) {
		if _, err := realworld.ReadYAMLConfigRaw(testYAMLPath); err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	})
}
