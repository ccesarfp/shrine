package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

type Config struct {
	Address  string
	Password string
	Database int
	Options  *redis.Options
}

// newOptions configure environment redis vars, configure options and return a new redis.Options
func newOptions() (*redis.Options, error) {
	c := &Config{}
	err := c.setupEnvironmentVars()
	if err != nil {
		return nil, err
	}
	c.setupOptions()
	return c.Options, nil
}

// setupEnvironmentVars setup environment vars
func (c *Config) setupEnvironmentVars() error {
	intConvert, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		return err
	}
	c.Address = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	c.Password = os.Getenv("REDIS_PASSWORD")
	c.Database = intConvert
	return nil
}

// setupOptions setup redis.Options
func (c *Config) setupOptions() {
	c.Options = &redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.Database,
	}
}
