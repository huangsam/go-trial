package realworld_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperReadAndGetYaml(t *testing.T) {
	viper.SetConfigType("yaml")
	file, err := os.Open("testdata/config.yml")
	assert.Nil(t, err)
	viper.ReadConfig(file)
	assert.Equal(t, "steve", viper.Get("name"))
}

func TestViperReadAndGetJson(t *testing.T) {
	viper.SetConfigType("json")
	file, err := os.Open("testdata/config.json")
	assert.Nil(t, err)
	viper.ReadConfig(file)
	assert.Equal(t, "Example Schema", viper.Get("title"))
}
