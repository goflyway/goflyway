package location

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"github.com/goflyway/goflyway/consts"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Location struct {
	IsFileSystem bool      // 是否系统文件目录
	Path         string    // 文件目录
	Sqls         []SqlFile // sql列表
}

func New(location ...string) ([]Location, error) {
	var locations []Location
	if len(location) > 0 {
		for _, item := range location {
			var location Location
			if strings.Index(item, consts.LOCATION_PREFIX_OS) == 0 {
				path := item[len(consts.LOCATION_PREFIX_OS):]
				location = Location{IsFileSystem: true, Path: path}
			} else {
				location = Location{IsFileSystem: false, Path: item}
			}
			err := loadSqlFile(&location)
			if err != nil {
				return locations, err
			}
			locations = append(locations, location)
		}
	}
	return locations, nil
}

func loadSqlFile(location *Location) error {
	if location.Sqls == nil {
		location.Sqls = make([]SqlFile, 0)
	}
	file, err := os.Open(location.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = filepath.Walk(location.Path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".sql" {
			return nil
		}
		fileName := info.Name()
		if !strings.HasPrefix(fileName, "V") && !strings.HasPrefix(fileName, "R") {
			return errors.New("sql file name must be V${version}__${description} or R${version}__${description}")
		}
		mod := fileName[0:1]
		fileName = fileName[1:]
		split := strings.Split(fileName, "__")
		if len(split) != 2 {
			return errors.New("sql file name must be V${version}__${description} or R${version}__${description}")
		}
		versionFull := split[0]
		version := strings.ReplaceAll(versionFull, "_", ".")
		abs, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		sf := SqlFile{
			Name:        info.Name(),
			Path:        path,
			AbsPath:     abs,
			Version:     version,
			Mod:         mod,
			Description: split[1],
		}
		location.Sqls = append(location.Sqls, sf)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// SqlFile sql变更文件
type SqlFile struct {
	Name        string // 文件名
	AbsPath     string // 绝对路径
	Path        string // 相对路径
	Mod         string
	Version     string // 版本
	Description string // 描述
}

// Content 文件内容
func (s SqlFile) Content() (string, error) {
	b, err := os.ReadFile(s.AbsPath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// MD5 文件 MD5 值
//func (l SqlFile) MD5() (string, error) {
//	file, err := os.Open(l.Path)
//	if err != nil {
//		return "", err
//	}
//	defer file.Close()
//	hash := md5.New()
//	_, err = io.Copy(hash, file)
//	if err != nil {
//		return "", err
//	}
//	m := hex.EncodeToString(hash.Sum(nil))
//	return m, nil
//}

func (l SqlFile) CheckSum() (int64, error) {
	file, err := os.Open(l.Path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return 0, err
	}
	hashSum := hash.Sum(nil)
	var checksum uint64
	if len(hashSum) >= 8 {
		checksum = binary.BigEndian.Uint64(hashSum[:8])
	} else {
		return 0, errors.New("hash value length is less than 8 bytes")
	}
	return int64(checksum), nil
}
