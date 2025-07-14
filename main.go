package main

import (
	"time"

	"github.com/Zeref010122/sleepy-time/configure"
)

func main() {
	config := configure.GhostConfig{
		Repetitions: 20,
		Delay:       1 * time.Second,
		Keys:        []string{"a", "b", "enter", "space"},
		MoveMouse:   true,
	}

	configure.RunGhostActions(config)
}
