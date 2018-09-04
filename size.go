package unit

import (
	"fmt"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

// Size represents integer size.
type Size uint64

// String returns "unit" presentation (value immediately followed by unit).
func (s Size) String() string {
	return s.format(false)
}

// Humanize returns "humanized" presentation (value and unit separated by single space).
func (s Size) Humanize() string {
	return s.format(true)
}

func (s Size) format(humanize bool) string {
	if s == 0 {
		return "0"
	}

	var f float32
	unit := ""

	switch {
	case s >= EB:
		f = float32(s) / EB
		unit = "EB"
	case s >= PB:
		f = float32(s) / PB
		unit = "PB"
	case s >= TB:
		f = float32(s) / TB
		unit = "TB"
	case s >= GB:
		f = float32(s) / GB
		unit = "GB"
	case s >= MB:
		f = float32(s) / MB
		unit = "MB"
	case s >= KB:
		f = float32(s) / KB
		unit = "KB"
	default:
		f = float32(s)
		unit = "B"
	}

	sf := strings.TrimSuffix(fmt.Sprintf("%.1f", f), ".0")

	if humanize {
		unit = " " + unit
	}

	return sf + unit
}
