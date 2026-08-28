package service

type IDGenerator struct {
	prefix string
	next   int
}

func NewIDGenerator(prefix string) *IDGenerator { return &IDGenerator{prefix: prefix, next: 1} }

func (g *IDGenerator) Next() string {
	id := g.prefix + "-" + formatNumber(g.next)
	g.next++
	return id
}

func formatNumber(value int) string {
	if value < 10 {
		return "00" + string(rune('0'+value))
	}
	if value < 100 {
		return "0" + string(rune('0'+value/10)) + string(rune('0'+value%10))
	}
	return numberString(value)
}

func numberString(value int) string {
	if value == 0 {
		return "0"
	}
	digits := make([]byte, 0, 8)
	for value > 0 {
		digits = append([]byte{byte('0' + value%10)}, digits...)
		value /= 10
	}
	return string(digits)
}
