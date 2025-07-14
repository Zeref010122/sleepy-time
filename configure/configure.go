package configure

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type GhostConfig struct {
	Repetitions int
	Delay       time.Duration
	Keys        []string
	MoveMouse   bool
}

func RunGhostActions(cfg GhostConfig) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	screenWidth, screenHeight := robotgo.GetScreenSize()

	for i := 0; i < cfg.Repetitions; i++ {
		if cfg.MoveMouse {
			x := r.Intn(screenWidth)
			y := r.Intn(screenHeight)
			robotgo.MoveMouse(x, y)
		}

		key := cfg.Keys[r.Intn(len(cfg.Keys))]
		robotgo.KeyTap(key)

		fmt.Println("key tap:", key)

		time.Sleep(cfg.Delay)
	}
}
