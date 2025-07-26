package bot

import "log"

var botRegistry = make(map[string]Bot)

func RegisterBot(b Bot) {
	platform := b.Platform()
	log.Printf("Registering bot for platform: %s", platform)

	botRegistry[platform] = b
}

func GetBot(platform string) (Bot, bool) {
	b, ok := botRegistry[platform]
	return b, ok
}

func StartAllBots() {
	log.Println("Starting all registered bots...")
	for platform, b := range botRegistry {
		go func(p string, botInstance Bot) {
			if err := botInstance.Start(); err != nil {
				log.Fatalf("Failed to start bot for platform %s: %v", platform, err)
			}
		}(platform, b)
	}
}
