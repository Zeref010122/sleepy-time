package main

import (
	"time"

	sleepytime "github.com/Zeref010122/sleepy-time/configure"
)

func main() {
	config := sleepytime.NewConfig()
	config.SetRunDuration(5 * time.Hour) // Set the run duration to 5 hours
	config.SetDelay(5 * time.Second)     // Set the delay between actions to 5 seconds
	config.SetKeys([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"})
	config.SetMoveMouse(true)
	config.SetTerminationKeys([]string{"esc", "a"})
	sleepytime.Run(config)
}
