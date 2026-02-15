package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[Forge OS - The Reactor] is igniting...")

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Printf("⚛️  Reactor Pulse at %s: Stability 100%% | Output: Pure Energy\n", t.Format("15:04:05"))
		}
	}()

	fmt.Println("🚀 Reactor is now live. Making reality malleable...")

	select {}
}
