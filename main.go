package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	notify()
	ticker := time.Tick(30 * time.Minute)
	for range ticker {
		notify()
	}
}

func notify() {
	beeep.Alert("Listen here, you little shit", "Change your Posture", "assets/warning.png")
}
