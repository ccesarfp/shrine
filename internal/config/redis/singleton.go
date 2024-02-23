package redis

import (
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

type Singleton struct {
	once     sync.Once
	instance *redis.Client
}

var s *Singleton

// GetInstance creates a new instance of client Redis
func GetInstance() (*redis.Client, error) {
	if s == nil {
		s = &Singleton{}
	}

	s.once.Do(func() {
		options, err := newOptions()
		if err != nil {
			log.Fatalln(err)
		}
		s.instance = redis.NewClient(options)
	})

	return s.instance, nil
}
