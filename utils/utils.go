package utils

import (
	"errors"
	"strconv"
	"strings"
)

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
