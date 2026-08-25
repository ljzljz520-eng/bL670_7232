package service

type Clock interface{ Now() string }

type FixedClock struct{ Value string }

func (c FixedClock) Now() string {
	if c.Value == "" {
		return "2000-01-01T00:00:00Z"
	}
	return c.Value
}
