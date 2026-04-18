package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseCustomDuration(input string) (time.Duration, error) {
	// REGEXP: Find a number  (\d+) and letter in brackets like h(Hour), m(Minute), s(Second)  \(([hmsd])\)
	// For Example: "24(h)" = 24 Hour(or 1 day).
	re := regexp.MustCompile(`^(\d+)\(([hmsd])\)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid format: %s", input)
	}

	// Parse Number.
	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	// Define letter in brackets: "(h)", "(m)", "(s)".
	unit := matches[2]
	var duration time.Duration
	switch unit {
	case "h":
		duration = time.Duration(value) * time.Hour
	case "m":
		duration = time.Duration(value) * time.Minute
	case "s":
		duration = time.Duration(value) * time.Second
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}

	return duration, nil
}
