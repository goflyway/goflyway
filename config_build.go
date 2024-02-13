package flyway

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/consts"
	"github.com/goflyway/goflyway/logger"
	"strings"
)

type configCallback func(c *Config) error
type configCallbackContain []configCallback

func (c *configCallbackContain) Registry(f func(c *Config) error) {
	val := append(*c, f)
	*c = val
}

// configBuild 配置构建：默认值设置、配置值规格校验
func configBuild(c *Config) error {
	var callbacks configCallbackContain
	// 注册处理器
	callbacks.Registry(configDefaultValCallback)
	callbacks.Registry(configLocationCallback)
	callbacks.Registry(configSchemaCallback)
	// 执行处理
	for _, callback := range callbacks {
		err := callback(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func configDefaultValCallback(c *Config) error {
	if c.Logger == nil {
		c.Logger = logger.Default
	}
	if c.Table == "" {
		c.Table = consts.DEFAULT_HISTORY_TABLE
	}
	if len(c.Locations) == 0 {
		// set default location
		c.Locations = append(c.Locations, consts.LOCATION_DEFAULT)
	}
	return nil
}

// 处理location值
func configLocationCallback(c *Config) error {
	prefixList := []string{consts.LOCATION_PREFIX_OS}
	for _, l := range c.Locations {
		if strings.Contains(l, consts.LOCATION_PREFIX_SEQ) {
			prefixIsErr := true
			for _, prefix := range prefixList {
				if strings.HasPrefix(l, prefix) {
					prefixIsErr = false
					break
				}
			}
			if prefixIsErr {
				return errors.New(fmt.Sprintf("The config location[%s]  prefix name error", l))
			}
		}

	}
	return nil
}

func configSchemaCallback(c *Config) error {
	if len(c.Schemas) == 0 {
		return nil
	}
	for _, item := range c.Schemas {
		if item != "" {
			c.Schemas = append(c.Schemas, item)
		}
	}
	return nil
}
