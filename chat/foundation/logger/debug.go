package logger

import (
	"context"
	"runtime/debug"
	"strconv"
	"strings"
)

// BuildInfo logs informations stored inside the Go binary.

func (log *Logger) BuildInfo(ctx context.Context) {
	var values []any

	info, _ := debug.ReadBuildInfo()

	for _, s := range info.Settings {
		key := s.Key

		if qouteKey(key) {
			key = strconv.Quote(key)
		}

		value := s.Value

		if qouteValue(value) {
			value = strconv.Quote(value)
		}

		values = append(values, key, value)
	}

	values = append(values, "goversion", info.GoVersion)
	values = append(values, "modversion", info.Main.Version)

	log.Info(ctx, "build info", values...)
}

// qouteKey reports whether key is required to be qouted.
func qouteKey(key string) bool {
	return len(key) == 0 || strings.ContainsAny(key, "= \t\r\n\"`")
}

// qouteValue reports whether value is required to be qouted.
func qouteValue(value string) bool {
	return len(value) == 0 || strings.ContainsAny(value, "= \t\r\n\"`")
}
