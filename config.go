package mgodb

import (
	"fmt"
)

type Config struct {
	User string //用户名
	Pass string //用户密码
	Host string //主机名(127.0.0.1:27017)
	Name string //数据库名
}

func (c *Config) Dsn() string {
	return fmt.Sprintf(`mongodb://%s:%s@%s/%s`, c.User, c.Pass, c.Host, c.Name)
}
