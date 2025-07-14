package sleepytime

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// Run uses the Config struct and its defaults from configure.go
func Run(cfg Config) {

	cfg.SetDefaults()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	screenWidth, screenHeight := robotgo.GetScreenSize()
	stopChan := make(chan struct{})

	go func() {
		hook.Register(hook.KeyDown, cfg.TerminationKeys, func(e hook.Event) {
			fmt.Println("🛑 Termination key(s) pressed!")
			stopChan <- struct{}{}
			hook.End()
		})
		s := hook.Start()
		<-hook.Process(s)
	}()

	endTime := time.Now().Add(cfg.RunDuration)

	for {
		select {
		case <-stopChan:
			fmt.Println("👋 sleepy-time stopped by user.")
			return
		default:
			if time.Now().After(endTime) {
				fmt.Println("⏰ sleepy-time duration ended.")
				return
			}

			if cfg.MoveMouse {
				x := r.Intn(screenWidth)
				y := r.Intn(screenHeight)
				robotgo.Move(x, y)
				if cfg.ClickMouse {
					robotgo.Click()
				}
			}

			if len(cfg.Keys) > 0 {
				key := cfg.Keys[r.Intn(len(cfg.Keys))]
				robotgo.KeyTap(key)
				fmt.Printf("🔑 Key '%s' pressed\n", key)
			}

			time.Sleep(cfg.Delay)
		}
	}
}
