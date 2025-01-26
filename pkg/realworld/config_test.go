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

func TestViperReadAndGetJson(t *testing.T) {
	viper.SetConfigType("json")
	file, err := os.Open("testdata/config.json")
	assert.Nil(t, err)
	assert.Nil(t, viper.ReadConfig(file))
	assert.Equal(t, "Example Schema", viper.Get("title"))
}

func TestViperReadAndGetYaml(t *testing.T) {
	viper.SetConfigType("yaml")
	file, err := os.Open("testdata/config.yml")
	assert.Nil(t, err)
	assert.Nil(t, viper.ReadConfig(file))
	assert.Equal(t, "steve", viper.Get("name"))
}

func TestKoanfReadAndGetJson(t *testing.T) {
	k := koanf.New(".")
	assert.Nil(t, k.Load(file.Provider("testdata/config.json"), json.Parser()))
	assert.Equal(t, "Example Schema", k.String("title"))
}

func TestKoanfReadAndGetYaml(t *testing.T) {
	k := koanf.New(".")
	assert.Nil(t, k.Load(file.Provider("testdata/config.yml"), yaml.Parser()))
	assert.Equal(t, "steve", k.String("name"))
}

func BenchmarkReadConfig(b *testing.B) {
	b.Run("Viper json load", func(b *testing.B) {
		viper.SetConfigType("json")
		for i := 0; i < b.N; i++ {
			if file, err := os.Open("testdata/config.json"); err != nil {
				b.Errorf("Unexpected error: %v", err)
			} else {
				_ = viper.ReadConfig(file)
			}
		}
	})
	b.Run("Koanf json load", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider("testdata/config.json"), json.Parser())
		}
	})
	b.Run("Viper yaml load", func(b *testing.B) {
		viper.SetConfigType("yaml")
		for i := 0; i < b.N; i++ {
			if file, err := os.Open("testdata/config.yml"); err != nil {
				b.Errorf("Unexpected error: %v", err)
			} else {
				_ = viper.ReadConfig(file)
			}
		}
	})
	b.Run("Koanf yaml load", func(b *testing.B) {
		k := koanf.New(".")
		for i := 0; i < b.N; i++ {
			_ = k.Load(file.Provider("testdata/config.yml"), json.Parser())
		}
	})
}
