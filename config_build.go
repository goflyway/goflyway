package flyway

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

func configLocationCallback(c *Config) error {
	if len(c.Locations) == 0 {
		// set default location
		c.Locations = append(c.Locations, "db_migration")
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
