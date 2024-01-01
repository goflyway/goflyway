package database

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

const metadataFileName = "createMetadata.sql"

// loadMetadataSql 加载元数据创建sql
func loadMetadataSql(t Type, schema, table string) (string, error) {
	fileName := fmt.Sprintf("./%s/%s", t, metadataFileName)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	tpl := template.Must(template.New("").Parse(string(b)))
	buffer := bytes.NewBuffer(nil)
	err = tpl.Execute(buffer, map[string]interface{}{
		"schema": schema,
		"table":  table,
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
