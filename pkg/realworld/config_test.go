package realworld_test

import (
	"os"
	"testing"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type appConfig struct {
	App   string
	Roles []string
	Extra map[string]string
}

func assertAppConfig(t *testing.T, cfg appConfig) {
	assert.Equal(t, "jenkins", cfg.App)
	assert.Equal(t, []string{"ci", "cd"}, cfg.Roles)
	assert.Equal(t, "-Xms256m -Xmx2048m", cfg.Extra["jvmopt"])
}

func TestJsonConfig(t *testing.T) {
	t.Run("Viper", func(t *testing.T) {
		viper.SetConfigType("json")
		file, err := os.Open("testdata/config.json")
		assert.Nil(t, err)
		assert.Nil(t, viper.ReadConfig(file))
		assert.Equal(t, "jenkins", viper.Get("app"))
		var cfg appConfig
		assert.Nil(t, viper.Unmarshal(&cfg))
		assertAppConfig(t, cfg)
	})
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider("testdata/config.json"), json.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg appConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, cfg)
	})
}

func TestYamlConfig(t *testing.T) {
	t.Run("Viper", func(t *testing.T) {
		viper.SetConfigType("yaml")
		file, err := os.Open("testdata/config.yml")
		assert.Nil(t, err)
		assert.Nil(t, viper.ReadConfig(file))
		assert.Equal(t, "jenkins", viper.Get("app"))
		var cfg appConfig
		assert.Nil(t, viper.Unmarshal(&cfg))
		assertAppConfig(t, cfg)
	})
	t.Run("Koanf", func(t *testing.T) {
		k := koanf.New(".")
		assert.Nil(t, k.Load(file.Provider("testdata/config.yml"), yaml.Parser()))
		assert.Equal(t, "jenkins", k.String("app"))
		var cfg appConfig
		assert.Nil(t, k.Unmarshal("", &cfg))
		assertAppConfig(t, cfg)
	})
}

func BenchmarkJsonConfig(b *testing.B) {
	b.Run("Viper", func(b *testing.B) {
		viper.SetConfigType("json")
		for i := 0; i < b.N; i++ {
			if file, err := os.Open("testdata/config.json"); err != nil {
				b.Errorf("Unexpected error: %v", err)
			} else {
				_ = viper.ReadConfig(file)
				var cfg appConfig
				_ = viper.Unmarshal(&cfg)
			}
		}
	})
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider("testdata/config.json"), json.Parser())
			var cfg appConfig
			_ = k.Unmarshal("", &cfg)
		}
	})
}

func BenchmarkYamlConfig(b *testing.B) {
	b.Run("Viper", func(b *testing.B) {
		viper.SetConfigType("yaml")
		for i := 0; i < b.N; i++ {
			if file, err := os.Open("testdata/config.yml"); err != nil {
				b.Errorf("Unexpected error: %v", err)
			} else {
				_ = viper.ReadConfig(file)
				var cfg appConfig
				_ = viper.Unmarshal(&cfg)
			}
		}
	})
	b.Run("Koanf", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider("testdata/config.yml"), json.Parser())
			var cfg appConfig
			_ = k.Unmarshal("", &cfg)
		}
	})
}
