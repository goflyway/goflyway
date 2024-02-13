package utils

import (
	"bytes"
	"errors"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

var flywaySourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	flywaySourceDir = regexp.MustCompile(`utils.utils\.go`).ReplaceAllString(file, "")
}

func FormatTemplate(tpl string, envs map[string]interface{}) (string, error) {
	t := template.Must(template.New("").Parse(tpl))
	buffer := bytes.NewBuffer(nil)
	err := t.Execute(buffer, envs)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// StringIfNull 返回第一个不为空的字符串
func StringIfNull(str ...string) string {
	for _, item := range str {
		if item != "" {
			return item
		}
	}
	return ""
}

func VersionToInt(version string) (int, error) {
	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, "V")
	versionParts := strings.Split(version, ".")
	var result int
	for _, part := range versionParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return 0, errors.New("Unable to convert version to integer:" + version)
		}
		result = result*1000 + num
	}
	return result, nil
}

var versionPrefix = []string{"V", "v", "R", "r"}

// VersionCompare 版本对比
// version1 > version2 -> return 1
// version1 = version2 -> return 0
// version1 < version2 -> return -1
func VersionCompare(version1, version2 string) (int, error) {
	for _, vp := range versionPrefix {
		version1 = strings.TrimPrefix(version1, vp)
		version2 = strings.TrimPrefix(version2, vp)
	}
	versionParts1 := strings.Split(version1, ".")
	versionParts2 := strings.Split(version2, ".")
	if len(versionParts1) > len(versionParts2) {
		return 1, nil
	} else if len(versionParts1) < len(versionParts2) {
		return -1, nil
	} else {
		l := len(versionParts1)
		for i := 0; i < l; i++ {
			num1, err := strconv.Atoi(versionParts1[i])
			if err != nil {
				return 0, errors.New("Unable to convert version to integer:" + version1)
			}
			num2, err := strconv.Atoi(versionParts2[i])
			if err != nil {
				return 0, errors.New("Unable to convert version to integer:" + version2)
			}
			if num1 > num2 {
				return 1, nil
			} else if num1 < num2 {
				return -1, nil
			}
		}
		return 0, nil
	}
}

func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, flywaySourceDir) || strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
