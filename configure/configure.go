package sleepytime

import (
	"time"
)

const (
	DefaultRunDuration     = 10 * time.Minute
	DefaultDelay           = 1 * time.Second
	DefaultKeys            = "a,b,enter,space"
	DefaultMoveMouse       = true
	DefaultTerminationKeys = "ctrl,alt,q"
)

type Config struct {
	RunDuration     time.Duration
	Delay           time.Duration
	Keys            []string
	MoveMouse       bool
	TerminationKeys []string
}

func (c *Config) SetDefaults() {
	if c.RunDuration == 0 {
		c.RunDuration = DefaultRunDuration
	}
	if c.Delay == 0 {
		c.Delay = DefaultDelay
	}
	if len(c.Keys) == 0 {
		c.Keys = []string{"a", "b", "enter", "space"}
	}

	if !c.MoveMouse {
		c.MoveMouse = DefaultMoveMouse
	}
	if len(c.TerminationKeys) == 0 {
		c.TerminationKeys = []string{"ctrl", "alt", "q"}
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.SetDefaults()
	return cfg
}

func (c *Config) SetRunDuration(duration time.Duration) {
	c.RunDuration = duration
}

func (c *Config) SetDelay(delay time.Duration) {
	c.Delay = delay
}

func (c *Config) SetKeys(keys []string) {
	c.Keys = keys
}

func (c *Config) SetMoveMouse(move bool) {
	c.MoveMouse = move
}

func (c *Config) SetTerminationKeys(keys []string) {
	c.TerminationKeys = keys
}
