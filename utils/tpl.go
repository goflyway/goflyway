package utils

import (
	"bytes"
	"text/template"
)

func FormatTemplate(tpl string, envs map[string]interface{}) (string, error) {
	t := template.Must(template.New("").Parse(tpl))
	buffer := bytes.NewBuffer(nil)
	err := t.Execute(buffer, envs)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
