package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("⚡ [FORGE OS - THE REACTOR] INITIALIZING NEW YEAR IGNITION...")

	rand.Seed(time.Now().UnixNano())

	manifests := []string{
		"BEYOND INFINITE", "REALITY_MALLEABLE", "NO_LIMITS_FOUND",
		"DOMINATING_TIME", "SYSTEM_EVOLVING", "UNSTOPPABLE_WILL",
	}

	go func() {
		for {
			delay := rand.Intn(1200) + 300
			time.Sleep(time.Duration(delay) * time.Millisecond)

			tag := manifests[rand.Intn(len(manifests))]
			output := 1000 + rand.Intn(500)

			fmt.Printf("🔥 [IGNITION] Output: %d%% | Status: %s | ⚛️\n", output, tag)
		}
	}()

	fmt.Println("🚀 NEW YEAR DETONATED. Reality is now being rewritten...")
	fmt.Println("🏗️  Build from the ruins. Rule the present.")

	select {}
}
