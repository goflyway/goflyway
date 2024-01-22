package utils

import (
	"fmt"
	"strconv"
	"testing"
)

func TestVersionToInteger(t *testing.T) {
	versions := []string{
		"V1.0.0.001",
		"V1.0.0.20240122",
	}
	for _, version := range versions {
		i, err := VersionToInt(version)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(fmt.Sprintf("%s to %d", version, i))
	}

}

func TestVersionCompare(t *testing.T) {
	versions := [][]string{
		{"V1.0.0.001", "V1.0.0.20240122", "-1"},
		{"V1.0.0.001", "V1.0.0.001", "0"},
		{"V1.0.0.20240122", "V1.0.0.001", "1"},
	}
	for _, arr := range versions {
		expect, _ := strconv.Atoi(arr[2])
		compare, err := VersionCompare(arr[0], arr[1])
		if err != nil {
			t.Fatal(err)
		}
		if compare != expect {
			t.Fatal(fmt.Sprintf("%s compare %s , result error,%d != %d", arr[0], arr[1], expect, compare))
		}
	}
}
