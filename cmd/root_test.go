package cmd

import "testing"

func TestRootCommandRegistersLanguageOnce(t *testing.T) {
	count := 0
	for _, command := range rootCmd.Commands() {
		if command.Name() == "language" {
			count++
		}
	}

	if count != 1 {
		t.Fatalf("expected language command to be registered once, got %d", count)
	}
}
