package database

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

const metadataFileName = "createMetadata.sql"

// loadMetadataSql 加载元数据创建sql
func loadMetadataSql(t Type, schema, table string) (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("get flyway runtime caller error")
	}
	dir := filepath.Dir(file)
	fileName := fmt.Sprintf("%s/%s/%s", dir, t, metadataFileName)
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
