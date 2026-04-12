package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseCustomDuration(input string) (time.Duration, error) {
	// Регулярное выражение: ищем число (\d+) и букву внутри скобок \(([hmsd])\)
	re := regexp.MustCompile(`^(\d+)\(([hmsd])\)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid format: %s", input)
	}

	// Парсим число
	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	// Определяем множитель времени на основе буквы в скобках
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
