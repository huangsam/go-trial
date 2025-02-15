package util_test

import (
	"testing"
	"text/template"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestRenderToString(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     any
		expected string
		valid    bool
	}{
		{
			name:     "simple template",
			template: "Hello, {{.Name}}!",
			data:     map[string]string{"Name": "World"},
			expected: "Hello, World!",
			valid:    true,
		},
		{
			name:     "missing data",
			template: "Hello, {{.Name}}!",
			data:     map[string]string{},
			expected: "Hello, <no value>!",
			valid:    true,
		},
		{
			name:     "invalid template",
			template: "Hello, {{.Name",
			data:     map[string]string{"Name": "World"},
			expected: "",
			valid:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := template.New("test").Parse(tt.template)
			if !tt.valid {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)

			result, err := util.RenderToString(tmpl, tt.data)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
