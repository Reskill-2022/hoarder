package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"sync"

	"github.com/Reskill-2022/hoarder/log"
	"github.com/spf13/cast"
)

type (
	Config interface {
		GetString(key string) string
		AddFromProvider(p Provider)
	}

	Provider interface {
		Provide() map[string]interface{}
	}

	cfg struct {
		mu  sync.RWMutex
		env map[string]interface{}
	}

	StaticProvider struct {
		env map[string]interface{}
	}
)

func New() Config {
	return &cfg{
		env: map[string]interface{}{},
	}
}

func NewStaticProvider(env map[string]interface{}) Provider {
	return &StaticProvider{
		env: env,
	}
}

func (c *cfg) AddFromProvider(p Provider) {
	if p == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.env == nil {
		c.env = map[string]interface{}{}
	}

	for k, v := range p.Provide() {
		c.env[k] = v
	}
}

func (c *cfg) GetString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.env[key]
	if !ok {
		return ""
	}
	return cast.ToString(value)
}

func (p *StaticProvider) Provide() map[string]interface{} {
	return p.env
}

func GetEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func MustGetEnv(ctx context.Context, key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.FromContext(ctx).Named("config.MustGetEnv").Fatal(fmt.Sprintf("key '%s' not found in environment", key))
	}
	return value
}

func GetBase64EncodedEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	// decode base64
	base64Decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return fallback
	}

	return string(base64Decoded)
}
