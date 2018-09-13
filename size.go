// Package unit contains type wrappers for common units like size. Conversion from and to human readable formats are
// supported.
package unit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Unit constants.
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

var pattern = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)\s*([KMGTPE]B?|B)?$`)

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

// ParseSize converts string to Size. Supported formats:
//  - 15
//  - 20 KB
//  - 30 M
//  - 40M
func ParseSize(s string) (Size, error) {
	f, unit, err := parseSize(s)

	if err != nil {
		return Size(0), err
	}

	if unit == "" {
		return Size(uint64(f)), nil
	}

	var bytes uint64

	switch unit {
	case "B":
		bytes = uint64(f)
	case "K":
		bytes = uint64(f * KB)
	case "M":
		bytes = uint64(f * KB)
	case "G":
		bytes = uint64(f * GB)
	case "T":
		bytes = uint64(f * TB)
	case "P":
		bytes = uint64(f * PB)
	default:
		bytes = uint64(f)
	}

	return Size(bytes), nil
}

func parseSize(s string) (float64, string, error) {
	parts := pattern.FindStringSubmatch(strings.TrimSpace(s))

	if len(parts) < 3 {
		return 0, "", errors.Errorf("invalid size format: %s", s)
	}

	f, err := strconv.ParseFloat(parts[1], 64)

	if err != nil {
		return 0, "", errors.Wrapf(err, "invalid size format: %s", s)
	}

	if f < 0 {
		return 0, "", errors.Errorf("negative size: %s", s)
	}

	unit := strings.ToUpper(parts[2])

	if len(unit) > 1 {
		unit = unit[:1]
	}

	return f, unit, nil
}
