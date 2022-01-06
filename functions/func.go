package common

import "github.com/xanzy/go-gitlab" gitlab

func GitClient() {
	git, err := gitlab.NewClient("5sat9MtD6VhstvsAHLxy")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}