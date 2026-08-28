package config

import "strings"

func (c Config) Validate() bool {
	return strings.TrimSpace(c.DatabasePath) != "" && strings.TrimSpace(c.Address) != ""
}

func (c Config) IsEphemeral() bool { return c.DatabasePath == ":memory:" }
