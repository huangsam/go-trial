package lesson

import (
	"strings"
	ttemplate "text/template"
)

// RenderToString renders text template to string.
//
// There is a html/template which avoids XSS vulnerabilities.
// That variant is not needed for this example as the templates in
// this project are mostly plaintext.
//
// We're using a string builder as our writer for demonstration purposes.
// In production, we can execute the template to different writers, such
// as an output file, HTTP response, etc.
func RenderToString(t *ttemplate.Template, d any) (string, error) {
	var buf strings.Builder
	err := t.Execute(&buf, d)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
