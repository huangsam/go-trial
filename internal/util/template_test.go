package util_test

import (
	"testing"
	"text/template"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			name:     "SimpleTemplate",
			template: "Hello, {{.Name}}!",
			data:     map[string]string{"Name": "World"},
			expected: "Hello, World!",
			valid:    true,
		},
		{
			name:     "MissingData",
			template: "Hello, {{.Name}}!",
			data:     map[string]string{},
			expected: "Hello, <no value>!",
			valid:    true,
		},
		{
			name:     "InvalidTemplate",
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
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			result, err := util.RenderToString(tmpl, tt.data)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
