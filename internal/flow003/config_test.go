package flow003

import (
	"testing"
	"training-review/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.Default()
	if !cfg.Validate() || cfg.DatabasePath == "" || cfg.Address == "" {
		t.Fatal("default config is invalid")
	}
}
