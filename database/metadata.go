package database

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/utils"
	"os"
	"path/filepath"
	"runtime"
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
	return utils.FormatTemplate(string(b), map[string]interface{}{
		"schema": schema,
		"table":  table,
	})
}
